package storage

import sq "github.com/Masterminds/squirrel"

func (s Storage) GetRole(token string) (string, error) {
	query, params, err := sq.Select(
		"role",
	).From(
		personTableName,
	).Where(
		sq.Eq{"token": token},
	).PlaceholderFormat(sq.Dollar).ToSql()

	var role string
	err = s.db.QueryRow(query, params...).Scan(
		&role,
	)
	if err != nil {
		return "", err
	}

	return role, nil
}
