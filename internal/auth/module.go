package auth

import (
	"example.com/contact/internal/user"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterModule(db *sqlx.DB) *chi.Mux {

	userRepo := user.NewUserRepository(db)
	authHandler := NewAuthHandler(userRepo)
	route := NewAuthRoute(authHandler)

	return route.GetRoutes()
}
