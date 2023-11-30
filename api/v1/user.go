package v1

import (
	"ginProject/global"
	"ginProject/model"
	"ginProject/model/common/response"
	req "ginProject/model/request"
	res "ginProject/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserApi struct{}

// Register
// @Tags     SysUser
// @Summary  用户注册账号
// @Produce   application/json
// @Param    data  body      req.UserRegister                                           true  "用户名, 昵称, 密码, 角色ID"
// @Success  200   {object}  response.Response{data=res.UserResponse,msg=string}  "用户注册账号,返回包括用户信息"
// @Router   /api/user/register [post]
func (u *UserApi) Register(c *gin.Context) {
	var r req.UserRegister
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user := &model.SysUser{Username: r.Username, NickName: r.NickName, Password: r.Password, HeaderImg: r.HeaderImg, Enable: r.Enable, Phone: r.Phone, Email: r.Email}

	userReturn, err := userService.Register(*user)

	if err != nil {
		global.GVA_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithDetailed(res.UserResponse{User: userReturn}, "注册失败", c)
		return
	}
	response.OkWithDetailed(res.UserResponse{User: userReturn}, "注册成功", c)
}
