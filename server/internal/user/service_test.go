package user

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type MockUserRepo struct {
	u   User
	err error
}

func (m *MockUserRepo) GetUserByUUID(s string) (User, error) {
	return m.u, m.err
}

func (m *MockUserRepo) GetPasswordByUsername(s string) (string, error) {
	return m.u.Password, m.err
}

func (m *MockUserRepo) CreateUser(newUser UserCreated) (int64, error) {
	return 1, nil
}

func (m *MockUserRepo) GetUserUUIDByUsername(uname string) (string, error) {
	return "da198c46-5b53-4988-986c-00df8f0a4086", nil
}

func TestServiceGetUserByUUID(t *testing.T) {
	want := User{
		ID:        1,
		UUID:      "eb2b0677-e035-45bd-8c25-54d03d6d1c11",
		Username:  "john even",
		Email:     "a@gmail.com",
		CreatedAt: time.Now(),
	}
	testTable := []struct {
		title   string
		id      string
		want    User
		wantErr error
	}{
		{"should get user", "eb2b0677-e035-45bd-8c25-54d03d6d1c11", want, nil},
	}

	for _, v := range testTable {
		t.Run(v.title, func(t *testing.T) {
			mRepo := MockUserRepo{want, nil}
			userService := NewUserService(&mRepo)
			actual, err := userService.GetUserByUUID(v.id)

			assert.Equalf(t, v.wantErr, err, "Unexpected error: %v", err)
			assert.Equalf(t, v.want, actual, "want %v but got %v", v.want, actual)
		})
	}
}

func TestServiceGetPasswordByUsername(t *testing.T) {
	want := User{
		ID:        1,
		UUID:      "eb2b0677-e035-45bd-8c25-54d03d6d1c11",
		Username:  "username",
		Email:     "a@gmail.com",
		Password:  "asdf;lkj",
		CreatedAt: time.Now(),
	}
	testTable := []struct {
		title    string
		username string
		want     string
		wantErr  error
	}{
		{"should get password", "username", want.Password, nil},
	}

	for _, v := range testTable {
		t.Run(v.title, func(t *testing.T) {
			mRepo := MockUserRepo{want, nil}
			userService := NewUserService(&mRepo)
			actual, err := userService.GetPasswordByUsername(v.username)

			assert.Equalf(t, v.wantErr, err, "Unexpected error: %v", err)
			assert.Equalf(t, v.want, actual, "want %v but got %v", v.want, actual)
		})
	}
}

func TestServiceCreateUser(t *testing.T) {
	testTable := []struct {
		title   string
		newUser UserCreated
		want    int64
		wantErr error
	}{
		{
			"should return last id",
			UserCreated{"35707a9d-a346-4cd4-ba0c-1dfd0b9ba96e", "abc123", "a@gmail.com", "asldkfjasjkdlf"},
			1,
			nil,
		},
	}

	for _, v := range testTable {
		t.Run(v.title, func(t *testing.T) {
			mRepo := MockUserRepo{}
			uService := NewUserService(&mRepo)
			actual, err := uService.CreateUser(v.newUser)

			assert.Equalf(t, v.wantErr, err, "Unexpected error : %v", err)
			assert.Equalf(t, v.want, actual, "Want id %v but got %v", v.want, actual)
		})
	}
}

func TestServiceGetUserUUIDByUsername(t *testing.T) {
	want := "da198c46-5b53-4988-986c-00df8f0a4086"
	input := "asdf"

	mRepo := MockUserRepo{}
	us := NewUserService(&mRepo)
	uuid, err := us.GetUserUUIDByUsername(input)
	assert.Equal(t, nil, err, "Unexpected error: %v", err)
	assert.Equal(t, want, uuid, "Want %v but got %v", want, uuid)
}
