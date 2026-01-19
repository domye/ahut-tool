package services

import (
	"encoding/json"
	"fmt"
)

// OAuthTokenResponse 完整的OAuth令牌响应结构
type OAuthTokenResponse struct {
	AccessToken   string      `json:"access_token"`
	TokenType     string      `json:"token_type"`
	RefreshToken  string      `json:"refresh_token"`
	ExpiresIn     int         `json:"expires_in"`
	Scope         string      `json:"scope"`
	PassWordLevel int         `json:"passWordLevel"`
	AvatarUrl     string      `json:"avatarUrl"`
	AccountType   int         `json:"accountType"`
	UserName      string      `json:"userName"`
	RoleType      string      `json:"roleType"`
	UserId        string      `json:"userId"`
	LastLoginTime string      `json:"lastLoginTime"`
	OauthId       interface{} `json:"oauthId"`
	AccountNo     string      `json:"accountNo"`
	TenantId      string      `json:"tenantId"`
	RoleName      string      `json:"roleName"`
	UserType      int         `json:"userType"`
	Detail        struct {
		SysAuthType         string `json:"sysAuthType"`
		IsSysUserSecondAuth bool   `json:"isSysUserSecondAuth"`
	} `json:"detail"`
	SchoolName string `json:"schoolName"`
	Jti        string `json:"jti"`
}

// reqToken 请求token
func (s *Service) reqToken() (OAuthTokenResponse, error) {
	var tokenResp OAuthTokenResponse

	resp, err := s.client.R().
		SetHeader("tenant-id", "000000").
		SetHeader("authorization", "Basic Zmx5c291cmNlX3dpc2VfYXBwOkRBNzg4YXNkVURqbmFzZF9mbHlzb3VyY2VfZHNkYWREQUlVaXV3cWU=").
		SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36").
		SetHeader("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8").
		Post("https://xskq.ahut.edu.cn/api/flySource-auth/oauth/token?tenantId=000000&username=249074506&password=15e2a01f8599cd9d21a8f9f22e193e84&type=account&grant_type=password&scope=all")

	if err != nil {
		return tokenResp, err
	}

	// 检查响应状态码
	if resp.StatusCode() >= 400 {
		return tokenResp, fmt.Errorf("请求失败，状态码: %d", resp.StatusCode())
	}

	err = json.Unmarshal(resp.Body(), &tokenResp)
	if err != nil {
		return tokenResp, err
	}

	return tokenResp, nil
}
