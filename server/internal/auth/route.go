package auth

import (
	"net/http"

	"github.com/gorilla/mux"
)

type AuthHandlerInterface interface {
	Login(http.ResponseWriter, *http.Request)
	Signup(http.ResponseWriter, *http.Request)
}

func RegisterAuthRouter(router *mux.Router, authHandler AuthHandlerInterface) {
	router.HandleFunc("/login", authHandler.Login).Methods(http.MethodPost)
	router.HandleFunc("/signup", authHandler.Signup).Methods(http.MethodPost)
}
