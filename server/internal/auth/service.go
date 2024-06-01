package auth

import "github.com/dsypasit/social-clone/server/internal/user"

type AuthService struct {
	usrService *user.UserService
}

func NewAuthService(usrService *user.UserService) *AuthService {
	return &AuthService{}
}
