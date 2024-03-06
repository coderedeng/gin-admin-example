package middleware

import (
	"strconv"
	"strings"

	"github.com/coderedeng/gin-admin-example/global"
	"github.com/coderedeng/gin-admin-example/model/common/response"
	"github.com/coderedeng/gin-admin-example/service"
	"github.com/coderedeng/gin-admin-example/utils"
	"github.com/gin-gonic/gin"
)

var casbinService = service.CasbinService{}

// CasbinHandler 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if global.GPA_CONFIG.System.Mode != "develop" {
			waitUse, _ := utils.GetClaims(c)
			//获取请求的PATH
			path := c.Request.URL.Path
			obj := strings.TrimPrefix(path, global.GPA_CONFIG.System.RouterPrefix)
			// 获取请求方法
			act := c.Request.Method
			// 获取用户的角色
			sub := strconv.Itoa(int(waitUse.AuthorityId))
			e := casbinService.Casbin() // 判断策略中是否存在
			success, _ := e.Enforce(sub, obj, act)
			if !success {
				response.FailWithDetailed(gin.H{}, "权限不足", c)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
