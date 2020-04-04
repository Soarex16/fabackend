package handlers

import (
	"encoding/json"
	"github.com/soarex16/fabackend/domain"
	"github.com/soarex16/fabackend/sql"
	"net/http"
	"time"
)

type AchievementsHandler struct {
	Handler
	Achievements *sql.AchievementsStore
	Users        *sql.UsersStore
}

// AddUserAchievement - return handler for route
// POST users/:username/achievements
func (h *AchievementsHandler) AddUserAchievement(w http.ResponseWriter, r *http.Request) {
	username := h.RouteParams(r)["username"]

	if !loginRegex.MatchString(username) {
		h.NotFound(w, r)
		return
	}

	// validate achievement
	ach := &domain.Achievement{}
	err := json.NewDecoder(r.Body).Decode(&ach)

	if err != nil {
		h.UnprocessableEntity(w, r)
		return
	}

	modelErrs := make(ValidationErrors)

	if time.Now().Sub(ach.Date).Milliseconds() == 0 || time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC).Sub(ach.Date).Milliseconds() == 0 {
		modelErrs["date"] = "Date can't be more than current and less than 2000"
	}

	if ach.Price < 0 {
		modelErrs["price"] = "Price can't be less than 0"
	}

	if len(ach.Title) == 0 {
		modelErrs["title"] = "Title can't be empty"
	} else if len(ach.IconName) > 100 {
		modelErrs["title"] = "Title length must be less than 100"
	}

	if len(ach.Title) == 0 {
		modelErrs["description"] = "Description can't be empty"
	} else if len(ach.IconName) > 300 {
		modelErrs["description"] = "Description length must be less than 100"
	}

	if len(modelErrs) > 0 {
		h.ModelValidationError(w, r, &modelErrs)
		return
	}

	// find user
	usr, err := h.Users.FindByName(username)
	if err != nil {
		h.NotFound(w, r)
		return
	}

	_, err = h.Achievements.Add(usr.ID, ach)

	if err != nil {
		h.InternalServerError(w, r, err, "Error while performing db operation")
	}

	h.Created(w, r, "Achievement was successfully added")
}

// GetUserAchievements - return handler for route
// GET users/:username/achievements
func (h *AchievementsHandler) GetUserAchievements(w http.ResponseWriter, r *http.Request) {
	username := h.RouteParams(r)["username"]

	if !loginRegex.MatchString(username) {
		h.NotFound(w, r)
		return
	}

	_, err := h.Users.FindByName(username)
	if err != nil {
		h.NotFound(w, r)
		return
	}

	achievements, err := h.Achievements.GetByUsername(username)

	if err != nil {
		h.InternalServerError(w, r, err, "Error while querying data from DB")
		return
	}

	h.WriteJsonBody(w, r, achievements)
}
