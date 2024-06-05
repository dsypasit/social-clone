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
