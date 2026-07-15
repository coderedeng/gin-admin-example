package initialize

import (
	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/coderedeng/gin-admin-example/global"
	"github.com/coderedeng/gin-admin-example/model"
	"go.uber.org/zap"
	"os"
)

// RegisterTables 注册数据库表专用
func RegisterTables() {
	db := global.GPA_DB
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
		adapter.CasbinRule{},
	)

	if err != nil {
		global.GPA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}

	global.GPA_LOG.Info("register table success")
}
