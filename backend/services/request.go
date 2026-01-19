package services

import (
	"fmt"
)

func (s *Service) getCookie() string {
	resp, err := s.client.R().Post("http://jwxt.ahut.edu.cn/jsxsd/")
	if err != nil {
		fmt.Println(err)
	}
	var cookie string
	for _, i := range resp.Cookies() {
		cookie += fmt.Sprintf("%s=%s; ", i.Name, i.Value)
	}
	println(cookie)
	return cookie
}
