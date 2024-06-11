package user

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

type MockHandler struct {
	getUserByUsernameCalled bool
}

func (m *MockHandler) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	m.getUserByUsernameCalled = true
}

func TestRoute(t *testing.T) {
	mhandler := MockHandler{false}
	router := mux.NewRouter()
	RegisterUserRouter(router, &mhandler)

	router.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/user", nil))
	assert.True(t, mhandler.getUserByUsernameCalled, "get user by username not called")
}
