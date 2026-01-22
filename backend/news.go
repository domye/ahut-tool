package backend

import (
	"ahut-tool/backend/models"
)

// GetAcademicNotifications 获取学业通知
func (a *App) GetAcademicNotifications() ([]models.News, error) {
	return NewsInstance.GetAcademicNotifications()
}

// GetAnnouncementNotifications 获取公告通知
func (a *App) GetAnnouncementNotifications() ([]models.News, error) {
	return NewsInstance.GetAnnouncementNotifications()
}

// GetSchoolNews 获取学校要闻
func (a *App) GetSchoolNews() ([]models.News, error) {
	return NewsInstance.GetSchoolNews()
}
