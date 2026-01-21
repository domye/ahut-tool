package models

// JwxtCredentials 教务系统凭证
type JwxtCredentials struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

// PayCredentials 缴费系统凭证
type PayCredentials struct {
	User     string `json:"user"`
	Password string `json:"password"`
}
