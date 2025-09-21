package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	store ProductStore
}

func NewHandler(store ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/products", h.getProducts)
}

func (h *Handler) getProducts(ctx *gin.Context) {
	products, err := h.store.GetProducts()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, products)
}
