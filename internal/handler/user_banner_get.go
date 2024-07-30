package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUserBanner(c *gin.Context) {
	tagID, err := strconv.Atoi(c.Param("tag_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid id")
		return
	}

	featureID, err := strconv.Atoi(c.Param("feature_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid id")
		return
	}

	useLastVersion, err := strconv.ParseBool(c.Param("use_last_version"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid id")
		return
	}

	fmt.Println(tagID, featureID, useLastVersion)

	c.JSON(http.StatusOK, GetUserBannerResponse{})
}
