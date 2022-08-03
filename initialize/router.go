package initialize

import (
	"Noteus/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	nousRouter := router.RouterGroupApp.Nous

	PublicGroup := Router.Group("")
	{
		// 健康检测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	nousRouter.InitNousRouter(PublicGroup)

	return Router
}
