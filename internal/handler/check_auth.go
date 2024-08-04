package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CheckAuth(c *gin.Context) {
	token := c.Request.Header.Get("Token")
	if token == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	role, err := h.iStorage.GetRole(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	if role != Admin {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
}

func (h *Handler) IsUserAdmin(c *gin.Context) bool {
	token := c.Request.Header.Get("Token")
	if token == "" {
		return false
	}

	role, err := h.iStorage.GetRole(token)
	if err != nil {
		return false
	}

	return role == Admin
}
