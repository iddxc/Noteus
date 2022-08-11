package initialize

import (
	"Noteus/utils/plugin"
	"fmt"
	"github.com/gin-gonic/gin"
)

func PluginInit(group *gin.RouterGroup, Plugin ...plugin.Plugin) {
	for _, item := range Plugin {
		PluginGroup := group.Group(item.RouterPath())
		item.Register(PluginGroup)
	}
}

func InstallPlugin(Router *gin.Engine) {
	PublicGroup := Router.Group("")
	fmt.Println("无鉴权插件安装==>", PublicGroup)
	PrivateGroup := Router.Group("")
	fmt.Println("鉴权插件安装==>", PrivateGroup)
	// PrivateGroup.User()
	// PluginInit(PrivateGroup, Plugin)
}
