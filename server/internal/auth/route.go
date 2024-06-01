package auth

import "github.com/gorilla/mux"

func RegisterAuthRouter(router *mux.Router, authHandler *AuthHandler) {
	router.HandleFunc("/login", authHandler.Login)
	router.HandleFunc("/signup", authHandler.Signup)
}
