package post

import "github.com/dsypasit/social-clone/server/internal/commend"

type PostHandler struct {
	postService    *PostService
	commendService *commend.CommendService
}

func NewPostHandler(postService *PostService, commendService *commend.CommendService) *PostHandler {
	return &PostHandler{postService, commendService}
}
