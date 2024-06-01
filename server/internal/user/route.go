package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterUserRouter(router *mux.Router, userHandler *UserHandler) {
	router.HandleFunc("/user", userHandler.GetUsers).Methods(http.MethodGet)
}
