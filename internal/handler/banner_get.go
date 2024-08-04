package handler

import (
	"backend-trainee-assignment-2024/internal/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetBanner(c *gin.Context) {
	token := c.Request.Header.Get("Token")

	role, err := h.iStorage.GetRole(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}

	if role != Admin {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	featureID, err := getQueryInt(c, "feature_id")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid feature_id")
		return
	}

	tagID, err := getQueryInt(c, "tag_id")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid tag_id")
		return
	}

	limit, err := getQueryInt(c, "limit")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid limit")
		return
	}

	offset, err := getQueryInt(c, "offset")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid offset")
		return
	}

	banners, err := h.iStorage.GetBannersAdmin(tagID, featureID, limit, offset)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map2Response(banners))
}

func getQueryInt(c *gin.Context, key string) (int, error) {
	if !c.Request.URL.Query().Has(key) {
		return 0, nil
	}
	result, err := strconv.Atoi(c.Request.URL.Query().Get(key))
	if err != nil {
		return 0, nil
	}
	return result, nil
}
func map2Response(banners []storage.Banner) []GetBannerResponse {
	resp := make([]GetBannerResponse, 0, len(banners))

	for _, banner := range banners {
		resp = append(resp, GetBannerResponse{
			BannerID:  banner.Id,
			TagIds:    banner.TagIds,
			FeatureID: banner.FeatureID,
			Content:   banner.Content,
			IsActive:  banner.IsActive,
			CreatedAt: banner.CreatedAt,
			UpdatedAt: banner.UpdatedAt,
		})
	}
	return resp
}
