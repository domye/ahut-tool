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

// SchedulesConfig 课程表默认节次配置
type SchedulesConfig struct {
	DefaultSchedule string `json:"defaultSchedule"`
}

// DormConfig 宿舍配置
type DormConfig struct {
	Campus     string `json:"campus"`
	BuildingId string `json:"buildingId"`
	RoomId     string `json:"roomId"`
}
