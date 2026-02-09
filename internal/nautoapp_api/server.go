package nautoapp_api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *Application) StartServer(r *chi.Mux) error {
	srv := &http.Server{
		Addr:    app.Config.addr,
		Handler: r,
	}
	app.AppLogger.Info(fmt.Sprintf("Server started @ %s", app.Config.addr))
	return srv.ListenAndServe()
}
