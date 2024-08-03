package storage

func (s Storage) GetRole(token string) (string, error) {
	var role string
	err := s.db.QueryRow("select role from person where token = $1", token).Scan(
		&role,
	)
	if err != nil {
		return role, err
	}

	return role, nil
}
