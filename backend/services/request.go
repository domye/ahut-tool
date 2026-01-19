package services

import (
	"ahut-tool/backend/utils"
	"fmt"
)

func (s *Service) getCookie() string {
	resp, err := s.client.R().Post("http://jwxt.ahut.edu.cn/jsxsd/xk/LoginToXk?userAccount=249074506&userPassword&encoded=MjQ5MDc0NTA2%%%M3Rtd2pkYVc=&pwdstr1&pwdstr2")
	if err != nil {
		fmt.Println(err)
	}
	var cookie string
	for _, i := range resp.Cookies() {
		cookie += fmt.Sprintf("%s=%s; ", i.Name, i.Value)
	}

	return cookie
}

func (s *Service) login(userId string, password string, cookie string) error {
	encoded := utils.Base64Encode(userId) + "%25%25%25" + utils.Base64Encode(password) + "%3D"
	println(encoded)
	s.client.SetFormData(map[string]string{"xnxq01id": "2025-2026-2",
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36"})
	response, err := s.client.R().Post("http://jwxt.ahut.edu.cn/jsxsd/xskb/xskb_list.do")
	println(response.String())
	return err
}
