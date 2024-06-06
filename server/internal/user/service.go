package user

import (
	"errors"

	"github.com/dsypasit/social-clone/server/internal/share/util"
	"github.com/google/uuid"
)

type IUserRepository interface {
	GetUserByUUID(string) (User, error)
	GetPasswordByUsername(string) (string, error)
	CreateUser(UserCreated) (int64, error)
	GetUserUUIDByUsername(string) (string, error)
}

type UserService struct {
	userRepo IUserRepository
}

func NewUserService(userRepo IUserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (us *UserService) GetUserByUUID(s string) (User, error) {
	return us.userRepo.GetUserByUUID(s)
}

func (us *UserService) GetPasswordByUsername(username string) (string, error) {
	return us.userRepo.GetPasswordByUsername(username)
}

func (us *UserService) CreateUser(newUser UserCreated) (int64, error) {
	var err error
	if !util.ValidateEmail(newUser.Email) {
		return 0, errors.New("invalid email")
	}
	newUser.UUID = uuid.New().String()
	newUser.Password, err = util.GeneratePassword(newUser.Password)
	if err != nil {
		return 0, err
	}
	return us.userRepo.CreateUser(newUser)
}

func (us *UserService) GetUserUUIDByUsername(username string) (string, error) {
	return us.userRepo.GetUserUUIDByUsername(username)
}
