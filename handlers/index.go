package handlers

import (
	"fmt"
	"net/http"
)

type IndexHandler struct {
	Handler
}

// Index - stub for index route
func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	//TODO: rewrite as json
	fmt.Fprintf(w, "Hello World!")
}
