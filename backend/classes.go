package backend

import (
	"ahut-tool/backend/models"
	"fmt"
)

// GetClassSchedule 获取课程表信息
func (a *App) GetClassSchedule(xnxq string) (*models.ClassScheduleResponse, error) {
	// 检查本地是否有缓存
	if ClassCacheServiceInstance.HasClassSchedule(xnxq) {
		// 从本地加载课程表
		classes, err := ClassCacheServiceInstance.LoadClassSchedule(xnxq)
		if err != nil {
			return nil, fmt.Errorf("failed to load class schedule from cache: %v", err)
		}
		return &models.ClassScheduleResponse{
			Classes: classes,
		}, nil
	}

	// 从网络获取课程表
	formData := JwxtInstance.GetClassesFormData(xnxq)
	classes, err := JwxtInstance.GetClasses(formData)
	if err != nil {
		return nil, err
	}

	// 保存课程表到本地缓存
	err = ClassCacheServiceInstance.SaveClassSchedule(xnxq, classes)
	if err != nil {
		// 保存失败不影响返回结果，可以记录日志
		fmt.Printf("Warning: failed to save class schedule to cache: %v\n", err)
	}

	return &models.ClassScheduleResponse{
		Classes: classes,
	}, nil
}
