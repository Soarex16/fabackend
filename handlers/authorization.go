package handlers

import (
	"net/http"

	ctx "github.com/soarex16/fabackend/context"
)

// LoginUser - return handler for route
// POST /users/login
func LoginUser(appCtx *ctx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}
}

// RefreshTokens - return handler for route
// POST /users/login/refresh
func RefreshTokens(appCtx *ctx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}
}
