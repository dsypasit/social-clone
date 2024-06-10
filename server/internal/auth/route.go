package auth

import (
	"net/http"

	"github.com/gorilla/mux"
)

type AuthHandlerInterface interface {
	Login(http.ResponseWriter, *http.Request)
	Signup(http.ResponseWriter, *http.Request)
	CheckToken(http.ResponseWriter, *http.Request)
}

func RegisterAuthRouter(router *mux.Router, authHandler AuthHandlerInterface) {
	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/login", authHandler.Login).Methods(http.MethodPost)
	authRouter.HandleFunc("/signup", authHandler.Signup).Methods(http.MethodPost)
	authRouter.HandleFunc("/checktoken", authHandler.CheckToken).Methods(http.MethodPost)
}
