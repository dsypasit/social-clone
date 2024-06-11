package post

import (
	"net/http"

	"github.com/gorilla/mux"
)

type IPostHandler interface {
	CreatePost(http.ResponseWriter, *http.Request)
	GetPostsByUserUUID(http.ResponseWriter, *http.Request)
}

func RegisterPostRouter(router *mux.Router, postHandler IPostHandler) {
	srouter := router.PathPrefix("/post").Subrouter()

	srouter.HandleFunc("", postHandler.GetPostsByUserUUID).Methods(http.MethodGet)
	srouter.HandleFunc("", postHandler.CreatePost).Methods(http.MethodPost)
}
