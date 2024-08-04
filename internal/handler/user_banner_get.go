package handler

import (
	"backend-trainee-assignment-2024/internal/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUserBanner(c *gin.Context) {
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

	var banners []storage.Banner

	if h.IsUserAdmin(c) {

		banners, err = h.iStorage.GetBannersAdmin(tagID, featureID, limit, offset)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			return
		}
	} else {
		banners, err = h.iStorage.GetBannersActive(tagID, featureID, limit, offset)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, map2Response(banners))
}
