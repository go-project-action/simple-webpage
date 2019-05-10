package http

import (
	"net/http"
)

func (a *AppServer) publicRoutes(r *http.ServeMux) {
	// Serving static files.
	fs := http.FileServer(http.Dir("web/static"))
	r.Handle("/static/", http.StripPrefix("/static/", fs))

	r.HandleFunc("/", a.Home())
}
