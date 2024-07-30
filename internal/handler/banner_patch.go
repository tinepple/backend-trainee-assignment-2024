package handler

import (
	"backend-trainee-assignment-2024/internal/storage"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) PatchBanner(c *gin.Context) {
	var req PatchBannerRequest

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid id")
		return
	}

	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	err = h.iStorage.PatchBanner(storage.Banner{
		TagIds:    req.TagIds,
		FeatureID: req.FeatureID,
		Content: storage.Content{
			Title: req.Content.Title,
			Text:  req.Content.Text,
			URL:   req.Content.URL,
		},
		IsActive: req.IsActive,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Print(id)

	c.Status(http.StatusOK)
}
