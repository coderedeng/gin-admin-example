package router

import (
	"ginProject/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUp() *gin.Engine {
	r := gin.New()
	r.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	return r
}
