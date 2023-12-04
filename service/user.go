package service

import (
	"errors"
	"fmt"
	"ginProject/global"
	"ginProject/model"
	"ginProject/model/common/request"
	"ginProject/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserService struct{}

func (userService *UserService) Register(u model.SysUser) (userInter model.SysUser, err error) {
	var user model.SysUser
	if !errors.Is(global.GVA_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return userInter, errors.New("用户名已注册")
	}
	// 否则 附加uuid 密码hash加密 注册
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.NewV4()
	err = global.GVA_DB.Create(&u).Error
	return u, err
}

func (userService *UserService) Login(u *model.SysUser) (userInter *model.SysUser, err error) {
	if nil == global.GVA_DB {
		return nil, fmt.Errorf("db not init")
	}

	var user model.SysUser
	err = global.GVA_DB.Where("username = ?", u.Username).First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
	}
	return &user, err
}

func (userService *UserService) GetUserList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.SysUser{})
	var userList []model.SysUser
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&userList).Error
	return userList, total, err
}
