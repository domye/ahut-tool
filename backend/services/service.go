package services

import (
	"github.com/go-resty/resty/v2"
)

type Service struct {
	client *resty.Client
	cookie string
}
