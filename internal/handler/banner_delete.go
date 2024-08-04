package handler

import (
	"backend-trainee-assignment-2024/internal/storage"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) DeleteBanner(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid id")
		return
	}

	err = h.iStorage.DeleteBannerWithTags(id)
	if err != nil {
		handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func handleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, storage.ErrNotFound):
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
}
