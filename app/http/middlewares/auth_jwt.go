package middlewares

import (
	"gohub/app/models/user"
	"gohub/pkg/jwt"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, err := jwt.NewJWT().ParseToken(ctx)
		if err != nil {
			response.Unauthorized(ctx, "你无权限")
			//return 掉，就会中断所有的后续请求
			return
		}

		_user := user.Get(claims.UserId)
		if _user.ID == 0 {
			response.Unauthorized(ctx, "找不到对应的用户")
			return
		}

		//将用户消息存入gin.context,后续 auth 包将从这里拿到当前用户数据
		ctx.Set("current_user_id", _user.ID)
		ctx.Set("current_user_name", _user.Name)
		ctx.Set("current_user", _user)

		ctx.Next()
	}
}
