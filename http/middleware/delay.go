package middleware

import (
	"net/http"
	"time"
)

/**
connection limit middleware fn:
*/

func DelayRequest(duration time.Duration, f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)
		f.ServeHTTP(w, r)
	})
}
