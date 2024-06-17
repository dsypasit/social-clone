package post

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dsypasit/social-clone/server/internal/auth"
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
	jwtSer := auth.NewJwtService("test")
	RegisterPostRouter(router, &mHandler, jwtSer)

	token, _ := jwtSer.GenerateToken("1234")

	req := httptest.NewRequest(http.MethodPost, "/post", nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	router.ServeHTTP(httptest.NewRecorder(), req)
	assert.True(t, mHandler.createPostCalled, "create post not called")

	req = httptest.NewRequest(http.MethodGet, "/post", nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	router.ServeHTTP(httptest.NewRecorder(), req)
	assert.True(t, mHandler.getPostsByUUIDCalled, "get post not called")
}
