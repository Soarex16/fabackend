package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/soarex16/fabackend/middlewares"
)

// Route - type for defining all route information
type Route struct {
	Name        string
	Method      string
	Pattern     string
	Private     bool
	HandlerFunc http.HandlerFunc
}

// Routes - collection of routes
type Routes []Route

// RouteParams - stores parameters for route
type RouteParams map[string]string

// NewRouter - constructs new router from Routes
func NewRouter(routes *Routes) *httprouter.Router {
	router := httprouter.New()

	for _, route := range *routes {
		var handler http.Handler
		handler = route.HandlerFunc

		handler = middlewares.RequestID(handler)
		handler = middlewares.Logging(handler)
		handler = middlewares.Authorization(handler, route.Private)

		router.Handler(
			route.Method,
			route.Pattern,
			route.HandlerFunc,
		)
	}

	return router
}
