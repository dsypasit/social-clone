package auth

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var secretKey = "test"

type MockUserService struct {
	u User
}

func (us *MockUserService) CreateUser(u User) (int, error) {
	return 1, nil
}

func (us *MockUserService) GetPasswordByUsername(s string) (string, error) {
	return us.u.Password, nil
}

func (us *MockUserService) GetUserUUIDByUsername(s string) (string, error) {
	return us.u.Password, nil
}

func TestJwtService(t *testing.T) {
	t.Run("jwt service should generate token", func(t *testing.T) {
		jwtService := NewJwtService(secretKey)
		token, err := jwtService.GenerateToken("c78047dd-f0cc-4e59-adc0-15e589410cd1")
		assert.Nil(t, err, "error should be null")
		assert.NotNil(t, token, "token should be not null")
	})

	t.Run("jwt service should parse token", func(t *testing.T) {
		jwtService := NewJwtService(secretKey)
		newJwtClaim := AuthJWTClaim{
			UserUUID: "b16b09cf-34ff-45c3-b1d6-b523ebffb57b",
		}
		token, _ := jwtService.GenerateToken("b16b09cf-34ff-45c3-b1d6-b523ebffb57b")
		assert.NotNil(t, token, "err should be nil")

		claim, err := jwtService.VerifyToken(token)
		assert.Nil(t, err, "err should be nil")
		assert.Equal(t, newJwtClaim.UserUUID, claim.UserUUID, fmt.Sprintf("should be %v but got %v", newJwtClaim, claim))
	})
}

func TestJwtSignUp(t *testing.T) {
	t.Run("should signup work and return token", func(t *testing.T) {
		userService := MockUserService{}
		jwtService := NewJwtService(secretKey)
		authService := NewAuthService(&userService, jwtService)
		newUser := User{}
		token, err := authService.Signup(newUser)

		assert.Nil(t, err, "err should be nil")
		assert.NotNil(t, token, "token should not nil")
	})
}

func TestJwtLogin(t *testing.T) {
	t.Run("should signin work and return token", func(t *testing.T) {
		loginedUser := User{
			Username: "ong",
			Password: "1234",
		}

		userService := MockUserService{loginedUser}
		jwtService := NewJwtService(secretKey)
		authService := NewAuthService(&userService, jwtService)

		token, err := authService.Login(loginedUser)

		assert.Nil(t, err, "err should be nil")
		assert.NotNil(t, token, "token should not nil")
	})

	t.Run("should got invalid password", func(t *testing.T) {
		loginedUser := User{
			Username: "ong",
			Password: "1234",
		}

		userService := MockUserService{loginedUser}
		jwtService := NewJwtService(secretKey)
		authService := NewAuthService(&userService, jwtService)

		loginedUser = User{
			Username: "ong",
			Password: "12345",
		}

		_, err := authService.Login(loginedUser)

		assert.NotNil(t, err, "err should be nil")
		assert.Equal(t, ErrInvalidPassword, err, "err should be invalid password")
	})
}
