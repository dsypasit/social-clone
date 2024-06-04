package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthHandler struct {
	authService AuthServiceInterface
}

type AuthServiceInterface interface {
	Signup(u User) (string, error)
	Login(u User) (string, error)
}

func NewAuthHandler(authService AuthServiceInterface) *AuthHandler {
	return &AuthHandler{authService}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginedUser User
	if err := json.NewDecoder(r.Body).Decode(&loginedUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "invalid structure format"})
		return
	}
	var empty User
	if loginedUser == empty {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "invalid structure format"})
		return
	}
	if loginedUser.Username == "" || loginedUser.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "username or password should not empty"})
		return
	}
	token, err := h.authService.Login(loginedUser)
	if err != nil {
		if err == ErrInvalidPassword {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"message": "invalid password"})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": fmt.Sprintf("%v", err)})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var newUser User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "invalid structure format"})
		return
	}
	var empty User
	if newUser == empty {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "invalid structure format"})
		return
	}
	if newUser.Username == "" || newUser.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "username or password should not empty"})
		return
	}
	token, err := h.authService.Signup(newUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": fmt.Sprintf("%v", err)})
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
