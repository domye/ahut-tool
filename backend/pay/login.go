package pay

import "ahut-tool/backend/utils"

func (s *Service) GetLoginFormData(userId string, password string) map[string]string {

	encryptedPwd, _ := utils.EncryptPasswordWithRSA(password)
	formData := map[string]string{
		"username": userId,
		"pwd":      encryptedPwd,
	}

	return formData
}

func (s *Service) GetToken(formData map[string]string) int {
	status, _ := s.sendLoginRequest(formData)
	return status
}
