package initialize

import (
	docs "ginProject/docs"
	"ginProject/global"
	"ginProject/router"
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
	{
		router.RouterGroupApp.InitUserRouter(PrivateGroup) // 注册用户路由
	}
	return Router
}