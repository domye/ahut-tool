package pay

import (
	"ahut-tool/backend/models"
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
