package backend

import (
	"ahut-tool/backend/models"
)

// GetGrades 获取成绩信息
func (a *App) GetGrades(kksj, kcxz, kcmc, xsfs string) (*models.GradesResponse, error) {
	formData := JwxtInstance.GetGradesFormData(kksj, kcxz, kcmc, xsfs)
	grades, summary, err := JwxtInstance.GetGrades(formData)
	if err != nil {
		return nil, err
	}

	return &models.GradesResponse{
		Grades:  grades,
		Summary: summary,
	}, nil
}
