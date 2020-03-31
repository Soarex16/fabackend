package handlers

import (
	"net/http"

	ctx "github.com/soarex16/fabackend/context"
)

// GetAllCources - return handler for route
// GET /courses
func GetAllCources(appCtx *ctx.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}
}
