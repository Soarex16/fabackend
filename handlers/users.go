package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/soarex16/fabackend/domain"
	"github.com/soarex16/fabackend/stores"
	"net/http"
	"regexp"
)

type UsersHandler struct {
	Handler
	Users stores.UsersStore
}

var loginRegex, _ = regexp.Compile(`^[A-Za-z\d]{8,50}$`)
var emailRegex, _ = regexp.Compile(`^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)
var pwdRegex, _ = regexp.Compile(`[A-Fa-f0-9]{64}`)

// CreateUser - return handler for route
// POST /users
func (h *UsersHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	usr := &domain.User{}
	err := json.NewDecoder(r.Body).Decode(usr)

	if err != nil {
		h.UnprocessableEntity(w, r)
		return
	}

	// validate user
	if modelErrs := getUserValidationErrors(usr); len(modelErrs) > 0 {
		h.ModelValidationError(w, r, &modelErrs)
		return
	}

	// check if user with this name already exists
	_, err = h.Users.FindByName(usr.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			_, err = h.Users.Add(usr)

			if err != nil {
				h.InternalServerError(w, r, err, "Error while creating user")
			}

			h.Created(w, r, "User was successfully added")
		}
		h.InternalServerError(w, r, err, "Error while getting user from db")
		return
	}

	h.ModelValidationError(w, r, &ValidationErrors{
		"username": "User with this name already exists",
	})
	return
}

// GetUserByName - return handler for route
// GET /users/{username}
func (h *UsersHandler) GetUserByName(w http.ResponseWriter, r *http.Request) {
	username := h.RouteParams(r)["username"]

	if !loginRegex.MatchString(username) {
		h.NotFound(w, r)
		return
	}

	usr, err := h.Users.FindByName(username)
	if err != nil {
		h.NotFound(w, r)
		return
	}

	// because UsersStore also returns pwd
	usr.Password = ""
	h.WriteJsonBody(w, r, http.StatusOK, usr)
}

// UpdateUser - return handler for route
// PUT /users/{username}
func (h *UsersHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	username := h.RouteParams(r)["username"]

	if !loginRegex.MatchString(username) {
		h.NotFound(w, r)
		return
	}

	usr := &domain.User{}
	err := json.NewDecoder(r.Body).Decode(usr)

	if err != nil {
		h.UnprocessableEntity(w, r)
		return
	}

	// validate user
	if modelErrs := getUserValidationErrors(usr); len(modelErrs) > 0 {
		h.ModelValidationError(w, r, &modelErrs)
		return
	}

	_, err = h.Users.Update(usr)

	if err != nil {
		h.InternalServerError(w, r, err, "Error while updating user")
	}

	h.Ok(w, r)
}

// DeleteUser - return handler for route
// DELETE /users/{username}
func (h *UsersHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
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

	_, err = h.Users.Delete(username)

	if err != nil {
		h.InternalServerError(w, r, err, "Error while deleting user from DB")
	}

	h.Ok(w, r)
}

func getUserValidationErrors(usr *domain.User) ValidationErrors {
	modelErrs := make(ValidationErrors)

	if len(usr.Username) == 0 {
		modelErrs["username"] = "Username can't be empty"
	} else if len(usr.Username) < 8 || len(usr.Username) > 50 {
		modelErrs["username"] = "Username must contain at least one digit and be more than 8 characters and less than 50 characters"
	} else if !loginRegex.MatchString(usr.Username) {
		modelErrs["username"] = "Username can only contain alphanumeric characters digits and underlines"
	}

	if len(usr.Email) == 0 {
		modelErrs["email"] = "Email can't be empty"
	} else if !emailRegex.MatchString(usr.Email) {
		modelErrs["email"] = "Email must be valid"
	}

	if !pwdRegex.MatchString(usr.Password) {
		modelErrs["password"] = "Password must be valid SHA-256"
	}

	return modelErrs
}
