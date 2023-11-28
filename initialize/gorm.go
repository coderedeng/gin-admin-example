package initialize

import (
	"ginProject/global"
	"ginProject/model"
)

// RegisterTables 注册数据库表专用
func RegisterTables() {
	db := global.GVA_DB
	db.AutoMigrate(
		// 系统用户表
		model.SysUser{},
	)
	global.GVA_LOG.Info("register table success")
}
