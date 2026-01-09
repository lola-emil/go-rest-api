package user

import (
	"github.com/go-chi/chi/v5"
)

type UserRoute struct {
	handler *UserHandler
}

func NewUserRoute(handler *UserHandler) *UserRoute {
	return &UserRoute{
		handler: handler,
	}
}

func (r *UserRoute) GetRoutes() *chi.Mux {
	mux := chi.NewRouter()

	mux.Get("/", r.handler.GetUsers)
	mux.Post("/", r.handler.PostUser)

	mux.Route("/{userId}", func(subRoute chi.Router) {
		subRoute.Get("/", r.handler.GetUser)
		subRoute.Delete("/", r.handler.DeleteUser)
		subRoute.Get("/contacts", r.handler.GetUserWithContacts)
	})

	return mux
}
