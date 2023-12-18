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
	"time"
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

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserAuthority
//@description: 设置一个用户的权限
//@param: uuid uuid.UUID, authorityId string
//@return: err error

func (userService *UserService) SetUserAuthority(id uint, authorityId uint) (err error) {
	assignErr := global.GVA_DB.Where("sys_user_id = ? AND sys_authority_authority_id = ?", id, authorityId).First(&model.SysUserAuthority{}).Error
	if errors.Is(assignErr, gorm.ErrRecordNotFound) {
		return errors.New("该用户无此角色")
	}
	err = global.GVA_DB.Where("id = ?", id).First(&model.SysUser{}).Update("authority_id", authorityId).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserAuthorities
//@description: 设置一个用户的权限
//@param: id uint, authorityIds []string
//@return: err error

func (userService *UserService) SetUserAuthorities(id uint, authorityIds []uint) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		TxErr := tx.Delete(&[]model.SysUserAuthority{}, "sys_user_id = ?", id).Error
		if TxErr != nil {
			return TxErr
		}
		var useAuthority []model.SysUserAuthority
		for _, v := range authorityIds {
			useAuthority = append(useAuthority, model.SysUserAuthority{
				SysUserId: id, SysAuthorityAuthorityId: v,
			})
		}
		TxErr = tx.Create(&useAuthority).Error
		if TxErr != nil {
			return TxErr
		}
		TxErr = tx.Where("id = ?", id).First(&model.SysUser{}).Update("authority_id", authorityIds[0]).Error
		if TxErr != nil {
			return TxErr
		}
		// 返回 nil 提交事务
		return nil
	})
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserInfo
//@description: 设置用户信息
//@param: reqUser model.SysUser
//@return: err error, user model.SysUser

func (userService *UserService) SetUserInfo(req model.SysUser) error {
	return global.GVA_DB.Model(&model.SysUser{}).
		Select("updated_at", "nick_name", "header_img", "phone", "email", "sideMode", "enable").
		Where("id=?", req.ID).
		Updates(map[string]interface{}{
			"updated_at": time.Now(),
			"nick_name":  req.NickName,
			"header_img": req.HeaderImg,
			"phone":      req.Phone,
			"email":      req.Email,
			"side_mode":  req.SideMode,
			"enable":     req.Enable,
		}).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserInfo
//@description: 设置用户信息
//@param: reqUser model.SysUser
//@return: err error, user model.SysUser

func (userService *UserService) SetSelfInfo(req model.SysUser) error {
	return global.GVA_DB.Model(&model.SysUser{}).
		Where("id=?", req.ID).
		Updates(req).Error
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 通过id获取用户信息
//@param: id int
//@return: err error, user *model.SysUser

func (userService *UserService) FindUserById(id int) (user *model.SysUser, err error) {
	var u model.SysUser
	err = global.GVA_DB.Where("`id` = ?", id).First(&u).Error
	return &u, err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserByUuid
//@description: 通过uuid获取用户信息
//@param: uuid string
//@return: err error, user *model.SysUser

func (userService *UserService) FindUserByUuid(uuid string) (user *model.SysUser, err error) {
	var u model.SysUser
	if err = global.GVA_DB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return &u, errors.New("用户不存在")
	}
	return &u, nil
}
