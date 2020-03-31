package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/soarex16/fabackend/middlewares"
)

// NewRouter - constructs new router from Routes
func NewRouter(routes *Routes) *mux.Router {
	router := mux.NewRouter()

	router.Use(middlewares.Logging)

	for _, route := range *routes {
		var handler http.Handler
		handler = route.HandlerFunc

		handler = middlewares.Authorization(handler, route.Private)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
