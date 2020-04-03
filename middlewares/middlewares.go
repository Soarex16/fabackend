package middlewares

import (
	"context"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// RequestContextKey - type definition for key of injected data into request.Context
type RequestContextKey string

const (
	RequestIDContextKey = "requestId"
	RequestIDHeader     = "X-Request-ID"
)

// RequestIDContext - structure of injected request id context
type RequestIDContext struct {
	ID uuid.UUID
}

func RequestID(next http.Handler) http.Handler {
	logrus.Error("RID")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID, parseErr := uuid.Parse(r.Header.Get(RequestIDHeader))

		if parseErr != nil {
			reqID = uuid.New()
		}

		logrus.Error("TEST!!!!")

		ctx := context.WithValue(
			r.Context(),
			RequestIDContextKey,
			RequestIDContext{
				ID: reqID,
			},
		)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Logging - middleware for logging requests
func Logging(next http.Handler) http.Handler {
	logrus.Error("LOGGING")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		reqID := r.Context().Value(RequestIDContextKey).(RequestIDContext).ID.String()

		logrus.WithFields(logrus.Fields{
			"method":         r.Method,
			"uri":            r.RequestURI,
			"requestId":      reqID,
			"processingTime": time.Since(start),
		}).Infof("Processed %v", r.RequestURI)
	})
}

// AuthorizationContextKey - key to get from request context
const AuthorizationContextKey RequestContextKey = "auth"

// AuthorizationContext - information about user
type AuthorizationContext struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int64
	UserID       uuid.UUID
}

// Authorization - wraps handler with authorization context and checks private routes
func Authorization(next http.Handler, private bool) http.Handler {
	logrus.Error("AUTH")
	//TODO
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
