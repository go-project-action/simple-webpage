package middlewares

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-project-action/simple-webpage/examples/max/pkg/ip"
)

// Logging middleware logs messages.
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now().UTC()
		path := r.URL.Path
		method := r.Method
		ip := ip.FromRequest(r)

		defer func() {
			// Skip static requests
			if strings.HasPrefix(path, "/static") {
				return
			}
			// Calculates the latency.
			end := time.Now().UTC()
			latency := end.Sub(start)
			log.Printf("%-13s | %-12s | %s %s", latency, ip, method, path)
		}()

		next.ServeHTTP(w, r)
	})
}
