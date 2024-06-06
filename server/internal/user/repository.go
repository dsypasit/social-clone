package user

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

var (
	ErrDupUsername  = errors.New("duplicate username")
	ErrUserNotFound = errors.New("user not found")
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) CreateUser(u UserCreated) (int64, error) {
	if err := ur.IsDuplicateUsername(u.Username); err != nil {
		return 0, err
	}
	var id int64
	err := ur.db.QueryRow("INSERT INTO app_user (uuid, username, email, password) VALUES ($1, $2, $3, $4) RETURNING id",
		u.UUID, u.Username, u.Email, u.Password).Scan(&id)
	return id, err
}

func (ur *UserRepository) GetUserByUsername(username string) (User, error) {
	var u User
	err := ur.db.QueryRow("SELECT id, uuid, username, email, created_at FROM app_user WHERE username=$1", username).
		Scan(&u.ID, &u.UUID, &u.Username, &u.Email, &u.CreatedAt)
	return u, err
}

func (ur *UserRepository) GetUserByUUID(username string) (User, error) {
	var u User
	err := ur.db.QueryRow("SELECT id, uuid, username, email, created_at FROM app_user WHERE uuid = $1 AND deleted_at is NULL", username).
		Scan(&u.ID, &u.UUID, &u.Username, &u.Email, &u.CreatedAt)
	if err != nil && err != sql.ErrNoRows {
		return User{}, err
	}
	if err == sql.ErrNoRows {
		return User{}, ErrUserNotFound
	}
	return u, err
}

func (ur *UserRepository) IsDuplicateUsername(username string) error {
	var existUsername string
	err := ur.db.QueryRow("SELECT username FROM app_user WHERE username=$1", username).Scan(&existUsername)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if existUsername != "" {
		return ErrDupUsername
	}
	return nil
}

func (ur *UserRepository) GetPasswordByUsername(username string) (string, error) {
	var uPassword string
	err := ur.db.QueryRow("SELECT password FROM app_user WHERE username=$1", username).Scan(&uPassword)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	if err == sql.ErrNoRows {
		return "", ErrUserNotFound
	}
	return uPassword, nil
}

func (ur *UserRepository) GetUserUUIDByUsername(username string) (string, error) {
	var uUUID string
	err := ur.db.QueryRow("SELECT uuid FROM app_user WHERE username=$1", username).Scan(&uUUID)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	if err == sql.ErrNoRows {
		return "", ErrUserNotFound
	}
	return uUUID, nil
}

func (ur *UserRepository) UpdateEmailByUUID(uuid string, email string) error {
	result, err := ur.db.Exec("UPDATE app_user SET email=$1 WHERE uuid=$2", email, uuid)
	if err != nil {
		return err
	}
	numAffect, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if numAffect == 0 {
		return ErrUserNotFound
	}

	return nil
}
