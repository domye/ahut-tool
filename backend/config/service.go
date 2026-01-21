package config

import (
	"ahut-tool/backend/models"
	"fmt"
)

// Service 配置服务，提供Wails可调用的方法
type Service struct{}

// NewService 创建新的配置服务实例
func NewService() *Service {
	return &Service{}
}

// GetJwxtCredentials 获取教务系统账号密码
func (s *Service) GetJwxtCredentials() (*models.JwxtCredentials, error) {
	config, err := LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %v", err)
	}

	user, password := config.GetJwxtCredentials()
	return &models.JwxtCredentials{
		User:     user,
		Password: password,
	}, nil
}

// GetPayCredentials 获取缴费系统账号密码
func (s *Service) GetPayCredentials() (*models.PayCredentials, error) {
	config, err := LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %v", err)
	}

	user, password := config.GetPayCredentials()
	return &models.PayCredentials{
		User:     user,
		Password: password,
	}, nil
}
