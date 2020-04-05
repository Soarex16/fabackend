package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/soarex16/fabackend/auth"
	"github.com/soarex16/fabackend/domain"
	"github.com/soarex16/fabackend/stores"
	"net/http"
	"time"
)

type AuthHandler struct {
	Handler
	JwtSecret []byte
	Users     stores.UsersStore
	Sessions  *auth.SessionStore
}

// Login - performs user authentication
// POST /auth/login
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// Note: в принципе, имеет смысл проверять, залогинился ли уже юзер,
	// но принято решение оставить возможность множественной авторизации (например, с двух разных устройств)

	usr := &domain.User{}
	err := json.NewDecoder(r.Body).Decode(usr)

	if err != nil {
		h.UnprocessableEntity(w, r)
		return
	}

	// validate user
	if !emailRegex.MatchString(usr.Email) || !pwdRegex.MatchString(usr.Password) {
		h.Unauthorized(w, r)
		return
	}

	// find in db
	dbUsr, err := h.Users.FindByEmail(usr.Email)

	if err != nil {
		h.InternalServerError(w, r, err, "Error while fetching user from db")
		return
	}

	if dbUsr.Password != usr.Password {
		h.Unauthorized(w, r)
		return
	}

	authResp, err := generateAuthCredentials(dbUsr.Username, h.JwtSecret)

	if err != nil {
		h.InternalServerError(w, r, err, fmt.Sprintf("Error while signing tokens for user %v", dbUsr.Username))
	}

	// fill available fo user routes
	accessedRoutes := map[string]string{
		fmt.Sprintf("/users/%v", dbUsr.Username):              "",
		fmt.Sprintf("/users/%v/achievements", dbUsr.Username): "",
	}

	h.Sessions.Add(&auth.Session{
		Credentials:    *authResp,
		UserID:         dbUsr.ID,
		AccessedRoutes: accessedRoutes,
	})

	h.WriteJsonBody(w, r, authResp)
}

type refreshTokensDto struct {
	RefreshToken string `json:"refreshToken"`
}

// RefreshTokens - session refreshing when token expires
// POST /auth/refresh
func (h *AuthHandler) RefreshTokens(w http.ResponseWriter, r *http.Request) {
	dto := &refreshTokensDto{}
	err := json.NewDecoder(r.Body).Decode(dto)

	if err != nil {
		h.UnprocessableEntity(w, r)
		return
	}

	// get prev session, if not exists - fail
	prevSession, ok := h.Sessions.GetByRefreshToken(dto.RefreshToken)

	if !ok {
		h.Unauthorized(w, r)
		return
	}

	claims := auth.Claims{}

	// parse token
	token, err := jwt.ParseWithClaims(dto.RefreshToken, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return h.JwtSecret, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			h.Unauthorized(w, r)
			return
		}

		h.BadRequest(w, r)
		return
	}

	if !token.Valid {
		h.Unauthorized(w, r)
		return
	}

	// if token expired
	if time.Now().Unix() > claims.ExpiresAt {
		h.Sessions.RemoveByRefreshToken(prevSession.RefreshToken)

		h.Unauthorized(w, r)
		return
	}

	// generate new tokens
	authResp, err := generateAuthCredentials(claims.Username, h.JwtSecret)

	if err != nil {
		h.InternalServerError(w, r, err, fmt.Sprintf("Error while signing tokens for user %v", claims.Username))
	}

	// remove prev session
	h.Sessions.RemoveByRefreshToken(prevSession.RefreshToken)
	// add new
	h.Sessions.Add(&auth.Session{
		Credentials:    *authResp,
		UserID:         prevSession.UserID,
		AccessedRoutes: prevSession.AccessedRoutes,
	})

	h.WriteJsonBody(w, r, authResp)
}

var accessTokenExpTime = time.Minute * 30
var refreshTokenExpTime = time.Hour * 24 * 14

func generateAuthCredentials(username string, jwtSecret []byte) (*auth.Credentials, error) {
	// init user session
	accessTokenExp := time.Now().Add(accessTokenExpTime).Unix()
	accessClaims := auth.Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			Subject:   "access",
			ExpiresAt: accessTokenExp,
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenStr, err := accessToken.SignedString(jwtSecret)

	if err != nil {
		return nil, err
	}

	refreshTokenExp := time.Now().Add(refreshTokenExpTime).Unix()
	refreshClaims := auth.Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			Subject:   "refresh",
			ExpiresAt: refreshTokenExp,
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS512, refreshClaims)
	refreshTokenStr, err := refreshToken.SignedString(jwtSecret)

	if err != nil {
		return nil, err
	}

	authResponse := &auth.Credentials{
		AccessToken:     accessTokenStr,
		AccessTokenExp:  accessTokenExp,
		RefreshToken:    refreshTokenStr,
		RefreshTokenExp: refreshTokenExp,
	}

	return authResponse, nil
}
