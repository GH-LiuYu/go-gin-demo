package initialize

import (
	"demo/router"//各个模块路由绑定对应详细下级路由
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine  {
	var Router =gin.Default()
	//global.DEMO_LOG.Debug("use middleware logger")
	// 方便统一添加路由组前缀 多服务器上线使用
	ApiGroup := Router.Group("")
	router.InitMenuRouter(ApiGroup)                  // 注册menu路由
	router.InitBaseRouter(ApiGroup)                  // 注册基础功能路由 不做鉴权
	//global.DEMO_LOG.Info("router register success")
	return Router
}
