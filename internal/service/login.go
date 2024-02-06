package service

func (s *Service) Authorized(userID string) (bool, error) {
	return s.NodeStorage.Authorized(userID)
}

func (s *Service) Login(login, password, userID string) (bool, error) {
	return s.NodeStorage.Login(login, password, userID)
}
