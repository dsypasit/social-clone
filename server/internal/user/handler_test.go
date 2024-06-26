package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dsypasit/social-clone/server/internal/share/util"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

type MockUserService struct {
	u   User
	err error
}

func (m *MockUserService) GetUserByUUID(string) (User, error) {
	return m.u, m.err
}

func (m *MockUserService) CreateUser(UserCreated) (int64, error) {
	return 1, m.err
}

func (m *MockUserService) GetUserByUsername(string) (User, error) {
	return m.u, m.err
}

func TestHandlerGetUserByUUID(t *testing.T) {
	passQuery, _ := util.GeneratePassword("wow")
	userQuery := User{
		ID:        1,
		UUID:      "98da0985-1b0c-47f5-95c5-ca63b5c4df35",
		Username:  "abc",
		Email:     "a@gmail.com",
		Password:  passQuery,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	testTable := []struct {
		title      string
		uuidParams string
		serviceErr error
		wantBody   User
		wantCode   int
	}{
		{"should return user", "98da0985-1b0c-47f5-95c5-ca63b5c4df35", nil, User{
			UUID:      "98da0985-1b0c-47f5-95c5-ca63b5c4df35",
			Username:  "abc",
			Email:     "a@gmail.com",
			CreatedAt: userQuery.CreatedAt,
		}, http.StatusOK},
	}

	for _, v := range testTable {
		t.Run(v.title, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "/", nil)
			assert.Nilf(t, err, "Unexpected error: %v", err)
			req = mux.SetURLVars(req, map[string]string{
				"uuid": v.uuidParams,
			})

			rec := httptest.NewRecorder()

			mService := MockUserRepo{
				userQuery,
				v.serviceErr,
			}

			uh := NewUserHandler(&mService)
			uh.GetUserByUUID(rec, req)

			var response User
			json.NewDecoder(rec.Body).Decode(&response)

			expectedCreatedAt := v.wantBody.CreatedAt.UnixMilli()
			actualCreatedAt := response.CreatedAt.UnixMilli()
			v.wantBody.CreatedAt = time.Time{}
			response.CreatedAt = time.Time{}

			// assertion
			assert.Equalf(t, v.wantCode, rec.Code, "Want %v but got %v", v.wantCode, rec.Code)
			assert.Equal(t, expectedCreatedAt, actualCreatedAt, "Want %v but got %v", expectedCreatedAt, actualCreatedAt)
			assert.Equalf(t, v.wantBody, response, "Want %v but got %v", v.wantCode, rec.Code)
		})
	}
}

func TestHandlerGetUserByUUID_Error(t *testing.T) {
	passQuery, _ := util.GeneratePassword("wow")
	userQuery := User{
		ID:        1,
		UUID:      "98da0985-1b0c-47f5-95c5-ca63b5c4df35",
		Username:  "abc",
		Email:     "a@gmail.com",
		Password:  passQuery,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	testTable := []struct {
		title      string
		uuidParams string
		serviceErr error
		wantBody   map[string]string
		wantCode   int
	}{
		{
			"should return bad request",
			"", nil,
			map[string]string{"message": "invalid uuid"},
			http.StatusBadRequest,
		}, {
			"should return not found",
			"089ff020-d4bd-4aa3-bf18-3a7fabb10dc5", ErrUserNotFound,
			map[string]string{"message": "user not found"},
			http.StatusNotFound,
		},
	}

	for _, v := range testTable {
		t.Run(v.title, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "/", nil)
			assert.Nilf(t, err, "Unexpected error: %v", err)
			req = mux.SetURLVars(req, map[string]string{
				"uuid": v.uuidParams,
			})

			rec := httptest.NewRecorder()

			mService := MockUserRepo{
				userQuery,
				v.serviceErr,
			}

			uh := NewUserHandler(&mService)
			uh.GetUserByUUID(rec, req)

			var response map[string]string
			json.NewDecoder(rec.Body).Decode(&response)

			// assertion
			assert.Equalf(t, v.wantCode, rec.Code, "Want %v but got %v", v.wantCode, rec.Code)
			assert.Equalf(t, v.wantBody, response, "Want %v but got %v", v.wantCode, rec.Code)
		})
	}
}

func TestHandlerCreateUser(t *testing.T) {
	testTable := []struct {
		title          string
		userCreated    UserCreated
		wantBody       map[string]string
		wantStatusCode int
	}{
		{
			"should create succes",
			UserCreated{
				Username: "ong",
				Email:    "a@gmail.com",
				Password: "abcd123",
			},
			map[string]string{"message": "User created successfully!"},
			http.StatusCreated,
		},
		{
			"invalid mail",
			UserCreated{
				Username: "ong",
				Email:    "a.gmail.com",
				Password: "abcd123",
			},
			map[string]string{"message": "Failed to create user", "error": "invalid email format"},
			http.StatusBadRequest,
		},
	}

	for _, v := range testTable {
		t.Run(v.title, func(t *testing.T) {
			mService := MockUserService{}
			uh := NewUserHandler(&mService)

			body, _ := json.Marshal(v.userCreated)
			req, _ := http.NewRequest(http.MethodPost, "/", bytes.NewReader(body))
			rec := httptest.NewRecorder()

			uh.CreateUser(rec, req)

			var response map[string]string
			json.NewDecoder(rec.Body).Decode(&response)
			assert.Equalf(t, v.wantStatusCode, rec.Code, "want %v but got %v", v.wantStatusCode, rec.Code)
			assert.Equalf(t, v.wantBody, response, "want %v but got %v", v.wantBody, response)
		})
	}
}

func TestHandlerGetUserByUsername(t *testing.T) {
	testTable := []struct {
		title      string
		input      string
		serviceErr error
		want       UserResponse
		wantStatus int
	}{
		{"should get user succesful", "ong2", nil, UserResponse{
			UUID:     "0e11819f-e780-4129-ad06-9c5634d0f054",
			Username: "ong2",
			Email:    "a@gmail.com",
		}, http.StatusOK},
	}

	for _, v := range testTable {
		t.Run(v.title, func(t *testing.T) {
			mService := MockUserService{
				u: User{
					UUID:     v.want.UUID,
					Email:    v.want.Email,
					Username: v.want.Username,
				},
				err: v.serviceErr,
			}
			h := NewUserHandler(&mService)

			url := fmt.Sprintf("/?username=%v", v.input)
			req, _ := http.NewRequest(http.MethodGet, url, nil)
			rec := httptest.NewRecorder()

			h.GetUserByUsername(rec, req)
			var res UserResponse
			json.NewDecoder(rec.Body).Decode(&res)
			assert.Equalf(t, v.wantStatus, rec.Code, "Want %v but got %v", v.wantStatus, rec.Code)
			assert.Equalf(t, v.want, res, "Want %v but got %v", v.want, res)
		})
	}
}

func TestHandlerGetUserByUsername_Error(t *testing.T) {
	testTable := []struct {
		title      string
		input      string
		serviceErr error
		want       map[string]string
		wantStatus int
	}{
		{"should invalid request", "", nil, util.BuildErrResponse("invalid request")(nil), http.StatusBadRequest},
		{
			"shoud service failure", "ong2", errors.New("service failed"), util.BuildErrResponse("service failure")(nil),
			http.StatusInternalServerError,
		},
		{
			"shoud user not found", "ong2", ErrUserNotFound, util.BuildErrResponse("user not found")(nil),
			http.StatusNotFound,
		},
	}

	for _, v := range testTable {
		t.Run(v.title, func(t *testing.T) {
			mService := MockUserService{
				u:   User{},
				err: v.serviceErr,
			}
			h := NewUserHandler(&mService)

			url := fmt.Sprintf("/?username=%v", v.input)
			req, _ := http.NewRequest(http.MethodGet, url, nil)
			rec := httptest.NewRecorder()

			h.GetUserByUsername(rec, req)
			var res map[string]string
			json.NewDecoder(rec.Body).Decode(&res)
			assert.Equalf(t, v.wantStatus, rec.Code, "Want %v but got %v", v.wantStatus, rec.Code)
			assert.Equalf(t, v.want["message"], res["message"], "Want %v but got %v", v.want, res)
		})
	}
}
