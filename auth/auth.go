package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

// Credentials - represents user auth response
type Credentials struct {
	AccessToken    string `json:"accessToken"`
	AccessTokenExp int64  `json:"accessTokenExp"`

	RefreshToken    string `json:"refreshToken"`
	RefreshTokenExp int64  `json:"refreshTokenExp"`
}

type AccessedRoutes map[string]string

type Session struct {
	Credentials
	UserID         uuid.UUID
	AccessedRoutes AccessedRoutes
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GetBearerToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")

	if authHeader != "" && strings.HasPrefix(authHeader, "Bearer") {
		// skip "Bearer " and remove spaces
		return strings.Trim(authHeader[7:], " "), nil
	}

	return "", errors.New("can't find Authorization header")
}
