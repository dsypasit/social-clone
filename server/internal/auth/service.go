package auth

import (
	"errors"

	"github.com/dsypasit/social-clone/server/internal/share/util"
	"github.com/dsypasit/social-clone/server/internal/user"
	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidPassword = errors.New("invalid password")
	ErrUserNotFound    = errors.New("username not found")
)

type UserServiceForAuth interface {
	CreateUser(user.UserCreated) (int64, error)
	GetPasswordByUsername(string) (string, error)
	GetUserUUIDByUsername(string) (string, error)
}

type JwtServiceInterface interface {
	GenerateToken(userUUID string) (string, error)
	VerifyToken(token string) (*AuthJWTClaim, error)
}

type AuthService struct {
	usrService UserServiceForAuth
	jwtService *JwtService
}

func NewAuthService(usrService UserServiceForAuth, jwtService *JwtService) *AuthService {
	return &AuthService{usrService: usrService, jwtService: jwtService}
}

func (as *AuthService) Signup(u user.UserCreated) (string, error) {
	_, err := as.usrService.CreateUser(u)
	if err != nil {
		return "", err
	}
	uuid, err := as.usrService.GetUserUUIDByUsername(u.Username)
	if err != nil {
		return "", err
	}
	return as.jwtService.GenerateToken(uuid)
}

func (as *AuthService) Login(u User) (string, error) {
	pass, err := as.usrService.GetPasswordByUsername(u.Username)
	if err != nil {
		return "", err
	}

	isMatch, _ := util.VerifyPassword(u.Password, pass)
	if !isMatch {
		return "", ErrInvalidPassword
	}
	uuid, err := as.usrService.GetUserUUIDByUsername(u.Username)
	if err != nil {
		return "", err
	}

	return as.jwtService.GenerateToken(uuid)
}

type JwtService struct {
	secretKey string
}

func NewJwtService(secretKey string) *JwtService {
	return &JwtService{secretKey: secretKey}
}

func (jService *JwtService) GenerateToken(userUUID string) (string, error) {
	claims := AuthJWTClaim{userUUID, jwt.RegisteredClaims{}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims, nil)
	return token.SignedString([]byte(jService.secretKey))
}

func (jService *JwtService) VerifyToken(token string) (*AuthJWTClaim, error) {
	verifiedToken, err := jwt.ParseWithClaims(token, &AuthJWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jService.secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	claim, ok := verifiedToken.Claims.(*AuthJWTClaim)
	if !ok {
		return nil, errors.New("failed to parse token with claim")
	}
	return claim, nil
}
