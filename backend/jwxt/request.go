package jwxt

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

func (s *Service) sendLoginRequest(formData map[string]string) (int, error) {

	s.client.SetRedirectPolicy(resty.NoRedirectPolicy())
	resp, err := s.client.R().
		SetFormData(formData).
		Post("http://jwxt.ahut.edu.cn/jsxsd/xk/LoginToXk")
	var cookie string
	for _, i := range resp.Cookies() {
		cookie += fmt.Sprintf("%s=%s; ", i.Name, i.Value)
	}
	// 设置全局cookie
	s.cookie = cookie
	// 将cookie设置到client的header中
	s.client.SetHeader("Cookie", cookie)
	println(resp.StatusCode())
	return resp.StatusCode(), err
}

// sendGetGradesRequest 获取成绩信息
func (s *Service) sendGetGradesRequest(formData map[string]string) (string, error) {
	resp, err := s.client.R().
		SetFormData(formData).
		Post("http://jwxt.ahut.edu.cn/jsxsd/kscj/cjcx_list")

	if err != nil {
		return "", err
	}

	return resp.String(), nil
}
