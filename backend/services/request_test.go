package services

import (
	"testing"

	"github.com/go-resty/resty/v2"
)

func TestReqToken(t *testing.T) {
	service := &Service{
		client: resty.New(),
	}
	a := service.getCookie()
	//service.login("249074506", "3tmwjdaW", a)
	service.login("249074506", "3tmwjdaW", a)
}
