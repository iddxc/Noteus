package system

import (
	v1 "Noteus/api/v1"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (r *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	userRouter := router.Group("user")
	systemApi := v1.ApiGroupApp.SystemApiGroup.SystemApi
	{
		userRouter.GET("/captcha", systemApi.GetCaptcha) // 获取验证码
		userRouter.POST("/login", systemApi.Login)       // 登录
		userRouter.POST("/register", systemApi.Register) // 注册接口
		userRouter.GET("/", systemApi.GetInfo)           // 获取账号信息
	}
}
