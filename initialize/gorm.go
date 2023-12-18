package initialize

import (
	"ginProject/global"
	"ginProject/model"
	"go.uber.org/zap"
	"os"
)

// RegisterTables 注册数据库表专用
func RegisterTables() {
	db := global.GVA_DB
	err := db.AutoMigrate(
		// 系统模块表
		model.SysApi{},
		model.SysUser{},
		model.SysBaseMenu{},
		model.JwtBlacklist{},
		model.SysAuthority{},
		model.SysBaseMenuParameter{},
		model.SysBaseMenuBtn{},
		model.SysAuthorityBtn{},
	)

	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}

	global.GVA_LOG.Info("register table success")
}
