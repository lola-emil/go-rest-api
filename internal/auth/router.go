package auth

import "github.com/go-chi/chi/v5"

type AuthRoute interface {
	GetRoutes() *chi.Mux
}

type route struct {
	handler AuthHandler
}

func NewAuthRoute(handler AuthHandler) AuthRoute {
	return &route{
		handler: handler,
	}
}

func (r *route) GetRoutes() *chi.Mux {
	mux := chi.NewRouter()

	mux.Post("/", r.handler.Login)

	return mux
}
