package handlers

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/soarex16/fabackend/domain"
	"github.com/soarex16/fabackend/sql"
	"net/http"
	"regexp"
)

type UsersHandler struct {
	Handler
	Users *sql.UsersStore
}

var loginRegex, _ = regexp.Compile(`^([a-zA-Z0-9]){5,50}$`)

// CreateUser - return handler for route
// POST /users
func (h *UsersHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	reqID := h.RequestID(r).String()

	usr := &domain.User{}
	err := json.NewDecoder(r.Body).Decode(&usr)

	if err != nil {
		logrus.
			WithField("requestId", reqID).
			Errorf("Error while reading request body: %v", err)

		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if !loginRegex.MatchString(usr.Username) {
		logrus.
			WithFields(logrus.Fields{
				"requestId": reqID,
				"field":     "username",
				"value":     usr.Username,
			}).Error("Error while creating user: incorrect username")

		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	emailRegex, _ := regexp.Compile(`^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)
	if !emailRegex.MatchString(usr.Email) {
		logrus.
			WithFields(logrus.Fields{
				"requestId": reqID,
				"field":     "email",
				"value":     usr.Email,
			}).Error("Error while creating user: incorrect email")

		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	pwdRegex, _ := regexp.Compile(`[A-Fa-f0-9]{64}`)
	if !pwdRegex.MatchString(usr.Password) {
		logrus.
			WithFields(logrus.Fields{
				"requestId": reqID,
				"field":     "password",
				"value":     usr.Username,
			}).Error("Error while creating user: password isn't in SHA256 format")

		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	_, err = h.Users.Add(usr)

	if err != nil {
		logrus.
			WithFields(logrus.Fields{
				"requestId": reqID,
				"username":  usr.Username,
			}).Errorf("Error while creating user: %v", err)
	}

	w.WriteHeader(http.StatusOK)

}

// GetUserByName - return handler for route
// GET /users/{username}
func (h *UsersHandler) GetUserByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

}

// UpdateUser - return handler for route
// PUT /users/{username}
func (h *UsersHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

}

// DeleteUser - return handler for route
// DELETE /users/{username}
func (h *UsersHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

}
