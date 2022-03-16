package bootstrap

import (
	"gohub/app/http/middlewares"
	"gohub/routes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	//注册全局中间件
	registerGlobalMiddleware(router)
	//注册路由
	routes.RegisterApiRouters(router)
	//404路由
	setup404Handler(router)
}

func registerGlobalMiddleware(router *gin.Engine) {
	router.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
	)
}

func setup404Handler(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		acceptStr := c.Request.Header.Get("Accept")
		if strings.Contains(acceptStr, "text/html") {
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}

	})
}
