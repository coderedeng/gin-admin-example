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

// Register
// @description: 用户注册
// @param: u *model.SysUser
// @return: userInter model.SysUser, err error
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

// Login
// @description: 用户登录
// @param: u *model.SysUser
// @return: userInter *model.SysUser, err error
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

// GetUserList
// @description: 获取用户列表
// @param: info request.PageInfo
// @return: list interface{}, total int64, err error
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

// ChangePassWord
// @description: 修改用户密码
// @param: u *model.SysUser, newPassword string
// @return: user model.SysUser, err error
func (userService *UserService) ChangePassWord(u *model.SysUser, newPassword string) (userInter *model.SysUser, err error) {
	var user model.SysUser
	if err = global.GVA_DB.Where("id = ?", u.ID).First(&user).Error; err != nil {
		return nil, err
	}
	if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
		return nil, errors.New("原密码错误")
	}
	user.Password = utils.BcryptHash(newPassword)
	err = global.GVA_DB.Save(&user).Error
	return &user, err
}

// GetUserInfo
// @description: 获取用户信息
// @param: uuid uuid.UUID
// @return: model.SysUser, user err error
func (userService *UserService) GetUserInfo(uuid uuid.UUID) (user model.SysUser, err error) {
	err = global.GVA_DB.First(&user, "uuid = ?", uuid).Error
	if err != nil {
		return user, err
	}
	//MenuServiceApp.UserAuthorityDefaultRouter(&user)
	return user, err
}

// DeleteUser
// @description: 根据用户id删除用户
// @param: id uint
// @return: err error
func (userService *UserService) DeleteUser(id uint) (err error) {
	var user model.SysUser
	err = global.GVA_DB.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	return err
}

// ResetPassword
// @description: 根据用户id重置密码
// @param: id uint
// @return: err error
func (userService *UserService) ResetPassword(id uint) (err error) {
	err = global.GVA_DB.Model(&model.SysUser{}).Where("id = ?", id).Update("password", utils.BcryptHash("123456")).Error
	return err
}
