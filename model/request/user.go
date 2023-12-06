package request

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
