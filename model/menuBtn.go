package model

import "github.com/coderedeng/gin-admin-example/global"

type SysBaseMenuBtn struct {
	global.GPA_MODEL
	Name          string `json:"name" gorm:"comment:按钮关键key"`
	Desc          string `json:"desc" gorm:"按钮备注"`
	SysBaseMenuID uint   `json:"sysBaseMenuID" gorm:"comment:菜单ID"`
}

func (SysBaseMenuBtn) TableName() string {
	return "base_menu_btn"
}
