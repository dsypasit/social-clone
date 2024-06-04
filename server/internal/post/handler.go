package post

import "github.com/dsypasit/social-clone/server/internal/comment"

type PostHandler struct {
	postService    *PostService
	commentService *comment.commentService
}

func NewPostHandler(postService *PostService, commentService *comment.commentService) *PostHandler {
	return &PostHandler{postService, commentService}
}
