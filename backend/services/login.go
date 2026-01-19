package services

import "ahut-tool/backend/utils"

type LoginResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Name   string `json:"name"`
		Cookie string `json:"cookie"`
	}
}

func getLoginFormData(userId string, password string) map[string]string {
	encoded := utils.Base64Encode(userId) + "%%%" + utils.Base64Encode(password) + "="
	FormData := map[string]string{
		"userAccount":  userId,
		"userPassword": "",
		"encoded":      encoded,
		"pwdstr1":      "",
		"pwdstr2":      "",
	}
	return FormData
}

func getToken(formData map[string]string) string {
	request.getCookie(formData)

}
