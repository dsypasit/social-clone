package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/dsypasit/social-clone/server/internal/auth"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

func AuthMiddleware(jwtService auth.JwtServiceInterface) mux.MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return Middleware(jwtService, h)
	}
}

func Middleware(jwtService auth.JwtServiceInterface, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authValue := r.Header.Get("Authorization")
		if authValue == "" {
			w.WriteHeader(http.StatusUnauthorized)

			json.NewEncoder(w).Encode(map[string]string{
				"message": "Missing authorization token",
			})
			return
		}

		parts := strings.SplitN(authValue, "Bearer ", 2)
		if len(parts) != 2 {
			w.WriteHeader(http.StatusUnauthorized)

			json.NewEncoder(w).Encode(map[string]string{
				"message": "Invalid authorization format",
			})
			return
		}
		token := parts[1]

		claim, err := jwtService.VerifyToken(token)
		if err != nil {
			if err == jwt.ErrTokenExpired {
				w.WriteHeader(http.StatusUnauthorized)

				json.NewEncoder(w).Encode(map[string]string{
					"message": "token expired",
				})
				return
			}

			w.WriteHeader(http.StatusUnauthorized)

			json.NewEncoder(w).Encode(map[string]string{
				"message": "Invalid authorization token",
			})
			return
		}

		// insert uuid value to context
		ctx := context.WithValue(r.Context(), "userUUID", claim.UserUUID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
