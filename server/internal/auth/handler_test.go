package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dsypasit/social-clone/server/internal/user"
	"github.com/stretchr/testify/assert"
)

type MockAuthService struct {
	user  User
	isErr error
}

func (m *MockAuthService) Signup(u user.UserCreated) (string, error) {
	if m.isErr != nil {
		return "", m.isErr
	}
	return "token", nil
}

func (m *MockAuthService) CheckToken(token string) bool {
	return m.isErr == nil
}

func (m *MockAuthService) Login(u User) (string, error) {
	if m.isErr != nil {
		return "", m.isErr
	}
	if m.user.Password != u.Password {
		return "", ErrInvalidPassword
	}
	return "token", nil
}

func TestSignup_InvalidFormat(t *testing.T) {
	testTable := []struct {
		title      string
		input      io.Reader
		wantStatus int
		wantBody   map[string]string
	}{
		{"should bad request cause input is string", bytes.NewReader([]byte("string")), http.StatusBadRequest, map[string]string{"message": "invalid structure format"}},
		{"should bad request cause input mismatch struct", bytes.NewReader([]byte("{\"invalid\":\"req\"}")), http.StatusBadRequest, map[string]string{"message": "invalid structure format"}},
		{"should bad request cause user empty", bytes.NewReader([]byte("{\"username\":\"\", \"password\":\"qwer\"}")), http.StatusBadRequest, map[string]string{"message": "username or password empty or invalid email format"}},
		{"should bad request cause password empty", bytes.NewReader([]byte("{\"username\":\"abc\", \"password\":\"\"}")), http.StatusBadRequest, map[string]string{"message": "username or password empty or invalid email format"}},
	}

	for _, v := range testTable {
		t.Run(v.title, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, "/api/v1/signup", v.input)
			assert.Nil(t, err, "Expected nil from create req")

			rec := httptest.NewRecorder()

			mockService := MockAuthService{}
			authHandler := NewAuthHandler(&mockService)
			authHandler.Signup(rec, req)

			expected := map[string]string{"message": "invalid structure format"}
			actualCode := rec.Code
			var actualResponse map[string]string
			err = json.NewDecoder(rec.Body).Decode(&actualResponse)
			assert.Nil(t, err, "Expected nil from response decoding")

			assert.Equalf(t, v.wantStatus, actualCode, "Expected bad request status code, but got %v", actualCode)
			assert.Equalf(t, v.wantBody, actualResponse, "Expected %v, but got %v", expected, actualCode)
		})
	}
}

func TestSignup_ServiceNotWorking(t *testing.T) {
	testTable := []struct {
		title      string
		input      io.Reader
		serviceErr error
		wantStatus int
		wantBody   map[string]string
	}{
		{"should internal error cause service not working", bytes.NewReader([]byte("{\"username\":\"abc\", \"password\":\"abc\", \"email\": \"a@gmail.com\"}")), errors.New("error"), http.StatusInternalServerError, map[string]string{"message": "error"}},
		{"should bad request cause duplicate user", bytes.NewReader([]byte("{\"username\":\"abc\", \"password\":\"asdf\", \"email\":\"a@gmail.com\"}")), user.ErrDupUsername, http.StatusBadRequest, map[string]string{"message": "duplicate username"}},
		{"should get token", bytes.NewReader([]byte("{\"username\":\"abc\", \"password\":\"abcd\", \"email\": \"a@gmail.com\"}")), nil, http.StatusCreated, map[string]string{"token": "token"}},
	}

	for _, v := range testTable {
		t.Run(v.title, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, "/api/v1/signup", v.input)
			assert.Nil(t, err, "Expected nil from create req")

			rec := httptest.NewRecorder()

			mockService := MockAuthService{User{}, v.serviceErr}
			authHandler := NewAuthHandler(&mockService)
			authHandler.Signup(rec, req)

			expected := map[string]string{"message": "invalid structure format"}
			actualCode := rec.Code
			var actualResponse map[string]string
			err = json.NewDecoder(rec.Body).Decode(&actualResponse)
			assert.Nil(t, err, "Expected nil from response decoding")

			assert.Equalf(t, v.wantStatus, actualCode, "Expected bad request status code, but got %v", actualCode)
			assert.Equalf(t, v.wantBody, actualResponse, "Expected %v, but got %v", expected, actualCode)
		})
	}
}

func TestLogin_InvalidFormat(t *testing.T) {
	testTable := []struct {
		title      string
		input      io.Reader
		wantStatus int
		wantBody   map[string]string
	}{
		{"should bad request cause input is string", bytes.NewReader([]byte("string")), http.StatusBadRequest, map[string]string{"message": "invalid structure format"}},
		{"should bad request cause input mismatch struct", bytes.NewReader([]byte("{\"invalid\":\"req\"}")), http.StatusBadRequest, map[string]string{"message": "invalid structure format"}},
		{"should bad request cause user empty", bytes.NewReader([]byte("{\"username\":\"\", \"password\":\"qwer\"}")), http.StatusBadRequest, map[string]string{"message": "username or password empty or invalid email format"}},
		{"should bad request cause password empty", bytes.NewReader([]byte("{\"username\":\"abc\", \"password\":\"\"}")), http.StatusBadRequest, map[string]string{"message": "username or password empty or invalid email format"}},
	}

	for _, v := range testTable {
		t.Run(v.title, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, "/api/v1/signup", v.input)
			assert.Nil(t, err, "Expected nil from create req")

			rec := httptest.NewRecorder()

			mockService := MockAuthService{}
			authHandler := NewAuthHandler(&mockService)
			authHandler.Login(rec, req)

			expected := map[string]string{"message": "invalid structure format"}
			actualCode := rec.Code
			var actualResponse map[string]string
			err = json.NewDecoder(rec.Body).Decode(&actualResponse)
			assert.Nil(t, err, "Expected nil from response decoding")

			assert.Equalf(t, v.wantStatus, actualCode, "Expected bad request status code, but got %v", actualCode)
			assert.Equalf(t, v.wantBody, actualResponse, "Expected %v, but got %v", expected, actualCode)
		})
	}
}

func TestLogin_Service(t *testing.T) {
	testTable := []struct {
		title       string
		input       io.Reader
		isErr       error
		initialUser User
		wantStatus  int
		wantBody    map[string]string
	}{
		{"should bad request cause password invalid", bytes.NewReader([]byte("{\"username\":\"abc\", \"password\":\"abc\"}")), nil, User{Password: "abcd"}, http.StatusBadRequest, map[string]string{"message": "invalid password"}},
		{"should internal error cause service not working", bytes.NewReader([]byte("{\"username\":\"abc\", \"password\":\"abc\", \"email\":\"a@gmail.com\"}")), errors.New("error"), User{Password: "abc"}, http.StatusInternalServerError, map[string]string{"error": "error"}},
		{"should get token", bytes.NewReader([]byte("{\"username\":\"abc\", \"password\":\"abc\", \"email\":\"a@gmail.com\"}")), nil, User{Password: "abc"}, http.StatusOK, map[string]string{"token": "token"}},
	}

	for _, v := range testTable {
		t.Run(v.title, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, "/api/v1/signup", v.input)
			assert.Nil(t, err, "Expected nil from create req")

			rec := httptest.NewRecorder()

			mockService := MockAuthService{v.initialUser, v.isErr}
			authHandler := NewAuthHandler(&mockService)
			authHandler.Login(rec, req)

			expected := map[string]string{"message": "invalid structure format"}
			actualCode := rec.Code
			var actualResponse map[string]string
			err = json.NewDecoder(rec.Body).Decode(&actualResponse)
			assert.Nil(t, err, "Expected nil from response decoding")

			assert.Equalf(t, v.wantStatus, actualCode, "Expected bad request status code, but got %v", actualCode)
			assert.Equalf(t, v.wantBody, actualResponse, "Expected %v, but got %v", expected, actualCode)
		})
	}
}

func TestHandlerCheckToken(t *testing.T) {
	t.Run("Should return no content status", func(t *testing.T) {
		mService := MockAuthService{isErr: nil}
		aHandler := NewAuthHandler(&mService)

		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		req.Header.Add("Authorization", "Bearer validtoken")

		rec := httptest.NewRecorder()
		aHandler.CheckToken(rec, req)

		wantStatus := http.StatusNoContent
		assert.Equalf(t, wantStatus, rec.Code, "Want %v but got %v", wantStatus, rec.Code)
	})

	t.Run("should return unautorized", func(t *testing.T) {
		mService := MockAuthService{isErr: errors.New("")}
		aHandler := NewAuthHandler(&mService)

		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		req.Header.Add("Authorization", "Bearer invalid")

		rec := httptest.NewRecorder()
		aHandler.CheckToken(rec, req)

		wantStatus := http.StatusUnauthorized
		assert.Equalf(t, wantStatus, rec.Code, "Want %v but got %v", wantStatus, rec.Code)
	})

	t.Run("should return unautorized cause invalid header", func(t *testing.T) {
		mService := MockAuthService{isErr: nil}
		aHandler := NewAuthHandler(&mService)

		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		req.Header.Add("Authorization", "")

		rec := httptest.NewRecorder()
		aHandler.CheckToken(rec, req)

		wantStatus := http.StatusUnauthorized
		assert.Equalf(t, wantStatus, rec.Code, "Want %v but got %v", wantStatus, rec.Code)
	})
}
