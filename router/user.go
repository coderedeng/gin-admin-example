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
		userRouter.POST("register", userApi.Register) // 用户注册账号
		userRouter.POST("login", userApi.Login)       // 用户登录
		userRouter.POST("captcha", userApi.Captcha)   // 获取验证码
	}
	{
		jwtUserRouter.POST("GetUserList", userApi.GetUserList)       // 获取系统用户列表
		jwtUserRouter.POST("ChangePassWord", userApi.ChangePassWord) // 用户修改密码
		jwtUserRouter.GET("GetUserInfo", userApi.GetUserInfo)        // 获取当前用户信息
		jwtUserRouter.GET("ResetPassword", userApi.ResetPassword)    // 根据用户id重置密码
		jwtUserRouter.DELETE("DeleteUser", userApi.DeleteUser)       // 根据用户id删除用户
	}
}
