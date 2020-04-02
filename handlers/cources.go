package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/soarex16/fabackend/middlewares"
	"github.com/soarex16/fabackend/repos"

	ctx "github.com/soarex16/fabackend/context"
)

// GetAllCources - return handler for route
// GET /courses
func GetAllCources(appCtx *ctx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		courses, err := repos.GetAllCourses(appCtx.DB)

		reqID := r.Context().Value(middlewares.LogContextKey).(middlewares.LogContext).ID.String()

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
}
