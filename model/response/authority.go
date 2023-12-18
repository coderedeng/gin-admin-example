package response

import (
	"ginProject/model"
)

type SysAuthorityResponse struct {
	Authority model.SysAuthority `json:"authority"`
}

type SysAuthorityCopyResponse struct {
	Authority      model.SysAuthority `json:"authority"`
	OldAuthorityId uint               `json:"oldAuthorityId"` // 旧角色ID
}
