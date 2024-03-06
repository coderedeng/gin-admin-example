package core

import (
	"fmt"
	"github.com/coderedeng/gin-admin-example/global"
	"github.com/coderedeng/gin-admin-example/initialize"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	// 1. 加载配置
	global.GVA_VP = Viper()
	// 2.初始化日志
	global.GVA_LOG = Zap()
	zap.ReplaceGlobals(global.GVA_LOG)
	// 3.初始化PgSQL连接
	global.GVA_DB = initialize.Pgsql()
	// 4.初始化redis连接
	global.GVA_REDIS = initialize.Redis()
	// 5. 初始化gorm框架
	if global.GVA_DB != nil {
		initialize.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}

	// 6.注册路由
	Router := initialize.Routers()

	// 初始化本地缓存
	initialize.OtherInit()

	// 7.启动服务（优雅关机）
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Port)
	fmt.Printf(`
	swagger后端文档地址:http://127.0.0.1%s/swagger/index.html
	knife4g后端文档地址:http://127.0.0.1%s/doc/index
	前端项目运行地址:http://127.0.0.1:8080
`, address, address)

	initServer(address, Router)

}
