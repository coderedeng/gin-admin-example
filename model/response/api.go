package response

import (
	"github.com/coderedeng/gin-admin-example/model"
)

type SysAPIResponse struct {
	Api model.SysApi `json:"api"`
}

type SysAPIListResponse struct {
	Apis []model.SysApi `json:"apis"`
}
