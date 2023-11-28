package response

import "ginProject/model"

type UserResponse struct {
	User model.SysUser `json:"user"`
}
