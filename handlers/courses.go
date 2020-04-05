package handlers

import (
	"github.com/soarex16/fabackend/stores"
	"net/http"
)

type CoursesHandler struct {
	Handler
	Courses stores.CoursesStore
}

// GetAllCources - return handler for route
// GET /courses
func (h *CoursesHandler) GetAllCources(w http.ResponseWriter, r *http.Request) {
	courses, err := h.Courses.GetAll()

	if err != nil {
		h.InternalServerError(w, r, err, "Error while courses data from db")
	}

	h.WriteJsonBody(w, r, courses)
}
