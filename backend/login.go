package backend

func (a *App) JwxtLogin() int {
	userInfo, _ := ConfigInstance.GetJwxtCredentials()
	formdate := JwxtInstance.GetLoginFormData(userInfo.User, userInfo.Password)
	LoginResponse := JwxtInstance.GetToken(formdate)
	return LoginResponse
}

func (a *App) PayLogin(userId string, password string) int {
	formData := PayInstance.GetLoginFormData(userId, password)
	loginResponse := PayInstance.GetToken(formData)
	return loginResponse
}
