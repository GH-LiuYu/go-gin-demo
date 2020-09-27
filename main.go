package main

import (
	"demo/core"
	"demo/global"
	"demo/initialize"
)

func main() {

	switch global.DEMO_CONFIG.System.DbType{
	case "mysql":
		initialize.Mysql()
	default:
		initialize.Mysql()
	}
	initialize.DBTables()//注冊迁移数据库
	////程序结束前关闭数据库连接
	defer global.DEMO_DB.Close()

	core.RunWindowsServer()//启动服务



	//以下为简单测试 上面为完整启动流程

	//router := gin.Default()
	//
	//// 匹配的url格式:  /welcome?firstname=Jane&lastname=Doe
	//router.GET("/welcome", func(c *gin.Context) {
	//	firstname := c.DefaultQuery("firstname", "Guest")
	//	lastname := c.Query("lastname") // 是 c.Request.URL.Query().Get("lastname") 的简写
	//
	//	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	//})
	//router.Run(":8888")


}
