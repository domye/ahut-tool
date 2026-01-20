package pay

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

func (s *Service) sendLoginRequest(formData map[string]string) (int, error) {
	s.client.SetRedirectPolicy(resty.NoRedirectPolicy())

	resp, err := s.client.R().
		SetHeader("Referer", "https://pay.ahut.edu.cn/Account/Login").
		SetFormData(formData).
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

	return resp.StatusCode(), nil
}
