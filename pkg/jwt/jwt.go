package jwt

import (
	"errors"
	"gohub/pkg/app"
	"gohub/pkg/config"
	"gohub/pkg/logger"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwtpkg "github.com/golang-jwt/jwt"
)

var (
	ErrTokenExpired           error = errors.New("令牌已过期")
	ErrTokenExpiredMaxRefresh error = errors.New("令牌已过最大刷新时间")
	ErrTokenMalformed         error = errors.New("请求令牌格式有误")
	ErrTokenInvalid           error = errors.New("请求令牌无效")
	ErrHeaderEmpty            error = errors.New("需要认证才能访问！")
	ErrHeaderMalformed        error = errors.New("请求头中 Authorization 格式有误")
)

type JWT struct {
	// 秘钥，用以加密 JWT，读取配置信息 app.key
	SignKey []byte
	//刷新token的最大时间
	MaxRefresh time.Duration
}

type JwtCustomClaims struct {
	UserId       string `json:"user_id"`
	UserName     string `json:"user_name"`
	ExpireAtTime int64  `json:"expire_time"`

	//继承jwt的StandardClaims
	jwtpkg.StandardClaims
}

func NewJWT() *JWT {
	return &JWT{
		SignKey: []byte(config.GetString("app.key")),

		MaxRefresh: time.Duration(config.GetInt64("jwt.max_refresh_time")) * time.Minute,
	}
}

//解析token,中间件中使用
func (jwt *JWT) ParseToken(c *gin.Context) (*JwtCustomClaims, error) {
	tokenString, parseErr := jwt.getTokenFromHeader(c)
	if parseErr != nil {
		return nil, parseErr
	}

	//调用 jwt 库解析用户传参的 Token
	token, err := jwt.parseTokenString(tokenString)

	//解析出错
	if err != nil {
		validationError, ok := err.(*jwtpkg.ValidationError)
		if ok {
			if validationError.Errors == jwtpkg.ValidationErrorMalformed {
				return nil, ErrTokenMalformed
			} else if validationError.Errors == jwtpkg.ValidationErrorExpired {
				return nil, ErrTokenExpired
			}
		}
		return nil, ErrTokenInvalid
	}

	// 将 token 中的 claims 信息解析出来和 JWTCustomClaims 数据结构进行校验
	if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrTokenInvalid

}

func (jwt *JWT) RefreshToken(c *gin.Context) (string, error) {
	// 1. 从 Header 里获取 token
	tokenString, parseErr := jwt.getTokenFromHeader(c)
	if parseErr != nil {
		return "", parseErr
	}

	// 2. 调用 jwt 库解析用户传参的 Token
	token, err := jwt.parseTokenString(tokenString)

	// 3. 解析出错，未报错证明是合法的 Token（甚至未到过期时间）
	if err != nil {
		validationErr, ok := err.(*jwtpkg.ValidationError)
		// 满足 refresh 的条件：只是单一的报错 ValidationErrorExpired
		if !ok || validationErr.Errors != jwtpkg.ValidationErrorExpired {
			return "", err
		}
	}

	// 4. 解析 JWTCustomClaims 的数据
	claims := token.Claims.(*JwtCustomClaims)

	//5 检查是否过了运行最大刷新时间
	x := app.TimeNowInTimezone().Add(-jwt.MaxRefresh).Unix()
	if claims.IssuedAt > x {
		//修改过期时间
		claims.StandardClaims.ExpiresAt = jwt.expireAtTime()
		return jwt.createToken(*claims)
	}
	return "", ErrTokenExpiredMaxRefresh
}

func (jwt *JWT) IssueToken(userID string, userName string) string {

	//构建claim
	expiredAtTime := jwt.expireAtTime()
	claims := JwtCustomClaims{
		userID,
		userName,
		expiredAtTime,
		jwtpkg.StandardClaims{
			NotBefore: app.TimeNowInTimezone().Unix(), //签名生效时间
			IssuedAt:  app.TimeNowInTimezone().Unix(), //首次签名时间
			ExpiresAt: expiredAtTime,
			Issuer:    config.GetString("app.name"),
		},
	}
	//根据claims 生成token
	token, err := jwt.createToken(claims)
	if err != nil {
		logger.LogIf(err)
		return ""
	}
	return token
}

//生成token，内部使用
func (jwt *JWT) createToken(claims JwtCustomClaims) (string, error) {
	//使用HS256算法进行token生成
	token := jwtpkg.NewWithClaims(jwtpkg.SigningMethodHS256, claims)
	return token.SignedString(jwt.SignKey)
}

//过期时间
func (jwt *JWT) expireAtTime() int64 {
	timenow := app.TimeNowInTimezone()

	var expireTime int64
	if config.GetBool("app.debug") {
		expireTime = config.GetInt64("jwt.debug_expire_time")
	} else {
		expireTime = config.GetInt64("jwt.expire_time")
	}
	expire := time.Duration(expireTime) * time.Minute
	return timenow.Add(expire).Unix()
}

//使用 jwtpkg.ParseWithClaims 解析 Token
func (jwt *JWT) parseTokenString(tokenString string) (*jwtpkg.Token, error) {
	return jwtpkg.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(t *jwtpkg.Token) (interface{}, error) {
		return jwt.SignKey, nil
	})
}

// 使用 jwtpkg.ParseWithClaims 解析 Token
// Authorization:Bearer xxxxx
func (jwt *JWT) getTokenFromHeader(c *gin.Context) (string, error) {

	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrHeaderEmpty
	}

	//按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", ErrHeaderMalformed
	}
	return parts[1], nil
}
