package handlers

import (
	"net/http"
)

type IndexHandler struct {
	Handler
}

// Index - stub for index route
func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	h.WriteJsonBody(w, r, "Hello World!")
}
