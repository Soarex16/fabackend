package handlers

import (
	"net/http"

	ctx "github.com/soarex16/fabackend/context"
)

// AddUserAchievement - return handler for route
// POST /achievements/{userId}
func AddUserAchievement(appCtx *ctx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}
}

// GetUserAchievements - return handler for route
// GET /achievements/{userId}
func GetUserAchievements(appCtx *ctx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}
}
