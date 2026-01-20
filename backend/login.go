package backend

func (a *App) JwxtLogin(userId string, password string) int {
	formdate := JwxtInstance.GetLoginFormData(userId, password)
	LoginResponse := JwxtInstance.GetToken(formdate)
	return LoginResponse
}

func (a *App) PayLogin(userId string, password string) int {
	formData := PayInstance.GetLoginFormData(userId, password)
	loginResponse := PayInstance.GetToken(formData)
	return loginResponse
}
