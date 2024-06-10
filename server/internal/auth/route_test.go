package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

type MockHandler struct {
	signupCalled bool
	loginCalled  bool
	tokenCalled  bool
}

func (m *MockHandler) Login(http.ResponseWriter, *http.Request) {
	m.loginCalled = true
}

func (m *MockHandler) Signup(http.ResponseWriter, *http.Request) {
	m.signupCalled = true
}

func (m *MockHandler) CheckToken(http.ResponseWriter, *http.Request) {
	m.tokenCalled = true
}

func TestRoute(t *testing.T) {
	router := mux.NewRouter()
	authHandler := MockHandler{false, false, false}
	RegisterAuthRouter(router, &authHandler)

	// Test signup route
	router.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("POST", "/auth/signup", nil))
	assert.True(t, authHandler.signupCalled, "signup handler not called")

	// Test login route
	router.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("POST", "/auth/login", nil))
	assert.True(t, authHandler.loginCalled, "login handler not called")

	// Test check token route
	router.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("POST", "/auth/checktoken", nil))
	assert.True(t, authHandler.tokenCalled, "login handler not called")
}
