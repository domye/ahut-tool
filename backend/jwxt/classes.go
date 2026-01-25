package jwxt

import (
	"ahut-tool/backend/models"
	"ahut-tool/backend/utils"
	"fmt"
)

// GetClassesFormData 获取课程表
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

// SaveSchedulesConfig 保存课程表配置数据
func (s *Service) SaveSchedulesConfig(data models.SchedulesConfig) {
	err := utils.SaveJSON(data, "classesConfig.json")
	if err != nil {
		return
	}
}

// LoadSchedulesConfig 加载课程表配置数据
func (s *Service) LoadSchedulesConfig() models.SchedulesConfig {
	var classes models.SchedulesConfig
	err := utils.LoadJSON("classesConfig.json", &classes)
	if err != nil {
		return models.SchedulesConfig{}
	}
	return classes
}

// SaveClassSchedule 保存课程表数据
func (s *Service) SaveClassSchedule(xnxq string, classes []models.Class) error {
	// 构造文件名
	filename := fmt.Sprintf("class_schedule_%s.json", xnxq)

	// 保存课程表数据
	err := utils.SaveJSON(classes, filename)
	if err != nil {
		return fmt.Errorf("failed to save class schedule: %v", err)
	}

	return nil
}

// LoadClassSchedule 加载课程表数据
func (s *Service) LoadClassSchedule(xnxq string) ([]models.Class, error) {
	// 构造文件名
	filename := fmt.Sprintf("class_schedule_%s.json", xnxq)

	// 检查文件是否存在
	if !utils.FileExists(filename) {
		return nil, fmt.Errorf("class schedule file does not exist")
	}

	// 加载课程表数据
	var classes []models.Class
	err := utils.LoadJSON(filename, &classes)
	if err != nil {
		return nil, fmt.Errorf("failed to load class schedule: %v", err)
	}

	return classes, nil
}

// HasClassSchedule 检查是否存在指定学期的课程表缓存
func (s *Service) HasClassSchedule(xnxq string) bool {
	// 构造文件名
	filename := fmt.Sprintf("class_schedule_%s.json", xnxq)

	// 检查文件是否存在
	return utils.FileExists(filename)
}
