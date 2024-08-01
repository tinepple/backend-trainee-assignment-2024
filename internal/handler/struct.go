package handler

import (
	"encoding/json"
	"time"
)

type GetBannerResponse struct {
	BannerID  int
	TagIds    []int
	FeatureID int
	Content   json.RawMessage
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AddBannerRequest struct {
	TagIds    []int       `json:"tag_ids"`
	FeatureID int         `json:"feature_id"`
	Content   interface{} `json:"content"`
	IsActive  bool        `json:"is_active"`
}

type AddBannerResponse struct {
	BannerID int
}

type PatchBannerRequest struct {
	TagIds    []int           `json:"tag_ids"`
	FeatureID int             `json:"feature_id"`
	Content   json.RawMessage `json:"content"`
	IsActive  bool            `json:"is_active"`
}
