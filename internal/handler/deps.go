package handler

import (
	storageDto "backend-trainee-assignment-2024/internal/storage"
)

type iStorage interface {
	CreateBannerWithTags(banner storageDto.Banner) (int, error)
	PatchBanner(banner storageDto.BannerPatch) error
	DeleteBannerWithTags(id int) error
	GetBannersAdmin(tagId int, featureId, limit, offset int) ([]storageDto.Banner, error)
	GetBannersActive(tagId int, featureId, limit, offset int) ([]storageDto.Banner, error)
	GetRole(token string) (string, error)
}
