package middlewares

const AuthorizationContextKey = "auth"

type AuthorizationContext struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int64
}
