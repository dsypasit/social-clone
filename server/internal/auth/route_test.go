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
}

func (m *MockHandler) Login(http.ResponseWriter, *http.Request) {
	m.loginCalled = true
}

func (m *MockHandler) Signup(http.ResponseWriter, *http.Request) {
	m.signupCalled = true
}

func TestRoute(t *testing.T) {
	router := mux.NewRouter()
	authHandler := MockHandler{false, false}
	RegisterAuthRouter(router, &authHandler)

	// Test signup route
	router.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("POST", "/signup", nil))
	assert.True(t, authHandler.signupCalled, "signup handler not called")

	// Test login route (similar approach)
	router.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest("POST", "/login", nil))
	assert.True(t, authHandler.loginCalled, "login handler not called")
}
