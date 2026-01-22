package news

import (
	"ahut-tool/backend/models"
	"ahut-tool/backend/utils"
)

// GetAcademicNotifications 获取学业通知
func (s *Service) GetAcademicNotifications() ([]models.News, error) {
	resp, err := s.sendGetAcademicRequest()
	if err != nil {
		return nil, err
	}

	news, err := utils.ParseNotifications(resp)
	if err != nil {
		return nil, err
	}

	return news, nil
}

// GetAnnouncementNotifications 获取公告通知
func (s *Service) GetAnnouncementNotifications() ([]models.News, error) {
	resp, err := s.sendGetAnnouncementRequest()
	if err != nil {
		return nil, err
	}

	news, err := utils.ParseNotifications(resp)
	if err != nil {
		return nil, err
	}

	return news, nil
}

// GetSchoolNews 获取学校要闻
func (s *Service) GetSchoolNews() ([]models.News, error) {
	resp, err := s.sendGetNewsRequest()
	if err != nil {
		return nil, err
	}

	news, err := utils.ParseNews(resp)
	if err != nil {
		return nil, err
	}

	return news, nil
}
