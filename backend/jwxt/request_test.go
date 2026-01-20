package jwxt

import (
	"github.com/go-resty/resty/v2"
	"testing"
)

func TestReqToken(t *testing.T) {
	service := &Service{
		client: resty.New(),
	}
	a := service.GetLoginFormData("249074506", "3tmwjdaW")
	service.sendLoginRequest(a)
	service.GetGrades(service.GetGradesFormData("2025-2026-1", "", "", "all"))
}
