package jwxt

import (
	"ahut-tool/backend/models"
	"ahut-tool/backend/utils"
)

func (s *Service) GetClassesFormData(kcxq string) map[string]string {
	formData := map[string]string{
		"xnxq01id": kcxq,
	}
	return formData
}

func (s *Service) GetClasses(formData map[string]string) ([]models.Class, error) {
	resp, err := s.sendGetClassesRequest(formData)
	if err != nil {
		return nil, err
	}

	classes, err := utils.ParseClassSchedule(resp)
	if err != nil {
		return nil, err
	}

	return classes, nil
}
