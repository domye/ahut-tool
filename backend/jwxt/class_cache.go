package jwxt

import (
	"ahut-tool/backend/models"
	"ahut-tool/backend/utils"
	"fmt"
)

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
