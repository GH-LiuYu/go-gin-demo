package router

import (
	"demo/api"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("login", api.Login)
		BaseRouter.POST("captcha", api.Captcha)
	}
	return BaseRouter
}
