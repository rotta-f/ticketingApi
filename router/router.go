package router

import (
	"net/http"
	"regexp"
	"github.com/rotta-f/ticketingApi/utils"
)

type Route struct {
	Path        string
	Method      string
	HandlerFunc http.HandlerFunc
	rg          *regexp.Regexp
}

type Router struct {
	Routes []Route
}

func (r *Router) AddRoute(method string, path string, handlerFunc http.HandlerFunc) {
	route := Route{
		Method:      method,
		Path:        path,
		HandlerFunc: handlerFunc,
	}
	route.rg = regexp.MustCompile(`^` + route.Path + `/*$`)
	r.Routes = append(r.Routes, route)
}

func NewRouter() *Router {
	router := &Router{
		Routes: make([]Route, 100),
	}
	return router
}

func UseRouter(router *Router) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, route := range router.Routes {
			if r.Method == route.Method && route.rg.Match([]byte(r.URL.Path)) {
				route.HandlerFunc(w, r)
				return
			}
		}
		utils.WriteError(w, http.StatusNotFound, http.StatusText(http.StatusNotFound), "")
	}
}
