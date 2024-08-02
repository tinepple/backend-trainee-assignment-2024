package handler

import storageDto "backend-trainee-assignment-2024/internal/storage"

type iStorage interface {
	CreateBanner(banner storageDto.Banner) error
	PatchBanner(banner storageDto.Banner) error
	DeleteBanner(id int) error
	GetBanners(tagId int, featureId, limit, offset int) ([]storageDto.Banner, error)
}
