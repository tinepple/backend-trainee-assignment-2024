package storage

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

func (s Storage) CreateBannerWithTags(banner Banner) (int, error) {

	tx, err := s.db.Beginx()
	if err != nil {
		return 0, err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()
	bannerId, err := s.createBanner(tx, banner)
	if err != nil {
		return 0, err
	}

	err = s.createBannerTags(tx, bannerId, banner.TagIds)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return bannerId, nil
}

func (s Storage) createBanner(tx *sqlx.Tx, banner Banner) (int, error) {

	query, params, err := sq.Insert(bannersTableName).
		Columns(
			"feature_id",
			"is_active",
			"content",
		).
		Values(
			banner.FeatureID,
			banner.IsActive,
			banner.Content,
		).
		Suffix("returning id").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	bannerId := 0

	err = tx.QueryRow(query, params...).Scan(&bannerId)
	if err != nil {
		return 0, err
	}

	return bannerId, nil
}

func (s Storage) createBannerTags(tx *sqlx.Tx, bannerID int, tagIDs []int64) error {
	for _, tagId := range tagIDs {
		query, params, err := sq.Insert(tagsTableName).
			Columns(
				"id",
			).
			Values(
				tagId,
			).
			PlaceholderFormat(sq.Dollar).
			ToSql()

		_, err = tx.Exec(query, params...)
		if err != nil {
			return err
		}

		query, params, err = sq.Insert(bannerTagsTableName).
			Columns(
				"banner_id",
				"tag_id",
			).
			Values(
				bannerID,
				tagId,
			).
			PlaceholderFormat(sq.Dollar).
			ToSql()

		_, err = tx.Exec(query, params...)
		if err != nil {
			return err
		}
	}

	return nil
}
