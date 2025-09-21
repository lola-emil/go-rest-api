package auth

import "github.com/gin-gonic/gin"

type Handler struct {
	store *Store
}

func NewHandler(store *Store) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/register", h.register)
	router.POST("/login", h.login)
}

func (h *Handler) register(ctx *gin.Context) {

}

func (h *Handler) login(ctx *gin.Context) {

}
