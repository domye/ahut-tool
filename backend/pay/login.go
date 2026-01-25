package pay

import (
	"ahut-tool/backend/models"
	"ahut-tool/backend/utils"
)

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

// SavePayLoginConfig 保存登录数据
func (s *Service) SavePayLoginConfig(data models.PayCredentials) {
	err := utils.SaveJSON(data, "PayCredentials.json")
	if err != nil {
		return
	}
}

// LoadPayLoginConfig 加载登陆数据
func (s *Service) LoadPayLoginConfig() models.PayCredentials {
	var config models.PayCredentials
	err := utils.LoadJSON("PayCredentials.json", &config)
	if err != nil {
		return models.PayCredentials{}
	}
	return config
}
