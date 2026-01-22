package news

// sendGetAcademicRequest 获取学业通知
func (s *Service) sendGetAcademicRequest() (string, error) {
	resp, err := s.client.R().
		Post(AcademicURL)

	if err != nil {
		return "", err
	}

	return resp.String(), nil
}

// sendGetAnnouncementRequest 获取公告通知
func (s *Service) sendGetAnnouncementRequest() (string, error) {
	resp, err := s.client.R().
		Post(AnnouncementURL)

	if err != nil {
		return "", err
	}

	return resp.String(), nil
}

// sendGetNewsRequest 获取学校要闻
func (s *Service) sendGetNewsRequest() (string, error) {
	resp, err := s.client.R().
		Post(SchoolNewsURL)

	if err != nil {
		return "", err
	}

	return resp.String(), nil
}
