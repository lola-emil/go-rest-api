package contact

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func RegisterModule(db *sqlx.DB) *chi.Mux {
	contactRepo := NewContactRepo(db)
	contactHandler := NewContactHandler(contactRepo)
	contactRouter := NewContactRoute(contactHandler)

	return contactRouter.GetRoutes()
}
