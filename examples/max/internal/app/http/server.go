package http

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// AppServer contains the information to run a server.
type AppServer struct{}

// Run will start the http server.
func (a *AppServer) Run(port string) {
	r := http.NewServeMux()
	a.publicRoutes(r)

	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%s", port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	log.Printf("Server started at port %v", port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
