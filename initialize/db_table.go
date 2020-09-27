package initialize

import (
	"demo/global"
	"demo/model"
)

// 注册数据库表专用
func DBTables() {
	db := global.DEMO_DB
	db.AutoMigrate(
		model.SysAuthority{},
		model.SysBaseMenu{},
		model.SysBaseMenuParameter{},
	)
	// global.DEMO_LOG.Debug("register table success")
}
