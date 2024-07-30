package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUserBanner(c *gin.Context) {
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

	useLastVersion, err := strconv.ParseBool(c.Request.URL.Query().Get("use_last_version"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid use_last_version")
		return
	}

	fmt.Println(tagID, featureID, useLastVersion)

	c.JSON(http.StatusOK, UserBanner{})
}
