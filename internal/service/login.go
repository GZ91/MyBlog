package service

func (s *Service) Authorized(userID string) (bool, error) {
	auth, err := s.NodeStorage.Authorized(userID)
	if err != nil {
		return false, err
	}
	if !auth {
		return false, nil
	}
	err = s.NodeStorage.InputFixation(userID)
	if err != nil {
		return true, err
	}
	return true, nil
}

func (s *Service) Login(login, password, userID string) (bool, error) {
	return s.NodeStorage.Login(login, password, userID)
}
