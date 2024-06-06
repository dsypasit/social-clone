package user

import (
	"github.com/gorilla/mux"
)

func RegisterUserRouter(router *mux.Router, userHandler *UserHandler) {
	_ = router.PathPrefix("/user").Subrouter()
}
