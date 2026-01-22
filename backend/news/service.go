package news

import (
	"github.com/go-resty/resty/v2"
)

type Service struct {
	client *resty.Client
}

// NewService 创建新的Service实例
func NewService() *Service {
	return &Service{
		client: resty.New(),
	}
}
