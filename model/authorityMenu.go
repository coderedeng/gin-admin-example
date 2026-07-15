package model

type SysMenu struct {
	SysBaseMenu
	MenuId      string                 `json:"menuId" gorm:"comment:菜单ID"`
	AuthorityId uint                   `json:"-" gorm:"comment:角色ID"`
	Children    []SysMenu              `json:"children" gorm:"-"`
	Parameters  []SysBaseMenuParameter `json:"parameters" gorm:"foreignKey:SysBaseMenuID;references:MenuId"`
	Btns        map[string]uint        `json:"btns" gorm:"-"`
}

type SysAuthorityMenu struct {
	MenuId      string `json:"menuId" gorm:"comment:菜单ID;column:baseMenuId"`
	AuthorityId string `json:"-" gorm:"comment:角色ID;column:authorityAuthorityId"`
}

func (s SysAuthorityMenu) TableName() string {
	return "authority_menus"
}
