package response

import (
	"github.com/coderedeng/gin-admin-example/model/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
