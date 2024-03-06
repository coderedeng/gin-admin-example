package initialize

import (
	docs "github.com/coderedeng/gin-admin-example/docs"
	"github.com/coderedeng/gin-admin-example/global"
	"github.com/coderedeng/gin-admin-example/router"
	"github.com/gin-gonic/gin"
	"github.com/hononet639/knife4g"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	// swagger文档地址 生成命令 swag init
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")

	// knife4g文档地址 生成命令 swag init
	Router.GET("/doc/*any", knife4g.Handler(knife4g.Config{RelativePath: "/doc"}))
	global.GVA_LOG.Info("register knife4g handler")

	PrivateGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	//PrivateGroup.Use(middleware.JWTAuth())
	{
		router.RouterGroupApp.InitUserRouter(PrivateGroup)               // 初始化用户路由
		router.RouterGroupApp.InitApiRouter(PrivateGroup)                // 初始化Api路由
		router.RouterGroupApp.InitAuthorityRouter(PrivateGroup)          // 初始化用户角色路由
		router.RouterGroupApp.InitAuthorityBtnRouterRouter(PrivateGroup) // 初始化用户角色按钮路由
		router.RouterGroupApp.InitCasbinRouter(PrivateGroup)             // 初始化Casbin路由
		router.RouterGroupApp.InitJwtRouter(PrivateGroup)                // 初始化JWT路由
		router.RouterGroupApp.InitMenuRouter(PrivateGroup)               // 初始化菜单路由
	}
	return Router

}
