package router

import (
	v1 "ginProject/api/v1"
	"ginProject/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	jwtUserRouter := Router.Group("user")
	jwtUserRouter.Use(middleware.JWTAuth())
	userApi := v1.ApiGroupApp.UserApi
	{
		userRouter.POST("register", userApi.Register) // 管理员注册账号
		userRouter.POST("login", userApi.Login)       // 管理员注册账号
		userRouter.POST("captcha", userApi.Captcha)   // 管理员注册账号
	}
	{
		jwtUserRouter.POST("GetUserList", userApi.GetUserList) // 获取系统用户列表
	}
}
