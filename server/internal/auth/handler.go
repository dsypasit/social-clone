package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/dsypasit/social-clone/server/internal/share/util"
	"github.com/dsypasit/social-clone/server/internal/user"
)

type AuthHandler struct {
	authService AuthServiceInterface
}

type AuthServiceInterface interface {
	Signup(u user.UserCreated) (string, error)
	Login(u User) (string, error)
	CheckToken(token string) bool
}

func NewAuthHandler(authService AuthServiceInterface) *AuthHandler {
	return &AuthHandler{authService}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginedUser User
	if err := json.NewDecoder(r.Body).Decode(&loginedUser); err != nil {
		util.SendJson(w, map[string]string{"message": "invalid structure format"}, http.StatusBadRequest)
		return
	}
	var empty User
	if loginedUser == empty {
		util.SendJson(w, map[string]string{"message": "invalid structure format"}, http.StatusBadRequest)
		return
	}
	if loginedUser.Username == "" || loginedUser.Password == "" {
		util.SendJson(w, map[string]string{"message": "username or password empty or invalid email format"}, http.StatusBadRequest)
		return
	}
	token, err := h.authService.Login(loginedUser)
	if err != nil {
		if err == ErrInvalidPassword {
			util.SendJson(w, map[string]string{"message": "invalid password"}, http.StatusBadRequest)
			return
		}

		if err == ErrUserNotFound {
			util.SendJson(w, map[string]string{"message": "user not found"}, http.StatusBadRequest)
			return
		}
		util.SendJson(w, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}
	util.SendJson(w, map[string]string{"token": token}, http.StatusOK)
}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var newUser user.UserCreated
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		util.SendJson(w, map[string]string{"message": "invalid structure format"}, http.StatusBadRequest)
		return
	}
	var empty user.UserCreated
	if newUser == empty {
		util.SendJson(w, map[string]string{"message": "invalid structure format"}, http.StatusBadRequest)
		return
	}
	if newUser.Username == "" || newUser.Password == "" || !util.ValidateEmail(newUser.Email) {
		util.SendJson(w, map[string]string{"message": "username or password empty or invalid email format"}, http.StatusBadRequest)
		return
	}
	token, err := h.authService.Signup(newUser)
	if err != nil {
		if err == user.ErrDupUsername {
			util.SendJson(w, map[string]string{"message": "duplicate username"}, http.StatusBadRequest)
			return
		}
		util.SendJson(w, map[string]string{"message": fmt.Sprintf("%v", err)}, http.StatusInternalServerError)
		return
	}
	util.SendJson(w, map[string]string{"token": token}, http.StatusCreated)
}

func (h *AuthHandler) CheckToken(w http.ResponseWriter, r *http.Request) {
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

	if !h.authService.CheckToken(token) {
		res := util.BuildResponse("Unauthorized")
		util.SendJson(w, res, http.StatusUnauthorized)
		return
	}
	util.SendJson(w, nil, http.StatusNoContent)
}
