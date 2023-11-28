package v1

import (
	"ginProject/model/common/response"
	req "ginProject/model/request"
	"github.com/gin-gonic/gin"
)

type userApi struct{}

// Register
// @Tags     SysUser
// @Summary  用户注册账号
// @Produce   application/json
// @Param    data  body      systemReq.Register                                            true  "用户名, 昵称, 密码, 角色ID"
// @Success  200   {object}  response.Response{data=systemRes.SysUserResponse,msg=string}  "用户注册账号,返回包括用户信息"
// @Router   /user/admin_register [post]
func (u *userApi) Register(c *gin.Context) {
	var r req.UserRegister
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//err = utils.Verify(r, utils.RegisterVerify)
	//if err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	//var authorities []system.SysAuthority
	//for _, v := range r.AuthorityIds {
	//	authorities = append(authorities, system.SysAuthority{
	//		AuthorityId: v,
	//	})
	//}
	//user := &model.SysUser{Username: r.Username, NickName: r.NickName, Password: r.Password, HeaderImg: r.HeaderImg, AuthorityId: r.AuthorityId, Authorities: authorities, Enable: r.Enable, Phone: r.Phone, Email: r.Email}
	//userReturn, err := userService.Register(*user)
	//if err != nil {
	//	global.GVA_LOG.Error("注册失败!", zap.Error(err))
	//	response.FailWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册失败", c)
	//	return
	//}
	//response.OkWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册成功", c)
}
