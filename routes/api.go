package routes

import (
	"gohub/app/http/controllers/api/v1/auth"
	"gohub/app/http/middlewares"

	"github.com/gin-gonic/gin"

	controllers "gohub/app/http/controllers/api/v1"
)

func RegisterApiRouters(r *gin.Engine) {
	v1 := r.Group("/v1")

	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
	// 作为参考 Github API 每小时最多 60 个请求（根据 IP）。
	// 测试时，可以调高一点。
	v1.Use(middlewares.LimitIP("200-H"))
	{
		authGroup := v1.Group("/auth")
		authGroup.Use(middlewares.LimitRouteIP("1000-H"))
		{
			//注册
			sc := new(auth.SignupController)
			authGroup.POST("/signup/phone/exists", middlewares.GuestJWT(), sc.IsPhoneExists)
			authGroup.POST("/signup/email/exists", middlewares.GuestJWT(), sc.IsEmailExists)
			authGroup.POST("/signup/using-phone", middlewares.GuestJWT(), middlewares.LimitRouteIP("60-H"), sc.SignupUsingPhone)
			authGroup.POST("/signup/using-email", middlewares.GuestJWT(), middlewares.LimitRouteIP("60-H"), sc.SignupUsingEmail)

			//登录
			lc := new(auth.LoginController)
			authGroup.POST("/login/using-phone", middlewares.GuestJWT(), lc.LoginByPhone)
			authGroup.POST("/login/using-password", middlewares.GuestJWT(), lc.LoginByPassword)
			authGroup.POST("/login/refresh-token", middlewares.AuthJWT(), lc.RefreshToken)

			//密码
			pc := new(auth.PasswordController)
			authGroup.POST("/password-reset/using-phone", middlewares.GuestJWT(), pc.ResetByPhone)
			authGroup.POST("/password-reset/using-email", middlewares.GuestJWT(), pc.ResetByEmail)

			//发送验证码
			vc := new(auth.VerifyCodeController)
			authGroup.POST("/verify-codes/captcha", middlewares.LimitRouteIP("20-H"), vc.ShowCaptcha)
			authGroup.POST("/verify-codes/phone", middlewares.LimitRouteIP("20-H"), vc.SendCodeByPhone)
			authGroup.POST("/verify-codes/email", middlewares.LimitRouteIP("50-H"), vc.SendCodeByEmail)
		}
		//用户
		uc := new(controllers.UsersController)
		v1.GET("/user", middlewares.AuthJWT(), uc.CurrentUser)
	}
}
