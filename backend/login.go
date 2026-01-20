package backend

func (a *App) Login(userId string, password string) int {
	formdate := Instance.GetLoginFormData(userId, password)
	LoginResponse := Instance.GetToken(formdate)
	return LoginResponse
}
