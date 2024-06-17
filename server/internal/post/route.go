package post

import (
	"net/http"

	"github.com/dsypasit/social-clone/server/internal/auth"
	"github.com/dsypasit/social-clone/server/internal/middleware"
	"github.com/gorilla/mux"
)

type IPostHandler interface {
	CreatePost(http.ResponseWriter, *http.Request)
	GetPostsByUserUUID(http.ResponseWriter, *http.Request)
}

func RegisterPostRouter(router *mux.Router, postHandler IPostHandler, jwtService *auth.JwtService) {
	srouter := router.PathPrefix("/post").Subrouter()
	srouter.Use(middleware.AuthMiddleware(jwtService))

	srouter.HandleFunc("", postHandler.GetPostsByUserUUID).Methods(http.MethodGet)
	srouter.HandleFunc("", postHandler.CreatePost).Methods(http.MethodPost)
}
