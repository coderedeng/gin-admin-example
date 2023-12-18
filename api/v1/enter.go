package v1

import (
	"ginProject/service"
)

type ApiGroup struct {
	UserApi
	JwtApi
	CasbinApi
	SystemApiApi
	AuthorityApi
	AuthorityMenuApi
	AuthorityBtnApi
}

var ApiGroupApp = new(ApiGroup)

var (
	apiService          = service.ApiService{}
	userService         = service.UserService{}
	jwtService          = service.JwtService{}
	casbinService       = service.CasbinService{}
	menuService         = service.MenuService{}
	baseMenuService     = service.BaseMenuService{}
	authorityService    = service.AuthorityService{}
	authorityBtnService = service.AuthorityBtnService{}
)
