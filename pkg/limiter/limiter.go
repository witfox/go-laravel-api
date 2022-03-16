package limiter

import (
	"gohub/pkg/config"
	"gohub/pkg/logger"
	"gohub/pkg/redis"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	limiterredis "github.com/ulule/limiter/v3/drivers/store/redis"
)

func GetKeyIP(c *gin.Context) string {
	return c.ClientIP()
}

func GetKeyRouteWithIP(c *gin.Context) string {
	return formatRoute(c.FullPath()) + c.ClientIP()
}

func CheckRate(c *gin.Context, key string, formatted string) (limiter.Context, error) {
	rate, err := limiter.NewRateFromFormatted(formatted)
	var context limiter.Context
	if err != nil {
		logger.LogIf(err)
		return context, err
	}

	//修改存储方式
	store, err := limiterredis.NewStoreWithOptions(redis.Redis.Client, limiter.StoreOptions{
		Prefix: config.GetString("app.name") + ":limiter",
	})
	if err != nil {
		logger.LogIf(err)
		return context, err
	}

	//初始化对象
	limiterObj := limiter.New(store, rate)

	//获取限流结果
	return limiterObj.Get(c, key)
}

func formatRoute(route string) string {
	route = strings.ReplaceAll(route, "/", "-")
	return strings.ReplaceAll(route, ":", "_")
}
