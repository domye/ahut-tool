// backend/pay/service.go
// 缴费系统服务，管理 HTTP 客户端和会话状态

package pay

import (
	"github.com/go-resty/resty/v2"
)

type Service struct {
	client *resty.Client
	cookie string
}

func NewService() *Service {
	return &Service{
		client: resty.New(),
	}
}
