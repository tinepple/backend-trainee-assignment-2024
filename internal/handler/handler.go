package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	router  *gin.Engine
	storage storage
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func (h *Handler) initRoutes() {
	h.router.GET("/user_baner", h.GetUserBanner)
}

func New(storage storage) *Handler {
	h := &Handler{
		router:  gin.New(),
		storage: storage,
	}

	h.initRoutes()

	return h
}
