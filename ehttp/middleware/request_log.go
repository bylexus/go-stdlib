package middleware

import (
	"log"
	"net/http"
)

func RequestLog(logger *log.Logger, f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		url := r.URL.String()
		logger.Printf("%s %s", ip, url)
		f.ServeHTTP(w, r)
	})
}
