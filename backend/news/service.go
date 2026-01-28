// backend/news/service.go
// 新闻服务，管理 HTTP 客户端和会话状态

package news

import (
	"github.com/go-resty/resty/v2"
)

type Service struct {
	client *resty.Client
}

func NewService() *Service {
	return &Service{
		client: resty.New(),
	}
}
