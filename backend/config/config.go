package config

import (
	"ahut-tool/backend/models"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

// Config 应用程序配置
type Config struct {
	Jwxt models.JWXTConfig `yaml:"jwxt"`
	Pay  models.PayConfig  `yaml:"pay"`
}

// LoadConfig 从config.yaml文件加载配置
func LoadConfig() (*Config, error) {
	// 获取可执行文件所在目录
	exeDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return nil, err
	}

	// 构造config.yaml文件路径
	configPath := filepath.Join(exeDir, "config.yaml")

	// 读取配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	// 解析YAML配置
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// GetJwxtCredentials 获取教务系统账号密码
func (c *Config) GetJwxtCredentials() (string, string) {
	return c.Jwxt.User, c.Jwxt.Password
}

// GetPayCredentials 获取缴费系统账号密码
func (c *Config) GetPayCredentials() (string, string) {
	return c.Pay.User, c.Pay.Password
}
