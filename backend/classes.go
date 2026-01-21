package backend

import (
	"ahut-tool/backend/models"
)

// GetClassSchedule 获取课程表信息
func (a *App) GetClassSchedule(xnxq string) (*models.ClassScheduleResponse, error) {
	formData := JwxtInstance.GetClassesFormData(xnxq)
	classes, err := JwxtInstance.GetClasses(formData)
	if err != nil {
		return nil, err
	}

	return &models.ClassScheduleResponse{
		Classes: classes,
	}, nil
}
