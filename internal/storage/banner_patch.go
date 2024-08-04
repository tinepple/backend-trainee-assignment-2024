package storage

import (
	sq "github.com/Masterminds/squirrel"
)

func (s Storage) PatchBanner(banner BannerPatch) error {

	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	builder := sq.StatementBuilder.Update(bannersTableName).
		Where(sq.Eq{"id": banner.Id}).
		Set("updated_at", "now()").
		PlaceholderFormat(sq.Dollar)

	if banner.FeatureID.Valid {
		builder = builder.Set("feature_id", sq.Eq{"feature_id": banner.FeatureID})
	}

	if len(banner.Content) != 0 {
		builder = builder.Set("content", banner.Content)
	}

	if banner.IsActive.Valid {
		builder = builder.Set("is_active", banner.IsActive)
	}

	query, params, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, params...)
	if err != nil {
		return err
	}

	query, params, err = sq.Delete(bannerTagsTableName).
		Where(sq.Eq{"banner_id": banner.Id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, params...)
	if err != nil {
		return err
	}

	for _, tagId := range banner.TagIds {
		_, err = tx.Exec("insert into tags values ($1) on conflict do nothing;", tagId)
		if err != nil {
			return err
		}
		query, params, err = sq.Insert("banner_tags").
			Columns(
				"banner_id",
				"tag_id",
			).
			Values(
				banner.Id,
				tagId,
			).
			PlaceholderFormat(sq.Dollar).
			ToSql()

		_, err = tx.Exec(query, params...)
		if err != nil {
			return err
		}

	}

	return tx.Commit()
}
