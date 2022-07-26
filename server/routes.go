package server

import (
	"github.com/go-chi/chi"
)

func (srv *Server) InjectRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", srv.Health)
	r.Route("/api", func(api chi.Router) {
		api.Get("/rate", srv.GetRate)
	})
	return r
}
