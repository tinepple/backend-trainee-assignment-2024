package storage

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

func (s Storage) DeleteBannerWithTags(bannerID int) error {
	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	err = s.deleteBannerTags(tx, bannerID)
	if err != nil {
		return err
	}

	err = s.deleteBanner(tx, bannerID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s Storage) deleteBannerTags(tx *sqlx.Tx, bannerID int) error {
	query, params, err := sq.Delete(bannerTagsTableName).
		Where(sq.Eq{"banner_id": bannerID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return handleError(err)
	}

	_, err = tx.Exec(query, params...)
	if err != nil {
		return err
	}

	return nil
}

func (s Storage) deleteBanner(tx *sqlx.Tx, bannerID int) error {
	query, params, err := sq.Delete(bannersTableName).
		Where(sq.Eq{"id": bannerID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, params...)
	if err != nil {
		return err
	}

	return nil
}
