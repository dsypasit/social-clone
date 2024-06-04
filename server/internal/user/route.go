package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterUserRouter(router *mux.Router, userHandler *UserHandler) {
	userRouter := router.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/", userHandler.GetUsers).Methods(http.MethodGet)
}
