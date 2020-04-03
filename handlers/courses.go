package handlers

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/soarex16/fabackend/sql"
	"net/http"
)

type CoursesHandler struct {
	Handler
	Courses *sql.CoursesStore
}

// GetAllCources - return handler for route
// GET /courses
func (h *CoursesHandler) GetAllCources(w http.ResponseWriter, r *http.Request) {
	courses, err := h.Courses.GetAll()

	reqID := h.RequestID(r).String()

	if err != nil {
		logrus.
			WithField("requestId", reqID).
			Errorf("Error while processing request: %v", err)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(courses)

	if err != nil {
		logrus.
			WithField("requestId", reqID).
			WithField("entity", courses).
			Errorf("Error while serializing courses: %v", err)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(bytes)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
