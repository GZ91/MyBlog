package storage

func (s *Storage) Login(login, password, userID string) (bool, error) {
	row := s.db.QueryRow("SELECT COUNT(*) FROM users WHERE login = $1 AND password = $2", login, password)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	_, err = s.db.Exec("UPDATE users SET userID = $1 WHERE login = $2 AND password = $3", userID, login, password)
	if err != nil {
		return true, err
	}
	return true, nil
}

func (s *Storage) Authorized(userID string) (bool, error) {
	row := s.db.QueryRow("SELECT COUNT(*) FROM users WHERE userID=$1", userID)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

func (s *Storage) InputFixation(userID string) error {
	row := s.db.QueryRow("SELECT login FROM users WHERE userID = $1", userID)
	var login string
	err := row.Scan(&login)
	if err != nil {
		return err
	}
	_, err = s.db.Exec("INSERT INTO input_fixation (login) VALUES ($1)", login)
	if err != nil {
		return err
	}
	return nil
}
