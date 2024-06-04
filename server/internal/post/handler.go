package post

import "github.com/dsypasit/social-clone/server/internal/comment"

type PostHandler struct {
	postService    *PostService
	commentService *comment.CommentService
}

func NewPostHandler(postService *PostService, commentService *comment.CommentService) *PostHandler {
	return &PostHandler{postService, commentService}
}
