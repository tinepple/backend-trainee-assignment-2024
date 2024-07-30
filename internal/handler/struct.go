package handler

import (
	"time"
)

type UserBanner struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	URL   string `json:"url"`
}

type GetBannerResponse struct {
	BannerID  int
	TagIds    []int
	FeatureID int
	Content   UserBanner
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AddBannerRequest struct {
	TagIds    []int      `json:"tag_ids"`
	FeatureID int        `json:"feature_id"`
	Content   UserBanner `json:"content"`
	IsActive  bool       `json:"is_active"`
}

type AddBannerResponse struct {
	BannerID int
}

type PatchBannerRequest struct {
	TagIds    []int      `json:"tag_ids"`
	FeatureID int        `json:"feature_id"`
	Content   UserBanner `json:"content"`
	IsActive  bool       `json:"is_active"`
}
