package services

import (
	"ahut-tool/backend/models"
	"ahut-tool/backend/utils"
	"fmt"
	"github.com/go-resty/resty/v2"
)

func (s *Service) getCookie(formData map[string]string) (string, error) {

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
	return cookie, err
}

// GetGrades 获取成绩信息
func (s *Service) GetGrades(kksj, kcxz, kcmc, xsfs string) ([]models.Grade, *models.GradeSummary, error) {
	formData := map[string]string{
		"kksj": kksj,
		"kcxz": kcxz,
		"kcmc": kcmc,
		"xsfs": xsfs,
	}

	// 打印请求参数和cookie
	fmt.Printf("请求参数: kksj=%s, kcxz=%s, kcmc=%s, xsfs=%s\n", kksj, kcxz, kcmc, xsfs)
	fmt.Printf("Cookie: %s\n", s.cookie)

	resp, err := s.client.R().
		SetFormData(formData).
		Post("http://jwxt.ahut.edu.cn/jsxsd/kscj/cjcx_list")

	if err != nil {
		return nil, nil, err
	}

	// 打印响应状态码和内容
	fmt.Printf("响应状态码: %d\n", resp.StatusCode())
	fmt.Printf("响应内容长度: %d\n", len(resp.String()))

	// 解析HTML获取成绩信息
	grades, summary, err := utils.ParseGrades(resp.String())

	if err != nil {
		return nil, nil, err
	}

	return grades, summary, nil
}
