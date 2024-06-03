package auth

import "github.com/golang-jwt/jwt/v5"

type Auth struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password"`
}

type AuthJWTClaim struct {
	UserUUID string
	jwt.RegisteredClaims
}
