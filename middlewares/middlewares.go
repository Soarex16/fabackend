package middlewares

import (
	"context"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/soarex16/fabackend/auth"
	"net/http"
	"strings"
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
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID, parseErr := uuid.Parse(r.Header.Get(RequestIDHeader))

		if parseErr != nil {
			reqID = uuid.New()
		}

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
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		reqID := r.Context().Value(RequestIDContextKey).(RequestIDContext).ID.String()

		logrus.WithFields(logrus.Fields{
			"method":         r.Method,
			"uri":            r.URL.Path,
			"requestId":      reqID,
			"processingTime": time.Since(start),
		}).Infof("Processed %v", r.URL.Path)
	})
}

// AuthorizationContextKey - key to get from request context
const AuthorizationContextKey RequestContextKey = "auth"

// AuthorizationContext - information about user
type AuthorizationContext struct {
	Session *auth.Session
}

// Authorization - wraps handler with authorization context and checks private routes
func Authorization(next http.Handler, store *auth.SessionStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := auth.GetBearerToken(r)

		// fast check
		if err != nil || token == "" {
			logrus.
				WithField("path", r.URL.Path).
				Warn("Attempt to access private route without authorization")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		session, ok := store.GetByAccessToken(token)
		if !ok {
			logrus.
				WithField("path", r.URL.Path).
				Warn("Attempt to access private route with invalid token")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if time.Now().Unix() > session.AccessTokenExp {
			// remove token from storage
			store.RemoveByAccessToken(token)

			logrus.
				WithField("path", r.URL.Path).
				Warn("Attempt to access private route with expired token")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// проверяем, что юзеру можно на этот раут
		canPass := false
		for route := range session.AccessedRoutes {
			if strings.Contains(route, r.URL.Path) {
				canPass = true
			}
		}

		if !canPass {
			logrus.
				WithField("path", r.URL.Path).
				Warn("none of the paths allowed to the user matches the requested path")

			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(
			r.Context(),
			AuthorizationContextKey,
			AuthorizationContext{
				Session: session,
			},
		)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")

		next.ServeHTTP(w, r)
	})
}
