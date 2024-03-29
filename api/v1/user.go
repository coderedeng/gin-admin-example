package v1

import (
	"github.com/coderedeng/gin-admin-example/global"
	"github.com/coderedeng/gin-admin-example/model"
	"github.com/coderedeng/gin-admin-example/model/common/request"
	"github.com/coderedeng/gin-admin-example/model/common/response"
	req "github.com/coderedeng/gin-admin-example/model/request"
	res "github.com/coderedeng/gin-admin-example/model/response"
	"github.com/coderedeng/gin-admin-example/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type UserApi struct{}

// 当开启多服务器部署时，替换下面的配置，使用redis共享存储验证码
// var store = captcha.NewDefaultRedisStore()
var store = base64Captcha.DefaultMemStore

// Register
// @Tags     User
// @Summary  用户注册账号
// @Produce   application/json
// @Param    data  body      req.UserRegister                                           true  "用户名, 昵称, 密码, 角色ID"
// @Success  200   {object}  response.Response{data=res.UserResponse,msg=string}  "用户注册账号,返回包括用户信息"
// @Router   /v1/user/register [post]
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
		global.GPA_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithDetailed(res.UserResponse{User: userReturn}, "注册失败", c)
		return
	}
	response.OkWithDetailed(res.UserResponse{User: userReturn}, "注册成功", c)
}

// Login
// @Tags     User
// @Summary  用户登录
// @Produce   application/json
// @Param    data  body      req.UserLogin                                        true  "用户名, 密码, 验证码"
// @Success  200   {object}  response.Response{data=res.UserResponse,msg=string}  "返回包括用户信息,token,过期时间"
// @Router   /v1/user/login [post]
func (u *UserApi) Login(c *gin.Context) {
	var l req.UserLogin
	err := c.ShouldBindJSON(&l)
	key := c.ClientIP()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 判断验证码是否开启
	openCaptcha := global.GPA_CONFIG.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := global.GPA_CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}

	var oc bool = openCaptcha == 0 || openCaptcha < interfaceToInt(v)

	if !oc || store.Verify(l.CaptchaId, l.Captcha, true) {
		sysUser := &model.SysUser{Username: l.Username, Password: l.Password}
		user, err := userService.Login(sysUser)
		if err != nil {
			global.GPA_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
			// 验证码次数+1
			global.BlackCache.Increment(key, 1)
			response.FailWithMessage("用户名不存在或者密码错误", c)
			return
		}
		if user.Enable != 1 {
			global.GPA_LOG.Error("登陆失败! 用户被禁止登录!")
			// 验证码次数+1
			global.BlackCache.Increment(key, 1)
			response.FailWithMessage("用户被禁止登录", c)
			return
		}
		u.TokenNext(c, *user)
		return
	}
	// 验证码次数+1
	global.BlackCache.Increment(key, 1)
	response.FailWithMessage("验证码错误", c)
}

// TokenNext 登录以后签发jwt
func (u *UserApi) TokenNext(c *gin.Context, user model.SysUser) {
	j := &utils.JWT{SigningKey: []byte(global.GPA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(req.BaseClaims{
		UUID:     user.UUID,
		ID:       user.ID,
		NickName: user.NickName,
		Username: user.Username,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GPA_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !global.GPA_CONFIG.System.UseMultipoint {
		response.OkWithDetailed(res.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
		return
	}

	if jwtStr, err := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			global.GPA_LOG.Error("设置登录状态失败!", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(res.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.GPA_LOG.Error("设置登录状态失败!", zap.Error(err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT model.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(res.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
	}
}

// Captcha
// @Tags      User
// @Summary   生成验证码
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=res.SysCaptchaResponse,msg=string}  "生成验证码,返回包括随机数id,base64,验证码长度,是否开启验证码"
// @Router    /v1/user/captcha [post]
func (u *UserApi) Captcha(c *gin.Context) {
	// 判断验证码是否开启
	openCaptcha := global.GPA_CONFIG.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := global.GPA_CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	key := c.ClientIP()
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}

	var oc bool
	if openCaptcha == 0 || openCaptcha < interfaceToInt(v) {
		oc = true
	}
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.GPA_CONFIG.Captcha.ImgHeight, global.GPA_CONFIG.Captcha.ImgWidth, global.GPA_CONFIG.Captcha.KeyLong, 0.7, 80)
	//cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))   // v8下使用redis
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		global.GPA_LOG.Error("验证码获取失败!", zap.Error(err))
		response.FailWithMessage("验证码获取失败", c)
		return
	}
	response.OkWithDetailed(res.SysCaptchaResponse{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: global.GPA_CONFIG.Captcha.KeyLong,
		OpenCaptcha:   oc,
	}, "验证码获取成功", c)
}

// 类型转换
func interfaceToInt(v interface{}) (i int) {
	switch v := v.(type) {
	case int:
		i = v
	default:
		i = 0
	}
	return
}

// GetUserList
// @Tags      User
// @Summary   获取系统用户列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.PageInfo true  "页码, 每页大小"
// @Success   200  {object}  response.Response{data=response.PageResult,msg=string}  "分页获取用户列表,返回包括列表,总数,页码,每页数量"
// @Router    /v1/user/GetUserList [post]
func (u *UserApi) GetUserList(c *gin.Context) {
	var pageInfo request.PageInfo

	err := c.ShouldBindJSON(&pageInfo)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := userService.GetUserList(pageInfo)

	if err != nil {
		global.GPA_LOG.Error("获取系统用户列表失败!", zap.Error(err))
		response.FailWithMessage("获取系统用户列表失败!", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// ChangePassWord
// @Tags      User
// @Summary   用户修改密码
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.ChangePassWord true  "账号, 密码, 确认密码"
// @Success   200  {object}  response.Response{msg=string}  "用户修改密码,返回包括修改成功，失败"
// @Router    /v1/user/ChangePassWord [post]
func (u *UserApi) ChangePassWord(c *gin.Context) {
	var p req.ChangePassWord

	err := c.ShouldBindJSON(&p)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := utils.GetUserID(c)
	user := &model.SysUser{GPA_MODEL: global.GPA_MODEL{ID: uid}, Password: p.Password}
	_, err = userService.ChangePassWord(user, p.ConfirmPassword)
	if err != nil {
		global.GPA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败，原密码与当前账户不符", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}

// GetUserInfo
// @Tags      User
// @Summary   获取用户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=map[string]interface{},msg=string}  "获取用户信息"
// @Router    /v1/user/GetUserInfo [get]
func (u *UserApi) GetUserInfo(c *gin.Context) {
	uuid := utils.GetUserUuid(c)
	userInfo, err := userService.GetUserInfo(uuid)
	if err != nil {
		global.GPA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(userInfo, "获取成功!", c)
}

// DeleteUser
// @Tags      User
// @Summary   删除用户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     id  query  uint                true  "用户ID"
// @Success   200   {object}  response.Response{msg=string}  "删除用户"
// @Router    /v1/user/DeleteUser [delete]
func (u *UserApi) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Query("id"), 10, 64)
	if err != nil {
		response.FailWithMessage("用户id不能为空", c)
		return
	}
	userId := uint(id)

	jwtId := utils.GetUserID(c)
	if jwtId == userId {
		response.FailWithMessage("删除失败，不能删除自己", c)
		return
	}

	err = userService.DeleteUser(userId)
	if err != nil {
		global.GPA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// ResetPassword
// @Tags      User
// @Summary   重置用户密码
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     id  query  uint                true  "用户ID"
// @Success   200   {object}  response.Response{msg=string}  "重置用户密码"
// @Router    /v1/user/ResetPassword [get]
func (u *UserApi) ResetPassword(c *gin.Context) {
	id, err := strconv.ParseUint(c.Query("id"), 10, 64)
	if err != nil {
		response.FailWithMessage("用户id不能为空", c)
		return
	}
	userId := uint(id)
	err = userService.ResetPassword(userId)
	if err != nil {
		global.GPA_LOG.Error("重置失败!", zap.Error(err))
		response.FailWithMessage("重置失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("重置成功", c)
}

// SetUserAuthority
// @Tags      User
// @Summary   更改用户权限
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      req.SetUserAuth          true  "用户UUID, 角色ID"
// @Success   200   {object}  response.Response{msg=string}  "设置用户权限"
// @Router    /v1/user/setUserAuthority [post]
func (u *UserApi) SetUserAuthority(c *gin.Context) {
	var sua req.SetUserAuth
	err := c.ShouldBindJSON(&sua)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	userID := utils.GetUserID(c)
	err = userService.SetUserAuthority(userID, sua.AuthorityId)
	if err != nil {
		global.GPA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	claims := utils.GetUserInfo(c)
	j := &utils.JWT{SigningKey: []byte(global.GPA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims.AuthorityId = sua.AuthorityId
	if token, err := j.CreateToken(*claims); err != nil {
		global.GPA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		c.Header("new-token", token)
		c.Header("new-expires-at", strconv.FormatInt(claims.ExpiresAt.Unix(), 10))
		response.OkWithMessage("修改成功", c)
	}
}

// SetUserAuthorities
// @Tags      User
// @Summary   设置用户权限
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      req.SetUserAuthorities   true  "用户UUID, 角色ID"
// @Success   200   {object}  response.Response{msg=string}  "设置用户权限"
// @Router    /v1/user/setUserAuthorities [post]
func (u *UserApi) SetUserAuthorities(c *gin.Context) {
	var sua req.SetUserAuthorities
	err := c.ShouldBindJSON(&sua)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userService.SetUserAuthorities(sua.ID, sua.AuthorityIds)
	if err != nil {
		global.GPA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}

// SetUserInfo
// @Tags      User
// @Summary   设置用户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      model.SysUser                                             true  "ID, 用户名, 昵称, 头像链接"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "设置用户信息"
// @Router    /v1/user/setUserInfo [put]
func (u *UserApi) SetUserInfo(c *gin.Context) {
	var user req.ChangeUserInfo
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if len(user.AuthorityIds) != 0 {
		err = userService.SetUserAuthorities(user.ID, user.AuthorityIds)
		if err != nil {
			global.GPA_LOG.Error("设置失败!", zap.Error(err))
			response.FailWithMessage("设置失败", c)
			return
		}
	}
	err = userService.SetUserInfo(model.SysUser{
		GPA_MODEL: global.GPA_MODEL{
			ID: user.ID,
		},
		NickName:  user.NickName,
		HeaderImg: user.HeaderImg,
		Phone:     user.Phone,
		Email:     user.Email,
		SideMode:  user.SideMode,
		Enable:    user.Enable,
	})
	if err != nil {
		global.GPA_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return
	}
	response.OkWithMessage("设置成功", c)
}

// SetSelfInfo
// @Tags      User
// @Summary   设置用户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      model.SysUser                                             true  "ID, 用户名, 昵称, 头像链接"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "设置用户信息"
// @Router    /v1/user/SetSelfInfo [put]
func (u *UserApi) SetSelfInfo(c *gin.Context) {
	var user req.ChangeUserInfo
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user.ID = utils.GetUserID(c)
	err = userService.SetSelfInfo(model.SysUser{
		GPA_MODEL: global.GPA_MODEL{
			ID: user.ID,
		},
		NickName:  user.NickName,
		HeaderImg: user.HeaderImg,
		Phone:     user.Phone,
		Email:     user.Email,
		SideMode:  user.SideMode,
		Enable:    user.Enable,
	})
	if err != nil {
		global.GPA_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return
	}
	response.OkWithMessage("设置成功", c)
}
