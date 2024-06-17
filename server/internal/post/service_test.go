package post

import (
	"database/sql"
	"testing"

	"github.com/dsypasit/social-clone/server/internal/share/util"
	"github.com/dsypasit/social-clone/server/internal/user"
	"github.com/stretchr/testify/assert"
)

type MockUserSrv struct{}

func (m *MockUserSrv) GetUserByUUID(s string) (user.User, error) {
	return user.User{}, nil
}

type MockRepo struct {
	repoErr error
	postRes []PostResponse
}

func (m *MockRepo) CreatePost(PostCreated) (int64, error) {
	if m.repoErr != nil {
		return 0, m.repoErr
	}
	return 1, nil
}

func (m *MockRepo) GetPostsByUserUUID(string) ([]PostResponse, error) {
	if m.repoErr != nil {
		return []PostResponse{}, m.repoErr
	}
	return m.postRes, nil
}

func (m *MockRepo) GetPosts() ([]PostResponse, error) {
	if m.repoErr != nil {
		return []PostResponse{}, m.repoErr
	}
	return m.postRes, nil
}

func TestServiceCreatePost(t *testing.T) {
	testTable := []struct {
		title   string
		input   PostCreated
		wantId  int64
		wantErr error
	}{
		{"should create success", PostCreated{
			Content:          "Hello",
			UserUUID:         "cfaecdf4-2a2a-47fb-a1fa-114b18383feb",
			VisibilityTypeId: 1,
		}, 1, nil},
	}

	for _, v := range testTable {
		t.Run(v.title, func(t *testing.T) {
			m := MockRepo{nil, nil}
			mu := MockUserSrv{}

			s := NewPostService(&m, &mu)
			id, err := s.CreatePost(v.input)
			assert.Equalf(t, v.wantErr, err, "Unexpected error: %v", err)
			assert.Equalf(t, v.wantId, id, "Want %v but got %v", v.wantId, id)
		})
	}
}

func TestServiceGetPostByUserUUID(t *testing.T) {
	testTable := []struct {
		title    string
		mErr     error
		input    string
		wantPost []PostResponse
		wantErr  error
	}{
		{
			"Should return []post response", nil, "ea151663-aad6-45b2-808b-e3f160956612",
			[]PostResponse{
				{
					UUID:             util.Ptr("c27e224d-b0af-4a45-8da8-8c5da69c5b03"),
					Content:          util.Ptr("Hello1"),
					NumLike:          12,
					UserUUID:         util.Ptr("ea151663-aad6-45b2-808b-e3f160956612"),
					VisibilityTypeId: 1,
				},
			},
			nil,
		}, {
			"should return no row", sql.ErrNoRows, "ea151663-aad6-45b2-808b-e3f160956612",
			[]PostResponse{},
			ErrNoRows,
		},
	}

	for _, v := range testTable {
		t.Run(v.title, func(t *testing.T) {
			m := MockRepo{v.mErr, v.wantPost}
			s := NewPostService(&m, &MockUserSrv{})
			_, err := s.GetPostsByUserUUID(v.input)
			assert.Equalf(t, v.wantErr, err, "Unexpected error: %v", err)
		})
	}
}

func TestServiceGetPost(t *testing.T) {
	testTable := []struct {
		title    string
		mErr     error
		input    string
		wantPost []PostResponse
		wantErr  error
	}{
		{
			"Should return []post response", nil, "ea151663-aad6-45b2-808b-e3f160956612",
			[]PostResponse{
				{
					UUID:             util.Ptr("c27e224d-b0af-4a45-8da8-8c5da69c5b03"),
					Content:          util.Ptr("Hello1"),
					NumLike:          12,
					UserUUID:         util.Ptr("ea151663-aad6-45b2-808b-e3f160956612"),
					VisibilityTypeId: 1,
				},
			},
			nil,
		}, {
			"should return no row", sql.ErrNoRows, "ea151663-aad6-45b2-808b-e3f160956612",
			[]PostResponse{},
			ErrNoRows,
		},
	}

	for _, v := range testTable {
		t.Run(v.title, func(t *testing.T) {
			m := MockRepo{v.mErr, v.wantPost}
			s := NewPostService(&m, &MockUserSrv{})
			_, err := s.GetPosts()
			assert.Equalf(t, v.wantErr, err, "Unexpected error: %v", err)
		})
	}
}
