package model

// SysUserAuthority 是 sysUser 和 sysAuthority 的连接表
type SysUserAuthority struct {
	SysUserId               uint `gorm:"column:userId"`
	SysAuthorityAuthorityId uint `gorm:"column:authorityAuthorityId"`
}

func (s *SysUserAuthority) TableName() string {
	return "userAuthority"
}
