package contact

import (
	"github.com/go-chi/chi/v5"
)

type ContactRoutes struct {
	handler *ContactHandler
}

func NewContactRoute(handler *ContactHandler) *ContactRoutes {
	return &ContactRoutes{
		handler: handler,
	}
}

func (r *ContactRoutes) GetRoutes() *chi.Mux {
	mux := chi.NewRouter()

	mux.Get("/", r.handler.GetContacts)
	mux.Post("/", r.handler.PostContact)

	mux.Route("/{contactId}", func(subMux chi.Router) {
		subMux.Get("/", r.handler.GetContact)
		subMux.Delete("/", r.handler.DeleteContact)
	})

	return mux
}
