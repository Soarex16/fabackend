package handlers

import (
	"github.com/soarex16/fabackend/sql"
	"net/http"
)

type AuthHandler struct {
	Handler
	JwtSecret []byte
	Users     *sql.UsersStore
	//Sessions *SessionsStore
}

// Login - performs user authentication
// POST /auth/login
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	h.Ok(w, r)
}

// RefreshTokens - session refreshing when token expires
// POST /auth/refresh
func (h *AuthHandler) RefreshTokens(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
