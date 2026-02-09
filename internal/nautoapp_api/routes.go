package nautoapp_api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Application) MountRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(app.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/api/v1/health", app.healthCheckHandler)
	r.Get("/api/v1/inventory", app.getInventory)
	r.Post("/api/v1/inventory", app.queryInventory)

	return r
}
