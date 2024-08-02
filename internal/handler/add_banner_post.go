package handler

import (
	"backend-trainee-assignment-2024/internal/storage"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddBanner(c *gin.Context) {
	var req AddBannerRequest

	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	err := h.iStorage.CreateBanner(storage.Banner{
		Content:   req.Content,
		FeatureID: req.FeatureID,
		TagIds:    req.TagIds,
		IsActive:  req.IsActive,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println(req.FeatureID)
}
