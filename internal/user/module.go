package user

import (
	"example.com/contact/internal/contact"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterModule(db *sqlx.DB) *chi.Mux {

	handler := NewUserHandler(
		NewUserRepository(db),
		contact.NewContactRepo(db),
	)

	router := NewUserRoute(handler)

	return router.GetRoutes()
}
