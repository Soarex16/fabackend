package middlewares

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	ctx "github.com/soarex16/fabackend/context"
)

// LogContextKey - key to get from request context
const LogContextKey ctx.RequestContextKey = "requestId"

// LogContext - structure of injected logging context
type LogContext struct {
	ID uuid.UUID
}

// Logging - middleware for logging requests
func Logging(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		requestIDCtx := context.WithValue(
			r.Context(),
			LogContextKey,
			LogContext{
				ID: uuid.New(),
			},
		)
		r = r.WithContext(requestIDCtx)

		inner.ServeHTTP(w, r)

		reqID := r.Context().Value(LogContextKey).(LogContext).ID.String()
		logrus.WithFields(logrus.Fields{
			"method":         r.Method,
			"uri":            r.RequestURI,
			"requestId":      reqID,
			"processingTime": time.Since(start),
		}).Infof("Processed %v", r.RequestURI)
	})
}
