package router

import (
	"net/http"
)

type RouterHandler interface {
	ServeRoute(w http.ResponseWriter, r *http.Request, matchedRoute MatchedRoute) error
}

type RouterHandlerFunc func(w http.ResponseWriter, r *http.Request, matchedRoute MatchedRoute) error

func (r RouterHandlerFunc) ServeRoute(w http.ResponseWriter, req *http.Request, matchedRoute MatchedRoute) error {
	err := r(w, req, matchedRoute)
	if err != nil {
		// TODO: handle specific http error
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return err
}

func Handler(routerHandler RouterHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		matchedRoute := GetMatchedRoute(r.Context())
		routerHandler.ServeRoute(w, r, matchedRoute)
	})
}
