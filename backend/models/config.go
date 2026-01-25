package models

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
	StartDate       string `json:"startDate"`
}

// DormConfig 宿舍配置
type DormConfig struct {
	Campus       string `json:"campus"`
	BuildingId   string `json:"buildingId"`
	BuildingName string `json:"buildingName"`
	RoomId       string `json:"roomId"`
}
