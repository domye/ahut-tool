package pay

import (
	"ahut-tool/backend/models"
	"ahut-tool/backend/utils"
	"encoding/json"
)

func (s *Service) GetIMSFormData(xiaoqu string, ld_Name string, ld_Id string, Room_No string, etype string) map[string]string {
	formData := map[string]string{
		"xiaoqu":  xiaoqu,
		"ld_Name": ld_Name,
		"ld_Id":   ld_Id,
		"Room_No": Room_No,
		"etype":   etype,
	}
	return formData
}

func (s *Service) GetIMS(formData map[string]string) (*models.IMSResponse, error) {
	body, err := s.sendGetIMSRequest(formData)
	if err != nil {
		return nil, err
	}

	var response models.IMSResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	print(response.Data.UsedAmp)
	return &response, nil
}

// SaveDormConfig 保存宿舍配置数据
func (s *Service) SaveDormConfig(data models.DormConfig) {
	err := utils.SaveJSON(data, "DormConfig.json")
	if err != nil {
		return
	}
}

// LoadDormConfig 加载宿舍配置数据
func (s *Service) LoadDormConfig() models.DormConfig {
	var dorm models.DormConfig
	err := utils.LoadJSON("DormConfig.json", &dorm)
	if err != nil {
		return models.DormConfig{}
	}
	return dorm
}
