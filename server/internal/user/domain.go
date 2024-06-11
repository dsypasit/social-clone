package user

import "time"

type User struct {
	ID        int       `json:"-" db:"id"`
	UUID      string    `json:"uuid" db:"uuid"`
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
}

type UserCreated struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	UUID     string `json:"uuid" db:"uuid"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
}
