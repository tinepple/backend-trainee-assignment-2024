package storage

import (
	sq "github.com/Masterminds/squirrel"
)

func (s Storage) CreateBanner(banner Banner) error {

	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	query, params, err := sq.Insert("banners").
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

	err = s.db.QueryRow(query, params...).Scan(&bannerId)
	if err != nil {
		return err
	}

	for _, tagId := range banner.TagIds {
		_, err = s.db.Exec("insert into tags values ($1) on conflict do nothing;", tagId)
		if err != nil {
			return err
		}
		query, params, err = sq.Insert("banner_tags").
			Columns(
				"banner_id",
				"tag_id",
			).
			Values(
				bannerId,
				tagId,
			).
			PlaceholderFormat(sq.Dollar).
			ToSql()

		_, err = s.db.Exec(query, params...)
		if err != nil {
			return err
		}

	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
