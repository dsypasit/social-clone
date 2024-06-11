package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type IUserHandler interface {
	GetUserByUsername(w http.ResponseWriter, r *http.Request)
}

func RegisterUserRouter(router *mux.Router, userHandler IUserHandler) {
	s := router.PathPrefix("/user").Subrouter()
	s.HandleFunc("", userHandler.GetUserByUsername)
}
