package response

import (
	"github.com/coderedeng/gin-admin-example/model"
)

type SysAuthorityResponse struct {
	Authority model.SysAuthority `json:"authority"`
}

type SysAuthorityCopyResponse struct {
	Authority      model.SysAuthority `json:"authority"`
	OldAuthorityId uint               `json:"oldAuthorityId"` // 旧角色ID
}
