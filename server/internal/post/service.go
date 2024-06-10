package post

import (
	"database/sql"

	"github.com/google/uuid"
)

type IPostRepository interface {
	CreatePost(Post) (int64, error)
	GetPostsByUserUUID(string) ([]PostResponse, error)
}

type PostService struct {
	postRepo IPostRepository
}

func NewPostService(postRepo IPostRepository) *PostService {
	return &PostService{postRepo}
}

func (s *PostService) CreatePost(p Post) (int64, error) {
	p.UUID = uuid.NewString()
	return s.postRepo.CreatePost(p)
}

func (s *PostService) GetPostsByUserUUID(userUUID string) ([]PostResponse, error) {
	posts, err := s.postRepo.GetPostsByUserUUID(userUUID)
	if err == sql.ErrNoRows {
		return nil, ErrNoRows
	}
	return posts, err
}
