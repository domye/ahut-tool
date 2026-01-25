package backend

import (
	"ahut-tool/backend/models"
)

func (a *App) JwxtLogin(config models.JwxtCredentials) int {
	formData := JwxtInstance.GetLoginFormData(config.User, config.Password)
	LoginResponse := JwxtInstance.GetToken(formData)
	return LoginResponse
}

func (a *App) PayLogin(config models.PayCredentials) int {
	formData := PayInstance.GetLoginFormData(config.User, config.Password)
	loginResponse := PayInstance.GetToken(formData)
	return loginResponse
}

func (a *App) SettingJwxtLogin(user, password string) {
	formData := JwxtInstance.GetLoginFormData(user, password)
	LoginResponse := JwxtInstance.GetToken(formData)
	if LoginResponse != 302 {
		return
	}
	userInfo := models.JwxtCredentials{
		User:     user,
		Password: password,
	}
	JwxtInstance.SaveJwxtLoginConfig(userInfo)
}
func (a *App) LoadJwxtLogin() models.JwxtCredentials {
	return JwxtInstance.LoadJwxtLoginConfig()
}

func (a *App) LoadPayLogin() models.PayCredentials {
	return PayInstance.LoadPayLoginConfig()
}

func (a *App) SettingPayLogin(user, password string) {
	formData := PayInstance.GetLoginFormData(user, password)
	LoginResponse := PayInstance.GetToken(formData)
	if LoginResponse != 200 {
		return
	}
	userInfo := models.PayCredentials{
		User:     user,
		Password: password,
	}
	PayInstance.SavePayLoginConfig(userInfo)
}
