package services

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

func (s *Service) getCookie(formData map[string]string) string {

	s.client.SetRedirectPolicy(resty.NoRedirectPolicy())
	resp, err := s.client.R().
		SetFormData(formData).
		Post("http://jwxt.ahut.edu.cn/jsxsd/xk/LoginToXk")
	if err != nil {
		fmt.Println(err)
	}
	var cookie string
	for _, i := range resp.Cookies() {
		cookie += fmt.Sprintf("%s=%s; ", i.Name, i.Value)
	}
	println(resp.StatusCode())
	return cookie
}
