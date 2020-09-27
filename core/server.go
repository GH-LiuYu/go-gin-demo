package core

import (
	"demo/global"
	"demo/initialize"
	"fmt"
	"time"
)

type server interface {
	ListenAndServe() error
}
func RunWindowsServer()  {
	if global.DEMO_CONFIG.System.UseMultipoint {
		// 初始化redis服务
		//initialize.Redis()
	}
	Router := initialize.Routers()
	//Router.Static("/form-generator", "./resource/page")//用于配置静态文件的，也就是前端要显示页面，需要配置，不过前后端分离就不需要了

	//InstallPlugs(Router)
	// end 插件描述

	address := fmt.Sprintf(":%d", global.DEMO_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	//global.DEMO_LOG.Debug("server run success on ", address)

	fmt.Printf(`欢迎使用 Gin-Vue-Admin
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
`, address)
	global.DEMO_LOG.Error(s.ListenAndServe())
}
