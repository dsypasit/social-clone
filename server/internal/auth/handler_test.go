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
	isErr bool
}

func (m *MockAuthService) Signup(u user.UserCreated) (string, error) {
	if m.isErr {
		return "", errors.New("error")
	}
	return "token", nil
}

func (m *MockAuthService) Login(u User) (string, error) {
	if m.isErr {
		return "", errors.New("error")
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
		isErr      bool
		wantStatus int
		wantBody   map[string]string
	}{
		{"should internal error cause service not working", bytes.NewReader([]byte("{\"username\":\"abc\", \"password\":\"abc\", \"email\": \"a@gmail.com\"}")), true, http.StatusInternalServerError, map[string]string{"message": "error"}},
		{"should get token", bytes.NewReader([]byte("{\"username\":\"abc\", \"password\":\"abcd\", \"email\": \"a@gmail.com\"}")), false, http.StatusCreated, map[string]string{"token": "token"}},
	}

	for _, v := range testTable {
		t.Run(v.title, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, "/api/v1/signup", v.input)
			assert.Nil(t, err, "Expected nil from create req")

			rec := httptest.NewRecorder()

			mockService := MockAuthService{User{}, v.isErr}
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
		isErr       bool
		initialUser User
		wantStatus  int
		wantBody    map[string]string
	}{
		{"should bad request cause password invalid", bytes.NewReader([]byte("{\"username\":\"abc\", \"password\":\"abc\"}")), false, User{Password: "abcd"}, http.StatusBadRequest, map[string]string{"message": "invalid password"}},
		{"should internal error cause service not working", bytes.NewReader([]byte("{\"username\":\"abc\", \"password\":\"abc\", \"email\":\"a@gmail.com\"}")), true, User{Password: "abc"}, http.StatusInternalServerError, map[string]string{"error": "error"}},
		{"should get token", bytes.NewReader([]byte("{\"username\":\"abc\", \"password\":\"abc\", \"email\":\"a@gmail.com\"}")), false, User{Password: "abc"}, http.StatusOK, map[string]string{"token": "token"}},
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
