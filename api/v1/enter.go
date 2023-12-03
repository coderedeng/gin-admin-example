package v1

import (
	"ginProject/service"
)

type ApiGroup struct {
	UserApi
}

var ApiGroupApp = new(ApiGroup)

var (
	userService = service.UserService{}
	jwtService  = service.JwtService{}
)
