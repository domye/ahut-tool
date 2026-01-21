package models

// Config 应用程序配置
type Config struct {
	Jwxt JWXTConfig `yaml:"jwxt"`
	Pay  PayConfig  `yaml:"pay"`
}

// JWXTConfig 教务系统配置
type JWXTConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// PayConfig 缴费系统配置
type PayConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
