package handler

import (
	"encoding/json"
	"time"
)

type GetBannerResponse struct {
	BannerID  int
	TagIds    []int64
	FeatureID int
	Content   json.RawMessage
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AddBannerRequest struct {
	TagIds    []int64         `json:"tag_ids"`
	FeatureID int             `json:"feature_id"`
	Content   json.RawMessage `json:"content"`
	IsActive  bool            `json:"is_active"`
}

type AddBannerResponse struct {
	BannerID int
}

type PatchBannerRequest struct {
	TagIds    []int64         `json:"tag_ids"`
	FeatureID int             `json:"feature_id"`
	Content   json.RawMessage `json:"content"`
	IsActive  bool            `json:"is_active"`
}
