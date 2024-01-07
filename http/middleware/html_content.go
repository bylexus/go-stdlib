package middleware

import "net/http"

func HtmlContent(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		f.ServeHTTP(w, r)
	})
}
