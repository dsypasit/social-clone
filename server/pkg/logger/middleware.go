package logger

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type LoggerKey string

var LogKey LoggerKey = "logger"

func Middleware(logger *zap.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return LogMiddleware(next, logger)
	}
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
	body       []byte
}

func (lrw *loggingResponseWriter) Write(p []byte) (n int, err error) {
	lrw.body = p
	return lrw.ResponseWriter.Write(p)
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func LogMiddleware(next http.Handler, logger *zap.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), LogKey, logger)
		lrw := &loggingResponseWriter{ResponseWriter: w}

		next.ServeHTTP(lrw, r.WithContext(ctx))
		logger.Info("From LogMiddleware",
			zap.String("METHOD", r.Method),
			zap.Any("PATH", r.URL),
			zap.Int("STATUS", lrw.statusCode),
			zap.String("Response", string(lrw.body)),
		)
	})
}
