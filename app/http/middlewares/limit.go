package middlewares

import (
	"gohub/pkg/limiter"
	"gohub/pkg/logger"
	"gohub/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func LimitIP(limit string) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		key := limiter.GetKeyIP(ctx)
		if ok := limitHandler(ctx, key, limit); !ok {
			return
		}
		ctx.Next()
	}
}

func LimitRouteIP(limit string) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		key := limiter.GetKeyRouteWithIP(ctx)
		if ok := limitHandler(ctx, key, limit); !ok {
			return
		}
		ctx.Next()
	}
}

func limitHandler(c *gin.Context, key, limit string) bool {
	rate, err := limiter.CheckRate(c, key, limit)
	if err != nil {
		logger.LogIf(err)
		response.Abort500(c)
		return false
	}

	//设置头部
	c.Header("X-RateLimit-Limit", cast.ToString(rate.Limit))
	c.Header("X-RateLimit-Remaining", cast.ToString(rate.Remaining))
	c.Header("X-RateLimit-Reset", cast.ToString(rate.Reset))

	//超额
	if rate.Reached {
		c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
			"message": "请求太频繁了",
		})
		return false
	}
	return true
}
