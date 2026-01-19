package services

type LoginResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Name   string `json:"name"`
		Cookie string `json:"cookie"`
	}
}

// LoginResult 包含完整登录结果信息
type LoginResult struct {
	OAuthTokenResponse
}

func (s *Service) Login(userId int, password string) (*LoginResult, error) {
	// TODO: 实现登录逻辑
	return nil, nil
}
