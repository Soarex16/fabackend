package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/soarex16/fabackend/middlewares"
)

// Route - type for defining all route information
type Route struct {
	Name    string
	Method  string
	Pattern string
	Private bool
	Handler http.Handler
}

// Routes - collection of routes
type Routes []Route

// RouteParams - stores parameters for route
type RouteParams map[string]string

// NewRouter - constructs new router from Routes
func NewRouter(routes *Routes) *httprouter.Router {
	router := httprouter.New()

	for _, route := range *routes {
		handler := route.Handler

		handler = middlewares.Logging(handler)
		handler = middlewares.RequestID(handler)

		router.Handler(
			route.Method,
			route.Pattern,
			handler,
		)
	}

	return router
}
