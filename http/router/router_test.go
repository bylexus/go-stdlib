package router

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/bylexus/go-stdlib/log"
)

type findRouteTestData struct {
	method               string
	url                  string
	expectMatch          bool
	expectedMethod       string
	expectedParams       *RouteParams
	expectedRoutePattern string
}

func findRouteByMethodAndUrlTestDataProvider() []findRouteTestData {
	testdata := make([]findRouteTestData, 0)

	// set
	testdata = append(testdata, findRouteTestData{
		method:               METHOD_GET,
		expectedMethod:       METHOD_GET,
		url:                  "/",
		expectMatch:          true,
		expectedRoutePattern: "/",
	},
	)

	// set
	testdata = append(testdata, findRouteTestData{
		method:               METHOD_POST,
		expectedMethod:       METHOD_POST,
		url:                  "/",
		expectMatch:          true,
		expectedRoutePattern: "/",
	},
	)

	// set
	testdata = append(testdata, findRouteTestData{
		method:               METHOD_PUT,
		expectedMethod:       METHOD_PUT,
		url:                  "/",
		expectMatch:          true,
		expectedRoutePattern: "/",
	},
	)

	testdata = append(testdata, findRouteTestData{
		method:               METHOD_PATCH,
		expectedMethod:       METHOD_PATCH,
		url:                  "/",
		expectMatch:          true,
		expectedRoutePattern: "/",
	},
	)

	testdata = append(testdata, findRouteTestData{
		method:               METHOD_DELETE,
		expectedMethod:       METHOD_DELETE,
		url:                  "/",
		expectMatch:          true,
		expectedRoutePattern: "/",
	},
	)

	testdata = append(testdata, findRouteTestData{
		method:               METHOD_OPTIONS,
		expectedMethod:       METHOD_OPTIONS,
		url:                  "/",
		expectMatch:          true,
		expectedRoutePattern: "/",
	},
	)

	testdata = append(testdata, findRouteTestData{
		method:               METHOD_HEAD,
		expectedMethod:       METHOD_HEAD,
		url:                  "/",
		expectMatch:          true,
		expectedRoutePattern: "/",
	},
	)

	testdata = append(testdata, findRouteTestData{
		method:               METHOD_TRACE,
		expectedMethod:       METHOD_TRACE,
		url:                  "/",
		expectMatch:          true,
		expectedRoutePattern: "/",
	},
	)

	testdata = append(testdata, findRouteTestData{
		method:               METHOD_PUT,
		expectedMethod:       METHOD_PUT,
		url:                  "/anyRoute",
		expectMatch:          true,
		expectedRoutePattern: "/anyRoute",
	},
	)

	// set
	testdata = append(testdata, findRouteTestData{
		method:               METHOD_GET,
		url:                  "/foo/bar",
		expectMatch:          false,
		expectedRoutePattern: "",
	},
	)

	// -------------- param match test data -----------------
	// match ok:
	testdata = append(testdata, findRouteTestData{
		method:               METHOD_GET,
		expectedMethod:       METHOD_GET,
		url:                  "/api/test_entity_1/1234",
		expectMatch:          true,
		expectedRoutePattern: "/api/{:entity}/{:id|[0-9]+}",
		expectedParams:       &RouteParams{"entity": "test_entity_1", "id": "1234"},
	},
	)
	// match with url tail:
	testdata = append(testdata, findRouteTestData{
		method:               METHOD_GET,
		expectedMethod:       METHOD_GET,
		url:                  "/api/test_entity_1/1234",
		expectMatch:          true,
		expectedRoutePattern: "/api/{:entity}/{:id|[0-9]+}",
		expectedParams:       &RouteParams{"entity": "test_entity_1", "id": "1234"},
	},
	)
	// no match
	testdata = append(testdata, findRouteTestData{
		method:               METHOD_POST,
		expectedMethod:       METHOD_POST,
		url:                  "/api/miau/foo/bar/123/",
		expectMatch:          true,
		expectedRoutePattern: "/api/{:param1}/{:urlTail|.*}",
		expectedParams:       &RouteParams{"param1": "miau", "urlTail": "foo/bar/123/"},
	},
	)
	return testdata
}

func TestFindRouteByMethodAndURL(t *testing.T) {
	// setup:
	l := log.NewSeverityLogger(io.Discard)
	r := NewRouter(&l)
	dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// just a dummy handler - not used
	})

	// define routes
	// simple routes
	r.Get("/", dummyHandler)
	r.Post("/", dummyHandler)
	r.Put("/", dummyHandler)
	r.Delete("/", dummyHandler)
	r.Patch("/", dummyHandler)
	r.Options("/", dummyHandler)
	r.Head("/", dummyHandler)
	r.Trace("/", dummyHandler)

	// match any method
	r.Any("/anyRoute", dummyHandler)
	r.Post("/anyRoute", dummyHandler)

	// match with params
	r.Get("/api/{:entity}/{:id|[0-9]+}", dummyHandler)

	// match with params
	r.Post("/api/{:param1}/{:urlTail|.*}", dummyHandler)

	// loop throug the testdata set:
	testdata := findRouteByMethodAndUrlTestDataProvider()
	for i, data := range testdata {
		t.Logf("Test %d: %v", i, data)

		// define request:
		req := &http.Request{
			Method: data.method,
			URL:    &url.URL{Path: data.url},
		}
		// execute:
		matchedRoute := r.findRoute(req)

		// verify:

		// route should not match:
		if data.expectMatch == false && matchedRoute != nil {
			t.Fatalf("Route for URL %s should not match", data.url)
		}
		// route should match:
		if data.expectMatch == true {
			// should match, check the route
			if matchedRoute == nil {
				t.Fatalf("Route not found.")
			}
			// check method
			if matchedRoute.Method != data.expectedMethod {
				t.Errorf("Route method should be %s, but is %s", data.expectedMethod, matchedRoute.Method)
			}
			// check route pattern
			if matchedRoute.Route.Pattern != data.expectedRoutePattern {
				t.Errorf("Route pattern should be %s, but is %s", data.expectedRoutePattern, matchedRoute.Route.Pattern)
			}
			// check route params
			if data.expectedParams != nil && !reflect.DeepEqual(matchedRoute.Params, *data.expectedParams) {
				t.Errorf("Route params do not match: should be: %s,  but is %s", *data.expectedParams, matchedRoute.Params)
			}
		}
	}
}

func TestRouterInjectsMatchedRouteToContext(t *testing.T) {
	// setup:
	l := log.NewSeverityLogger(io.Discard)
	r := NewRouter(&l)
	var routeParams RouteParams

	dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		matchedRoute := GetMatchedRoute(r.Context())
		routeParams = matchedRoute.Params
	})

	// Define route with some parameters
	r.Get("/api/{:entity}/{:id|[0-9]+}/{:urlTail|.*}", dummyHandler)

	// define request:
	req := &http.Request{
		Method: METHOD_GET,
		URL:    &url.URL{Path: "/api/foo_123/345/more/things/to/come"},
	}
	// execute:
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// verify:
	if routeParams["entity"] != "foo_123" {
		t.Errorf("Route param for 'entity' do not match: should be: %s,  but is %s", "foo_123", routeParams["entity"])
	}
	if routeParams["id"] != "345" {
		t.Errorf("Route param for 'id' do not match: should be: %s,  but is %s", "345", routeParams["id"])
	}
	if routeParams["urlTail"] != "more/things/to/come" {
		t.Errorf("Route param for 'urlTail' do not match: should be: %s,  but is %s", "more/things/to/come", routeParams["urlTail"])
	}
}
