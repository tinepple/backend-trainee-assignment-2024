package handler

import (
	storageDto "backend-trainee-assignment-2024/internal/storage"
)

type iStorage interface {
	CreateBanner(banner storageDto.Banner) error
	PatchBanner(banner storageDto.Banner) error
	DeleteBanner(id int) error
	GetBannersAdmin(tagId int, featureId, limit, offset int) ([]storageDto.Banner, error)
	GetBannersActive(tagId int, featureId, limit, offset int) ([]storageDto.Banner, error)
	GetRole(token string) (string, error)
}
