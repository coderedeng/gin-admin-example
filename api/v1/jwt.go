package v1

import (
	"ginProject/global"
	"ginProject/model"
	"ginProject/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type JwtApi struct{}

// JsonInBlacklist
// @Tags      Jwt
// @Summary   jwt加入黑名单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{msg=string}  "jwt加入黑名单"
// @Router    /v1/jwt/jsonInBlacklist [post]
func (j *JwtApi) JsonInBlacklist(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	jwt := model.JwtBlacklist{Jwt: token}
	err := jwtService.JsonInBlacklist(jwt)
	if err != nil {
		global.GVA_LOG.Error("jwt作废失败!", zap.Error(err))
		response.FailWithMessage("jwt作废失败", c)
		return
	}
	response.OkWithMessage("jwt作废成功", c)
}