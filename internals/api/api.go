package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/lola-emil/go-rest-api/internals/modules/auth"
	"github.com/lola-emil/go-rest-api/internals/modules/product"
)

type APIServer struct {
	addr string
	db   *sqlx.DB
}

func NewApiServer(addr string, db *sqlx.DB) *APIServer {
	return &APIServer{addr: addr, db: db}
}

func (s *APIServer) Run() error {
	route := gin.Default()

	v1 := route.Group("v1")

	// Register auth route
	authStore := auth.NewStore(s.db)
	authHandler := auth.NewHandler(authStore)
	authHandler.RegisterRoutes(v1)

	// Register products route
	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(v1)

	return route.Run()
}
