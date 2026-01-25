package jwxt

import (
	"ahut-tool/backend/models"
	"ahut-tool/backend/utils"
)

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

// SaveJwxtLoginConfig 保存登录数据
func (s *Service) SaveJwxtLoginConfig(data models.JwxtCredentials) {
	err := utils.SaveJSON(data, "JwxtCredentials.json")
	if err != nil {
		return
	}
}

// LoadJwxtLoginConfig 加载课程表配置数据
func (s *Service) LoadJwxtLoginConfig() models.JwxtCredentials {
	var config models.JwxtCredentials
	err := utils.LoadJSON("JwxtCredentials.json", &config)
	if err != nil {
		return models.JwxtCredentials{}
	}
	return config
}

func (s *Service) ExistJwxtLoginConfig() bool {
	return utils.FileExists("JwxtCredentials.json")
}
