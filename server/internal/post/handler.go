package post

import (
	"errors"

	"github.com/dsypasit/social-clone/server/internal/comment"
)

var ErrNoRows = errors.New("no posts")

type PostHandler struct {
	postService    *PostService
	commentService *comment.CommentService
}

// TODO: add post handler method
func NewPostHandler(postService *PostService, commentService *comment.CommentService) *PostHandler {
	return &PostHandler{postService, commentService}
}
