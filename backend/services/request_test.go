package services

import (
	"testing"

	"github.com/go-resty/resty/v2"
)

func TestReqToken(t *testing.T) {
	service := &Service{
		client: resty.New(),
	}
	service.getCookie()
}
