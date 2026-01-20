package jwxt

import "ahut-tool/backend/utils"

func (s *Service) GetLoginFormData(userId string, password string) map[string]string {
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

func (s *Service) GetToken(formData map[string]string) int {
	status, _ := s.sendLoginRequest(formData)
	return status
}
