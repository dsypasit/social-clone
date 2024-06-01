package auth

import "net/http"

type AuthHandler struct {
	authService *AuthService
}

func NewAuthHandler(authService *AuthService) *AuthHandler {
	return &AuthHandler{authService}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
}
