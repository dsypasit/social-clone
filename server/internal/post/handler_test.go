package post

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dsypasit/social-clone/server/internal/share/util"
	"github.com/stretchr/testify/assert"
)

type mService struct {
	postsResp []PostResponse
	isErr     error
}

func (m *mService) CreatePost(PostCreated) (int64, error) {
	return 1, m.isErr
}

func (m *mService) GetPostsByUserUUID(string) ([]PostResponse, error) {
	return m.postsResp, m.isErr
}

func TestHandlerCreatePost(t *testing.T) {
	post, _ := json.Marshal(PostCreated{
		Content: "hello", UserUUID: "1eb64cd3-03ef-4ac7-9008-e0ab63f4105f",
		VisibilityTypeId: 1,
	})
	testTable := []struct {
		title      string
		post       []byte
		serviceErr error
		wantStatus int
		wantBody   map[string]string
	}{
		{
			"should create post success", post, nil, http.StatusCreated,
			util.BuildResponse("created post successful!"),
		},
	}

	for _, v := range testTable {
		t.Run(v.title, func(t *testing.T) {
			mService := mService{}
			h := NewPostHandler(&mService)

			req, _ := http.NewRequest(http.MethodGet, "/", bytes.NewReader(v.post))
			ctx := context.WithValue(req.Context(), "userUUID", "d8ff2276-9c8c-47f3-9b98-9ac1c8413c18")
			req = req.WithContext(ctx)
			rec := httptest.NewRecorder()

			h.CreatePost(rec, req)

			var res map[string]string
			json.NewDecoder(rec.Body).Decode(&res)
			assert.Equalf(t, v.wantStatus, rec.Code, "Want %v but got %v", v.wantStatus, rec.Code)
			assert.Equalf(t, v.wantBody, res, "Want %v but got %v", v.wantBody, res)
		})
	}
}

func TestHandlerCreatePost_Invalid(t *testing.T) {
	post, _ := json.Marshal(PostCreated{
		Content: "hello", UserUUID: "25864f86-da09-4009-8585-56d143e99da6",
		VisibilityTypeId: 1,
	})
	testTable := []struct {
		title      string
		post       []byte
		serviceErr error
		wantStatus int
		wantBody   map[string]string
	}{
		{
			"should create post success", []byte(""), nil, http.StatusBadRequest,
			util.BuildErrResponse("invalid request")(nil),
		},
		{
			"should create post success", []byte("{\"user_id\":\"3\"}"), nil, http.StatusBadRequest,
			util.BuildErrResponse("invalid request")(nil),
		},
		{
			"should create post success", []byte("{\"user_id\":\"3\", \"content\": \"hello\"}"), nil, http.StatusBadRequest,
			util.BuildErrResponse("invalid request")(nil),
		},
		{
			"should create post success", []byte("{\"user_id\":\"3\", \"Visibility_type_id\": \"hello\"}"), nil, http.StatusBadRequest,
			util.BuildErrResponse("invalid request")(nil),
		},
		{
			"should create post success", post, errors.New("service err"), http.StatusInternalServerError,
			util.BuildErrResponse("service error")(nil),
		},
	}

	for _, v := range testTable {
		t.Run(v.title, func(t *testing.T) {
			mService := mService{nil, v.serviceErr}
			h := NewPostHandler(&mService)

			req, _ := http.NewRequest(http.MethodGet, "/", bytes.NewReader(v.post))
			ctx := context.WithValue(req.Context(), "userUUID", "1e41893e-d6a0-44e5-9dd1-9bb1949bdb0d")
			req = req.WithContext(ctx)
			rec := httptest.NewRecorder()

			h.CreatePost(rec, req)

			var res map[string]string
			json.NewDecoder(rec.Body).Decode(&res)
			assert.Equalf(t, v.wantStatus, rec.Code, "Want %v but got %v", v.wantStatus, rec.Code)
			assert.Equalf(t, v.wantBody["message"], res["message"], "Want %v but got %v", v.wantBody, res)
		})
	}
}

func TestHandlerGetPostByUserUUID(t *testing.T) {
	testTable := []struct {
		title      string
		useruuid   string
		want       []PostResponse
		wantStatus int
	}{
		{
			"should return post", "4a1ec88b-380e-4dc4-bba8-a88e85dc6663",
			[]PostResponse{
				{
					UUID:             util.Ptr("0ee1abd0-a330-488d-b170-b33f58dd6178"),
					Content:          util.Ptr("hello"),
					NumLike:          10,
					UserUUID:         util.Ptr("4a1ec88b-380e-util.4dc4-bba8-a88e85dc6663"),
					VisibilityTypeId: 1,
				},
			},
			http.StatusOK,
		},
	}

	for _, v := range testTable {
		t.Run(v.title, func(t *testing.T) {
			ms := mService{v.want, nil}
			h := NewPostHandler(&ms)

			url := fmt.Sprintf("?useruuid=%s", v.useruuid)
			req, _ := http.NewRequest(http.MethodGet, url, nil)
			ctx := context.WithValue(req.Context(), "userUUID", "7a053eee-a70d-442c-81ba-c36d72d3f87b")
			req = req.WithContext(ctx)

			rec := httptest.NewRecorder()

			h.GetPostsByUserUUID(rec, req)

			var res []PostResponse
			json.NewDecoder(rec.Body).Decode(&res)
			assert.Equalf(t, v.wantStatus, rec.Code, "want %v but got %v", v.wantStatus, rec.Code)
			assert.Equalf(t, v.want, res, "want %v but got %v", v.want, res)
		})
	}
}

func TestHandlerGetPostByUserUUID_Invalid(t *testing.T) {
	testTable := []struct {
		title      string
		useruuid   string
		serviceErr error
		want       map[string]string
		wantStatus int
	}{
		{
			"should return bad request cause missing uuid", "", nil,
			util.BuildErrResponse("invalid request")(nil),
			http.StatusBadRequest,
		},
		{
			"should return bad request cause service error", "9329518b-4d55-4665-b640-b95261d2b204", errors.New("service error"),
			util.BuildErrResponse("service failed")(nil),
			http.StatusInternalServerError,
		},
	}

	for _, v := range testTable {
		t.Run(v.title, func(t *testing.T) {
			ms := mService{nil, v.serviceErr}
			h := NewPostHandler(&ms)

			url := fmt.Sprintf("?useruuid=%s", v.useruuid)
			req, _ := http.NewRequest(http.MethodGet, url, nil)
			ctx := context.WithValue(req.Context(), "userUUID", "2d36ecaa-e321-4257-a8f0-3064446a4378")
			req = req.WithContext(ctx)

			rec := httptest.NewRecorder()

			h.GetPostsByUserUUID(rec, req)

			var res map[string]string
			json.NewDecoder(rec.Body).Decode(&res)
			assert.Equalf(t, v.wantStatus, rec.Code, "want %v but got %v", v.wantStatus, rec.Code)
			assert.Equalf(t, v.want["message"], res["message"], "want %v but got %v", v.want, res)
		})
	}
}
