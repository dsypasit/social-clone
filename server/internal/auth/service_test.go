package auth

import (
	"fmt"
	"testing"

	"github.com/dsypasit/social-clone/server/internal/share/util"
	"github.com/dsypasit/social-clone/server/internal/user"
	"github.com/stretchr/testify/assert"
)

var secretKey = "test"

type MockUserService struct {
	u      *user.UserCreated
	uLogin User
}

func (us *MockUserService) CreateUser(u user.UserCreated) (int64, error) {
	return 1, nil
}

func (us *MockUserService) GetPasswordByUsername(s string) (string, error) {
	return us.uLogin.Password, nil
}

func (us *MockUserService) GetUserUUIDByUsername(s string) (string, error) {
	return "b3c5d2af-5cd3-4164-979d-1dcc705411bc", nil
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
		newUser := user.UserCreated{}
		token, err := authService.Signup(newUser)

		assert.Nil(t, err, "err should be nil")
		assert.NotNil(t, token, "token should not nil")
	})
}

func TestJwtLogin(t *testing.T) {
	t.Run("should signin work and return token", func(t *testing.T) {
		loginedUser := User{
			Username: "ong",
		}
		loginedUser.Password, _ = util.GeneratePassword("1234")

		userService := MockUserService{nil, loginedUser}
		jwtService := NewJwtService(secretKey)
		authService := NewAuthService(&userService, jwtService)

		loginedUser.Password = "1234"
		token, err := authService.Login(loginedUser)

		assert.Nil(t, err, "err should be nil")
		assert.NotNil(t, token, "token should not nil")
	})

	t.Run("should got invalid password", func(t *testing.T) {
		loginedUser := User{
			Username: "ong",
			Password: "1234",
		}

		userService := MockUserService{nil, loginedUser}
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

func TestCheckToken(t *testing.T) {
	jService := NewJwtService(secretKey)
	userUUID := "53d84415-9650-4271-b701-93fd9ed14adf"
	token, _ := jService.GenerateToken(userUUID)

	jService = NewJwtServiceWithExpireDuration(secretKey, 0)
	expireToken, _ := jService.GenerateToken(userUUID)

	testTable := []struct {
		title string
		input string
		want  bool
	}{
		{"should return true", token, true},
		{"should return false", "", false},
		{"should return false cause expired", expireToken, false},
	}

	for _, v := range testTable {
		t.Run(v.title, func(t *testing.T) {
			jService = NewJwtService(secretKey)
			muService := MockUserService{}
			aService := NewAuthService(&muService, jService)
			actual := aService.CheckToken(v.input)
			assert.Equalf(t, v.want, actual, "Want %v but got %v", v.want, actual)
		})
	}
}
