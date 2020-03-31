package handlers

import (
	"net/http"

	ctx "github.com/soarex16/fabackend/context"
)

// CreateUser - return handler for route
// POST /users
func CreateUser(appCtx *ctx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}
}

// GetUserByName - return handler for route
// GET /users/{username}
func GetUserByName(appCtx *ctx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}
}

// UpdateUser - return handler for route
// PUT /users/{username}
func UpdateUser(appCtx *ctx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}
}

// DeleteUser - return handler for route
// DELETE /users/{username}
func DeleteUser(appCtx *ctx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}
}
