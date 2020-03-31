package middlewares

import (
	"net/http"

	"github.com/google/uuid"

	ctx "github.com/soarex16/fabackend/context"
)

// AuthorizationContextKey - key to get from request context
const AuthorizationContextKey ctx.RequestContextKey = "auth"

// AuthorizationContext - information about user
type AuthorizationContext struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int64
	UserID       uuid.UUID
}

// Authorization - wraps handler with authorization context and checks private routes
func Authorization(inner http.Handler, private bool) http.Handler {
	//TODO
	return inner
}
