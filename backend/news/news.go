package news

import (
	"ahut-tool/backend/models"
	"ahut-tool/backend/utils"
)

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
