package router

import (
	"context"
	"net/http"

	"github.com/bylexus/go-stdlib/log"
)

const METHOD_GET = "GET"
const METHOD_POST = "POST"
const METHOD_PUT = "PUT"
const METHOD_PATCH = "PATCH"
const METHOD_DELETE = "DELETE"
const METHOD_OPTIONS = "OPTIONS"
const METHOD_HEAD = "HEAD"
const METHOD_TRACE = "TRACE"
const METHOD_CONNECT = "CONNECT"
const METHOD_ANY = "ANY"

type MatchedRouteKeyType string

const MatchedRouteKey MatchedRouteKeyType = "bylexus/http/router/MatchedRoute"

// The Router type implements a http.Handler that is capable of a more sophisticated routing
// than the routing mechanism of the standard library. For example, it supports placeholders and
// regular expressions for matching part of the route.
type Router struct {
	logger *log.SeverityLogger
	routes []Route
}

func NewRouter(logger *log.SeverityLogger) Router {
	return Router{
		logger: logger,
		routes: make([]Route, 0),
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	matchedRoute := r.findRoute(req)
	r.logger.Debug("Routing URL: %s", req.URL.Path)
	if matchedRoute == nil {
		r.logger.Debug("No route found for %s", req.URL.Path)
		http.NotFound(w, req)
		return
	}
	r.logger.Debug("Route found for %s: %s", req.URL.Path, matchedRoute.Route.Pattern)
	r.logger.Debug("Route params: %#v", matchedRoute.Params)
	// inject MatchedRoute into context:
	ctx := context.WithValue(req.Context(), MatchedRouteKey, matchedRoute)
	req = req.WithContext(ctx)
	(*matchedRoute.Route.Handler).ServeHTTP(w, req)
}

func (r *Router) Add(method string, pattern string, handler http.Handler) error {
	route, err := NewRoute(method, pattern, &handler)
	if err != nil {
		return err
	}
	r.routes = append(r.routes, *route)
	return nil
}

func (r *Router) Get(pattern string, handler http.Handler) error {
	return r.Add(METHOD_GET, pattern, handler)
}
func (r *Router) Post(pattern string, handler http.Handler) error {
	return r.Add(METHOD_POST, pattern, handler)
}

func (r *Router) Put(pattern string, handler http.Handler) error {
	return r.Add(METHOD_PUT, pattern, handler)
}

func (r *Router) Delete(pattern string, handler http.Handler) error {
	return r.Add(METHOD_DELETE, pattern, handler)
}

func (r *Router) Head(pattern string, handler http.Handler) error {
	return r.Add(METHOD_HEAD, pattern, handler)
}

func (r *Router) Options(pattern string, handler http.Handler) error {
	return r.Add(METHOD_OPTIONS, pattern, handler)
}

func (r *Router) Trace(pattern string, handler http.Handler) error {
	return r.Add(METHOD_TRACE, pattern, handler)
}

func (r *Router) Connect(pattern string, handler http.Handler) error {
	return r.Add(METHOD_CONNECT, pattern, handler)
}

func (r *Router) Patch(pattern string, handler http.Handler) error {
	return r.Add(METHOD_PATCH, pattern, handler)
}

func (r *Router) Any(pattern string, handler http.Handler) error {
	return r.Add(METHOD_ANY, pattern, handler)
}

func (r *Router) findRoute(req *http.Request) *MatchedRoute {
	// First exact-match implementation, does just a 1:1 match of the url path
	for _, route := range r.routes {
		// 1. check method:
		if route.Method != req.Method && route.Method != METHOD_ANY {
			continue
		}
		// 2. check path:
		match := route.regexPattern.FindStringSubmatch(req.URL.Path)
		if match != nil {
			paramNames := route.regexPattern.SubexpNames()[1:]
			paramValues := match[1:]
			params := RouteParamsFromKeyValues(paramNames, paramValues)
			return &MatchedRoute{
				Route:  route,
				Params: params,
				URL:    *req.URL,
				Method: req.Method,
			}
		}
	}
	return nil
}

func GetMatchedRoute(ctx context.Context) MatchedRoute {
	matchedRoute, ok := ctx.Value(MatchedRouteKey).(*MatchedRoute)
	if !ok {
		return MatchedRoute{}
	}
	return *matchedRoute
}
