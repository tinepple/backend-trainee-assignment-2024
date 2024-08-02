package storage

import (
	"time"

	"github.com/lib/pq"
)

type Banner struct {
	Id        int           `db:"id"`
	TagIds    pq.Int64Array `db:"tag_ids"`
	FeatureID int           `db:"feature_id"`
	Content   []byte        `db:"content"`
	IsActive  bool          `db:"is_active"`
	CreatedAt time.Time     `db:"created_at"`
	UpdatedAt time.Time     `db:"updated_at"`
}
