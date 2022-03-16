package middlewares

import (
	"gohub/pkg/jwt"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

func GuestJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if len(ctx.GetHeader("Authorization")) > 0 {
			_, err := jwt.NewJWT().ParseToken(ctx)
			if err == nil {
				response.Unauthorized(ctx, "请用游客身份登录")
				ctx.Abort()
				return
			}
		}
		ctx.Next()
	}
}
