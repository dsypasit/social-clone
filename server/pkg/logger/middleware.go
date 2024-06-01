package logger

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type LoggerKey string

var logKey LoggerKey = "logger"

func Middleware(logger *zap.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return LogMiddleware(next, logger)
	}
}

func LogMiddleware(next http.Handler, logger *zap.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), logKey, logger)

		logger.Info("From LogMiddleware", zap.String("METHOD", r.Method), zap.Any("PATH", r.URL))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
