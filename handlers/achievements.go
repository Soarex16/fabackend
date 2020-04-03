package handlers

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/soarex16/fabackend/sql"
	"net/http"
)

type AchievementsHandler struct {
	Handler
	Achievements *sql.AchievementsStore
	Users        *sql.UsersStore
}

// AddUserAchievement - return handler for route
// POST users/:username/achievements
func (h *AchievementsHandler) AddUserAchievement(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// GetUserAchievements - return handler for route
// GET users/:username/achievements
func (h *AchievementsHandler) GetUserAchievements(w http.ResponseWriter, r *http.Request) {
	username := h.RouteParams(r)["username"]

	reqID := h.RequestID(r).String()

	if !loginRegex.MatchString(username) {
		logrus.
			WithField("requestId", reqID).
			Error("Error while processing request: incorrect username")

		w.WriteHeader(http.StatusNotFound)
		return
	}

	_, err := h.Users.FindByName(username)
	if err != nil {
		logrus.
			WithFields(logrus.Fields{
				"requestId": reqID,
				"username":  username,
			}).Errorf("Cannot find user: %v", err)

		w.WriteHeader(http.StatusNotFound)
		return
	}

	achievements, err := h.Achievements.GetByUsername(username)

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
}
