package request

import (
	"github.com/coderedeng/gin-admin-example/global"
	"github.com/coderedeng/gin-admin-example/model"
)

// Add menu authority info structure
type AddMenuAuthorityInfo struct {
	Menus       []model.SysBaseMenu `json:"menus"`
	AuthorityId uint                `json:"authorityId"` // 角色ID
}

func DefaultMenu() []model.SysBaseMenu {
	return []model.SysBaseMenu{{
		GVA_MODEL: global.GVA_MODEL{ID: 1},
		ParentId:  "0",
		Path:      "dashboard",
		Name:      "dashboard",
		Component: "view/dashboard/index.vue",
		Sort:      1,
		Meta: model.Meta{
			Title: "仪表盘",
			Icon:  "setting",
		},
	}}
}
