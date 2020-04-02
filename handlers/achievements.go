package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/soarex16/fabackend/repos"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	ctx "github.com/soarex16/fabackend/context"
	"github.com/soarex16/fabackend/middlewares"
)

// AddUserAchievement - return handler for route
// POST users/{userName}/achievements
func AddUserAchievement(appCtx *ctx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}
}

// GetUserAchievements - return handler for route
// GET users/{userName}/achievements
func GetUserAchievements(appCtx *ctx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := mux.Vars(r)["username"]

		reqID := r.Context().Value(middlewares.LogContextKey).(middlewares.LogContext).ID.String()

		if !loginRegex.MatchString(username) {
			logrus.
				WithField("requestId", reqID).
				Error("Error while processing request: incorrect username")

			w.WriteHeader(http.StatusNotFound)
			return
		}

		_, err := repos.FindUserByName(appCtx.DB, username)
		if err != nil {
			logrus.
				WithFields(logrus.Fields{
					"requestId": reqID,
					"username":  username,
				}).Errorf("Cannot find user: %v", err)

			w.WriteHeader(http.StatusNotFound)
			return
		}

		achievements, err := repos.GetUserAchievements(appCtx.DB, username)

		if err != nil {
			logrus.
				WithField("requestId", reqID).
				Errorf("Error while processing request: %v", err)

			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		bytes, err := json.Marshal(achievements)

		if err != nil {
			logrus.
				WithField("requestId", reqID).
				WithField("entity", achievements).
				Errorf("Error while serializing courses: %v", err)

			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(bytes)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}
}
