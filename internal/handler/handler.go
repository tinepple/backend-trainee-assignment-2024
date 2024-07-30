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
	h.router.GET("/user_banner", h.GetUserBanner)
	h.router.GET("/banner", h.GetBanner)
	h.router.POST("/banner", h.AddBanner)
	h.router.PATCH("/banner/:id", h.PatchBanner)
	h.router.DELETE("/banner/:id", h.DeleteBanner)
}

func New(iStorage iStorage) *Handler {
	h := &Handler{
		router:   gin.New(),
		iStorage: iStorage,
	}

	h.initRoutes()

	return h
}
