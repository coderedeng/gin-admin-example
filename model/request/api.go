package request

import (
	"github.com/coderedeng/gin-admin-example/model"
	"github.com/coderedeng/gin-admin-example/model/common/request"
)

// SearchApiParams api分页条件查询及排序结构体
type SearchApiParams struct {
	model.SysApi
	request.PageInfo
	OrderKey string `json:"orderKey"` // 排序
	Desc     bool   `json:"desc"`     // 排序方式:升序false(默认)|降序true
}
