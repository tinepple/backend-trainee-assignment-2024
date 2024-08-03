package storage

import (
	"database/sql"
	"errors"
)

var ErrNotFound = errors.New("not found")

func handleError(err error) error {
	if err == sql.ErrNoRows {
		return ErrNotFound
	}
	return err
}
