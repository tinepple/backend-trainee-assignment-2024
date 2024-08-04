package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	router   *gin.Engine
	iStorage iStorage
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func (h *Handler) initRoutes() {
	adminGroup := h.router.Group("/", h.CheckAuth)
	adminGroup.GET("/user_banner", h.GetUserBanner)
	adminGroup.GET("/banner", h.GetBanner)
	adminGroup.POST("/banner", h.AddBanner)
	adminGroup.PATCH("/banner/:id", h.PatchBanner)
	adminGroup.DELETE("/banner/:id", h.DeleteBanner)
}

func New(iStorage iStorage) *Handler {
	h := &Handler{
		router:   gin.New(),
		iStorage: iStorage,
	}

	h.initRoutes()

	return h
}
