package response

import "github.com/coderedeng/gin-admin-example/model"

type SysMenusResponse struct {
	Menus []model.SysMenu `json:"menus"`
}

type SysBaseMenusResponse struct {
	Menus []model.SysBaseMenu `json:"menus"`
}

type SysBaseMenuResponse struct {
	Menu model.SysBaseMenu `json:"menu"`
}
