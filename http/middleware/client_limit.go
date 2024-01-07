package middleware

import "net/http"

// ClientLimit defines an HTTP middleware that ensures no more than
// maxClients requests are processed concurrently to the given handler f.
func ClientLimit(h http.Handler, maxClients int) http.HandlerFunc {
	// Counting semaphore using a buffered channel
	sema := make(chan struct{}, maxClients)

	return func(w http.ResponseWriter, req *http.Request) {
		// fill the buffer with an entry:
		sema <- struct{}{}
		// read from the buffer to release the entry, after the request is processed:
		defer func() { <-sema }()
		h.ServeHTTP(w, req)
	}
}
