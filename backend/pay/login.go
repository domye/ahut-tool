package pay

func (s *Service) GetToken(username string, password string) int {
	status, _ := s.sendLoginRequest(username, password)
	return status
}
