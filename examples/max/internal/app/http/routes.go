package http

import (
	"net/http"

	"github.com/go-project-action/simple-webpage/examples/max/internal/app/http/middlewares"
	"github.com/gorilla/mux"
)

func (a *AppServer) publicRoutes(r *mux.Router) {
	r.Use(middlewares.Logging)
	// Serving static files.
	fs := http.FileServer(http.Dir("web/static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	r.HandleFunc("/", a.Home())
}
