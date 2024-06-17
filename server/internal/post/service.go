package post

import (
	"database/sql"
	"errors"

	"github.com/dsypasit/social-clone/server/internal/user"
	"github.com/google/uuid"
)

var ErrUserNotFound = errors.New("user not found")

type IUserServiceForPost interface {
	GetUserByUUID(s string) (user.User, error)
}

type IPostRepository interface {
	CreatePost(PostCreated) (int64, error)
	GetPostsByUserUUID(string) ([]PostResponse, error)
	GetPosts() ([]PostResponse, error)
}

type PostService struct {
	postRepo    IPostRepository
	userService IUserServiceForPost
}

func NewPostService(postRepo IPostRepository, userService IUserServiceForPost) *PostService {
	return &PostService{postRepo, userService}
}

func (s *PostService) CreatePost(p PostCreated) (int64, error) {
	_, err := s.userService.GetUserByUUID(p.UserUUID)
	if err == user.ErrUserNotFound {
		return 0, ErrUserNotFound
	}
	p.UUID = uuid.NewString()
	return s.postRepo.CreatePost(p)
}

func (s *PostService) GetPostsByUserUUID(userUUID string) ([]PostResponse, error) {
	if userUUID == "" {
		posts, err := s.postRepo.GetPosts()
		if err == sql.ErrNoRows {
			return nil, ErrNoRows
		}
		return posts, err
	}

	_, err := s.userService.GetUserByUUID(userUUID)
	if err == user.ErrUserNotFound {
		return nil, ErrNoRows
	}
	posts, err := s.postRepo.GetPostsByUserUUID(userUUID)
	if err == sql.ErrNoRows {
		return nil, ErrNoRows
	}
	return posts, err
}

func (s *PostService) GetPosts() ([]PostResponse, error) {
	posts, err := s.postRepo.GetPosts()
	if err == sql.ErrNoRows {
		return nil, ErrNoRows
	}
	return posts, err
}
