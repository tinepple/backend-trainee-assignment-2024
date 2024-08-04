package handler

import (
	"backend-trainee-assignment-2024/internal/storage"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"go.octolab.org/pointer"

	"github.com/gin-gonic/gin"
)

func (h *Handler) PatchBanner(c *gin.Context) {
	var req PatchBannerRequest

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

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid id")
		return
	}

	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	featureID := int64(0)

	if req.FeatureID != nil {
		featureID = int64(*req.FeatureID)
	}

	err = h.iStorage.PatchBanner(storage.BannerPatch{
		Id: sql.NullInt64{
			Int64: int64(id),
			Valid: id != 0,
		},
		TagIds: req.TagIds,
		FeatureID: sql.NullInt64{
			Int64: featureID,
			Valid: req.FeatureID != nil,
		},
		Content: req.Content,
		IsActive: sql.NullBool{
			Bool:  pointer.ValueOfBool(req.IsActive),
			Valid: req.IsActive != nil,
		},
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Print(id)

	c.Status(http.StatusOK)
}
