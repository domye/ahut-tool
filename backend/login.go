package backend

import (
	"ahut-tool/backend/services"
	"fmt"
)

var Instance *services.Service

func (a *App) Login(userId string, password string) string {
	formdate := Instance.GetLoginFormData(userId, password)
	LoginResponse := Instance.GetToken(formdate)
	fmt.Println(LoginResponse)
	return LoginResponse
}
