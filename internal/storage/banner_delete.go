package storage

import sq "github.com/Masterminds/squirrel"

func (s Storage) DeleteBanner(id int) error {
	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	query, params, err := sq.Delete(bannerTagsTableName).
		Where(sq.Eq{"banner_id": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return handleError(err)
	}

	_, err = tx.Exec(query, params...)
	if err != nil {
		return err
	}

	query, params, err = sq.Delete(bannersTableName).
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, params...)
	if err != nil {
		return err
	}

	return tx.Commit()
}
