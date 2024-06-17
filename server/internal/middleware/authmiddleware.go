package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/dsypasit/social-clone/server/internal/auth"
	"github.com/dsypasit/social-clone/server/internal/share/util"
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
			util.SendJson(w, map[string]string{
				"message": "Missing authorization token",
			}, http.StatusUnauthorized)
			return
		}

		parts := strings.SplitN(authValue, "Bearer ", 2)
		if len(parts) != 2 {
			util.SendJson(w, map[string]string{
				"message": "Invalid authorization format",
			}, http.StatusUnauthorized)
			return
		}
		token := parts[1]

		claim, err := jwtService.VerifyToken(token)
		if err != nil {
			if err == jwt.ErrTokenExpired {
				util.SendJson(w, map[string]string{
					"message": "token expired",
				}, http.StatusUnauthorized)
				return
			}

			util.SendJson(w, map[string]string{
				"message": "Invalid authorization token",
			}, http.StatusUnauthorized)
			return
		}

		// insert uuid value to context
		ctx := context.WithValue(r.Context(), "userUUID", claim.UserUUID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
