package router

import (
	v1 "ginProject/api/v1"
	"github.com/gin-gonic/gin"
)

type CasbinRouter struct{}

func (s *CasbinRouter) InitCasbinRouter(Router *gin.RouterGroup) {
	casbinRouter := Router.Group("casbin")
	casbinRouterWithoutRecord := Router.Group("casbin")
	casbinApi := v1.ApiGroupApp.CasbinApi
	{
		casbinRouter.POST("updateCasbin", casbinApi.UpdateCasbin)
	}
	{
		casbinRouterWithoutRecord.POST("getPolicyPathByAuthorityId", casbinApi.GetPolicyPathByAuthorityId)
	}
}
