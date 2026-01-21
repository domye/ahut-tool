package backend

func (a *App) JwxtLogin() int {
	userInfo, _ := ConfigInstance.GetJwxtCredentials()
	formdate := JwxtInstance.GetLoginFormData(userInfo.User, userInfo.Password)
	LoginResponse := JwxtInstance.GetToken(formdate)
	return LoginResponse
}

func (a *App) PayLogin() int {
	userInfo, _ := ConfigInstance.GetPayCredentials()
	formData := PayInstance.GetLoginFormData(userInfo.User, userInfo.Password)
	loginResponse := PayInstance.GetToken(formData)
	return loginResponse
}
