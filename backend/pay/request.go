package pay

import (
	"fmt"

	"github.com/go-resty/resty/v2"

	"ahut-tool/backend/utils"
)

func (s *Service) sendLoginRequest(username, password string) (int, error) {
	// 使用utils包中的RSA加密函数加密密码
	encryptedPwd, err := utils.EncryptPasswordWithRSA(password)
	if err != nil {
		return 0, err
	}

	s.client.SetRedirectPolicy(resty.NoRedirectPolicy())

	resp, err := s.client.R().
		SetHeader("Referer", "https://pay.ahut.edu.cn/Account/Login").
		SetFormData(map[string]string{
			"username": username,
			"pwd":      encryptedPwd,
		}).
		Post("https://pay.ahut.edu.cn/Account/LoginService")

	if err != nil {
		return 0, err
	}

	var cookie string
	for _, i := range resp.Cookies() {
		cookie += fmt.Sprintf("%s=%s; ", i.Name, i.Value)
	}
	// 设置全局cookie
	s.cookie = cookie
	// 将cookie设置到client的header中
	s.client.SetHeader("Cookie", cookie)

	fmt.Println(resp.String())
	return resp.StatusCode(), nil
}
