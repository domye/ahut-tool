package backend

import (
	"ahut-tool/backend/models"
)

// GradesResponse 成绩响应结构体
type GradesResponse struct {
	Grades  []models.Grade       `json:"grades"`
	Summary *models.GradeSummary `json:"summary"`
}

// GetGrades 获取成绩信息
func (a *App) GetGrades(kksj, kcxz, kcmc, xsfs string) (*GradesResponse, error) {
	formData := Instance.GetGradesFormData(kksj, kcxz, kcmc, xsfs)
	grades, summary, err := Instance.GetGrades(formData)
	if err != nil {
		return nil, err
	}

	return &GradesResponse{
		Grades:  grades,
		Summary: summary,
	}, nil
}
