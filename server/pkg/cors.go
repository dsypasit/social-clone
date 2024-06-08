package pkg

import (
	"log"
	"net/http"

	"github.com/dsypasit/social-clone/server/pkg/logger"
	"go.uber.org/zap"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, content-type, X-Auth-Token, Authorization")
		if logger, ok := r.Context().Value(logger.LogKey).(*zap.Logger); ok {
			logger.Info("option ok")
		}
		if r.Method == http.MethodOptions {
			log.Println("option kub")
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
