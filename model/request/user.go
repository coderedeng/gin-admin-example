package request

import "ginProject/model"

// UserRegister User register structure
type UserRegister struct {
	Username  string `json:"userName" example:"用户名"`
	Password  string `json:"passWord" example:"密码"`
	NickName  string `json:"nickName" example:"昵称"`
	HeaderImg string `json:"headerImg" example:"头像链接"`
	Enable    int    `json:"enable" swaggertype:"string" example:"int 是否启用"`
	Phone     string `json:"phone" example:"电话号码"`
	Email     string `json:"email" example:"电子邮箱"`
}

// UserLogin User login structure
type UserLogin struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

// ChangePassWord User change password
type ChangePassWord struct {
	ID              uint   `json:"-"` // 从 JWT 中提取 user id，避免越权
	Password        string `json:"passWord" example:"密码"`
	ConfirmPassword string `json:"confirmPassword" example:"确认密码"`
}

// Modify  user's auth structure
type SetUserAuth struct {
	AuthorityId uint `json:"authorityId"` // 角色ID
}

// Modify  user's auth structure
type SetUserAuthorities struct {
	ID           uint
	AuthorityIds []uint `json:"authorityIds"` // 角色ID
}

type ChangeUserInfo struct {
	ID           uint                 `gorm:"primarykey"`                                                                           // 主键ID
	NickName     string               `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                            // 用户昵称
	Phone        string               `json:"phone"  gorm:"comment:用户手机号"`                                                          // 用户手机号
	AuthorityIds []uint               `json:"authorityIds" gorm:"-"`                                                                // 角色ID
	Email        string               `json:"email"  gorm:"comment:用户邮箱"`                                                           // 用户邮箱
	HeaderImg    string               `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	SideMode     string               `json:"sideMode"  gorm:"comment:用户侧边主题"`                                                      // 用户侧边主题
	Enable       int                  `json:"enable" gorm:"comment:冻结用户"`                                                           //冻结用户
	Authorities  []model.SysAuthority `json:"-" gorm:"many2many:userAuthority;"`
}
