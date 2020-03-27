package middlewares

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const LogContextKey = "requestId"

type LogContext struct {
	Id uuid.UUID
}

// Middleware for logging requests
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		requestIdCtx := context.WithValue(r.Context(), LogContextKey, LogContext{Id: uuid.New()})
		r = r.WithContext(requestIdCtx)

		next.ServeHTTP(w, r)

		reqId := r.Context().Value(LogContextKey).(LogContext).Id.String()
		logrus.WithFields(logrus.Fields{
			"method":         r.Method,
			"uri":            r.RequestURI,
			"requestId":      reqId,
			"processingTime": time.Since(start),
		}).Infof("Processed %v", r.RequestURI)
	})
}
