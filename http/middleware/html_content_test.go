package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bylexus/go-stdlib/http/middleware"
)

func TestHtmlContentMiddleware(t *testing.T) {
	// setup: create a test handler
	testH := middleware.HtmlContent(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	// serve the handler by using the http test response recorder:
	rr := httptest.NewRecorder()
	testH.ServeHTTP(rr, &http.Request{})

	// check if the Content-Type header is set to text/html
	if rr.Result().Header["Content-Type"][0] != "text/html" {
		t.Fatalf("HtmlContent middleware should add Content-Type header.")
	}
}
