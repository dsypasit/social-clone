package post

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

type MockHandler struct {
	createPostCalled     bool
	getPostsByUUIDCalled bool
}

func (m *MockHandler) GetPostsByUserUUID(w http.ResponseWriter, r *http.Request) {
	m.getPostsByUUIDCalled = true
}

func (m *MockHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	m.createPostCalled = true
}

func TestRoute(t *testing.T) {
	router := mux.NewRouter()
	mHandler := MockHandler{false, false}
	RegisterPostRouter(router, &mHandler)

	router.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, "/post", nil))
	assert.True(t, mHandler.createPostCalled, "create post not called")

	router.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/post", nil))
	assert.True(t, mHandler.getPostsByUUIDCalled, "get post not called")
}
