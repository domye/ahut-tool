package jwxt

import (
	"ahut-tool/backend/models"
	"ahut-tool/backend/utils"
)

func (s *Service) GetGradesFormData(kksj, kcxz, kcmc, xsfs string) map[string]string {
	formData := map[string]string{
		"kksj": kksj,
		"kcxz": kcxz,
		"kcmc": kcmc,
		"xsfs": xsfs,
	}
	return formData
}

func (s *Service) GetGrades(formData map[string]string) ([]models.Grade, *models.GradeSummary, error) {

	resp, err := s.sendGetGradesRequest(formData)

	grades, summary, err := utils.ParseGrades(resp)
	if err != nil {
		return nil, nil, err
	}

	return grades, summary, nil
}
