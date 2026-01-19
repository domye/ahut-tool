package services

import (
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"
)

func TestReqToken(t *testing.T) {
	service := &Service{
		client: resty.New(),
	}
	_, err := service.reqToken()
	if err != nil {
		fmt.Printf("%#v \n", err)
	}
	//fmt.Printf("%#v \n", result)
}
