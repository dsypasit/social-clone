package post

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/dsypasit/social-clone/server/internal/share/util"
)

var (
	ErrNoRows         = errors.New("no posts")
	ErrInCompleteInfo = errors.New("incomplete information")
	ErrInvalidUUID    = errors.New("invalid uuid format")
)

type IPostService interface {
	CreatePost(p PostCreated) (int64, error)
	GetPostsByUserUUID(string) ([]PostResponse, error)
}

type PostHandler struct {
	postService IPostService
}

func NewPostHandler(postService IPostService) *PostHandler {
	return &PostHandler{postService}
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var newPost PostCreated
	errInvalidReq := util.BuildErrResponse("invalid request")
	if err := json.NewDecoder(r.Body).Decode(&newPost); err != nil {
		util.SendJson(w, errInvalidReq(err), http.StatusBadRequest)
		return
	}

	if v := r.Context().Value("userUUID"); v != nil {
		log.Println("uuid", v)
		if userUUID, ok := v.(string); ok {
			log.Println("assigend successful")
			newPost.UserUUID = userUUID
		} else {
			util.SendJson(w, errInvalidReq(ErrInvalidUUID), http.StatusBadRequest)
			return
		}
	} else {
		util.SendJson(w, errInvalidReq(ErrInvalidUUID), http.StatusUnauthorized)
		return
	}

	if newPost.UserUUID == "" || newPost.Content == "" || newPost.VisibilityTypeId == 0 {
		log.Println(newPost)
		util.SendJson(w, errInvalidReq(ErrInCompleteInfo), http.StatusBadRequest)
		return
	}

	if !util.IsValidUUID(newPost.UserUUID) {
		util.SendJson(w, errInvalidReq(ErrInCompleteInfo), http.StatusBadRequest)
		return
	}

	_, err := h.postService.CreatePost(newPost)
	if err != nil {
		errServicErr := util.BuildErrResponse("service error")
		util.SendJson(w, errServicErr(err), http.StatusInternalServerError)
		return
	}

	response := util.BuildResponse("created post successful!")
	util.SendJson(w, response, http.StatusCreated)
}

func (h *PostHandler) GetPostsByUserUUID(w http.ResponseWriter, r *http.Request) {
	userUUID := r.URL.Query().Get("useruuid")
	errInvalidReq := util.BuildErrResponse("invalid request")
	if userUUID == "" {
		util.SendJson(w, errInvalidReq(ErrInCompleteInfo), http.StatusBadRequest)
		return
	}

	if !util.IsValidUUID(userUUID) {
		util.SendJson(w, errInvalidReq(ErrInCompleteInfo), http.StatusBadRequest)
		return
	}

	posts, err := h.postService.GetPostsByUserUUID(userUUID)
	if err != nil {
		errRes := util.BuildErrResponse("service failed")
		util.SendJson(w, errRes(err), http.StatusInternalServerError)
		return
	}

	util.SendJson(w, posts, http.StatusOK)
}
