package storage

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

func (s Storage) GetBannersAdmin(tagId int, featureId, limit, offset int) ([]Banner, error) {
	builder := sq.StatementBuilder.Select(
		"b.id",
		"b.feature_id",
		"b.is_active",
		"b.content",
		"array_agg(bt.tag_id) as tag_ids",
		"b.created_at",
		"b.updated_at",
	).From(fmt.Sprintf("%s b", bannersTableName)).
		InnerJoin(fmt.Sprintf("%s bt on bt.banner_id = b.id", bannerTagsTableName)).
		GroupBy(
			"b.id",
			"b.feature_id",
			"b.is_active",
			"b.content",
			"b.created_at",
			"b.updated_at",
		).PlaceholderFormat(sq.Dollar)

	if tagId != 0 {
		builder = builder.Where(sq.Eq{"bt.tag_id": tagId})
	}

	if featureId != 0 {
		builder = builder.Where(sq.Eq{"b.feature_id": featureId})
	}

	if limit != 0 {
		builder = builder.Limit(uint64(limit))
	}

	if offset != 0 {
		builder = builder.Offset(uint64(offset))
	}

	query, params, err := builder.ToSql()

	if err != nil {
		return nil, err
	}

	dest := make([]Banner, 0, 10)
	err = s.db.Select(&dest, s.db.Rebind(query), params...)
	if err != nil {
		return nil, err
	}

	return dest, nil
}

func (s Storage) GetBannersActive(tagId int, featureId, limit, offset int) ([]Banner, error) {
	builder := sq.StatementBuilder.Select(
		"b.id",
		"b.feature_id",
		"b.is_active",
		"b.content",
		"array_agg(bt.tag_id) as tag_ids",
		"b.created_at",
		"b.updated_at",
	).From(fmt.Sprintf("%s b", bannersTableName)).
		InnerJoin(fmt.Sprintf("%s bt on bt.banner_id = b.id", bannerTagsTableName)).
		GroupBy(
			"b.id",
			"b.feature_id",
			"b.is_active",
			"b.content",
			"b.created_at",
			"b.updated_at",
		).
		Where("b.is_active = true").
		PlaceholderFormat(sq.Dollar)

	if tagId != 0 {
		builder = builder.Where(sq.Eq{"bt.tag_id": tagId})
	}

	if featureId != 0 {
		builder = builder.Where(sq.Eq{"b.feature_id": featureId})
	}

	if limit != 0 {
		builder = builder.Limit(uint64(limit))
	}

	if offset != 0 {
		builder = builder.Offset(uint64(offset))
	}

	query, params, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	dest := make([]Banner, 0, 10)
	err = s.db.Select(&dest, s.db.Rebind(query), params...)
	if err != nil {
		return nil, err
	}

	return dest, nil
}
