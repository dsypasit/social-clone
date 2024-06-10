package post

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockRepo struct {
	repoErr error
	postRes []PostResponse
}

func (m *MockRepo) CreatePost(Post) (int64, error) {
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

func TestServiceCreatePost(t *testing.T) {
	testTable := []struct {
		title   string
		input   Post
		wantId  int64
		wantErr error
	}{
		{"should create success", Post{
			Content:          "Hello",
			UserId:           1,
			VisibilityTypeId: 1,
		}, 1, nil},
	}

	for _, v := range testTable {
		t.Run(v.title, func(t *testing.T) {
			m := MockRepo{nil, nil}
			s := NewPostService(&m)
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
					UUID:             "c27e224d-b0af-4a45-8da8-8c5da69c5b03",
					Content:          "Hello1",
					NumLike:          12,
					UserUUID:         "ea151663-aad6-45b2-808b-e3f160956612",
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
			s := NewPostService(&m)
			_, err := s.GetPostsByUserUUID(v.input)
			assert.Equalf(t, v.wantErr, err, "Unexpected error: %v", err)
		})
	}
}
