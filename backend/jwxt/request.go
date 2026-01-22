package jwxt

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

func (s *Service) sendLoginRequest(formData map[string]string) (int, error) {

	s.client.SetRedirectPolicy(resty.NoRedirectPolicy())
	resp, err := s.client.R().
		SetFormData(formData).
		Post(LoginURL)
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
		Post(GradesURL)

	if err != nil {
		return "", err
	}

	return resp.String(), nil
}

// sendGetClassesRequest 获取课程信息
func (s *Service) sendGetClassesRequest(formData map[string]string) (string, error) {
	resp, err := s.client.R().
		SetFormData(formData).
		Post(ClassesURL)

	if err != nil {
		return "", err
	}

	return resp.String(), nil
}
