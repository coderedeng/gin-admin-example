package model

import (
	"github.com/coderedeng/gin-admin-example/global"
)

type JwtBlacklist struct {
	global.GPA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
