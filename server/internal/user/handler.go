package user

import (
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	userSrv *UserService
}

func NewUserHandler(userSrv *UserService) *UserHandler {
	return &UserHandler{userSrv}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	result := []User{
		{},
		{},
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
