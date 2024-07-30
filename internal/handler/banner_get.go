package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetBanner(c *gin.Context) {
	featureID, err := strconv.Atoi(c.Request.URL.Query().Get("feature_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid feature_id")
		return
	}

	tagID, err := strconv.Atoi(c.Request.URL.Query().Get("tag_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid tag_id")
		return
	}

	limit, err := strconv.Atoi(c.Request.URL.Query().Get("limit"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid limit")
		return
	}

	offset, err := strconv.Atoi(c.Request.URL.Query().Get("offset"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid offset")
		return
	}

	fmt.Println(tagID, featureID, limit, offset)

	c.JSON(http.StatusOK, []GetBannerResponse{}) //вообще там массив таких структур!!!
}
