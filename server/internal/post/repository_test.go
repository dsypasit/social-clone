package post

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dsypasit/social-clone/server/internal/share/util"
	"github.com/stretchr/testify/assert"
)

func TestCreatePost(t *testing.T) {
	testTable := []struct {
		title   string
		post    PostCreated
		wantId  int64
		wantErr error
	}{
		{
			"create success",
			PostCreated{
				UUID:             "f307d2db-d2ea-4ec9-8d31-27b7443d7c72",
				Content:          "Hello",
				NumLike:          0,
				UserUUID:         "e936e164-52fa-4fd5-b0e0-597c2f270245",
				VisibilityTypeId: 1,
			},
			1, nil,
		},
	}

	for _, v := range testTable {
		t.Run(v.title, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			mock.ExpectQuery("INSERT INTO post").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

			postRepo := NewPostRepository(db)
			id, err := postRepo.CreatePost(v.post)
			assert.Equalf(t, v.wantErr, err, "Want %v but got %v", v.wantErr, err)
			assert.Equalf(t, v.wantId, id, "Want %v but got %v", v.wantId, id)
		})
	}
}

func TestGetPostByUserUUID(t *testing.T) {
	testTable := []struct {
		title    string
		userUUID string
		wantPost []PostResponse
		wantErr  error
	}{
		{
			"create success",
			"f6630558-b800-48ff-9a09-5863d6055154",
			[]PostResponse{
				{
					UUID:             util.Ptr("f307d2db-d2ea-4ec9-8d31-27b7443d7c72"),
					Content:          util.Ptr("Hello"),
					NumLike:          0,
					VisibilityTypeId: 1,
					UserUUID:         util.Ptr("f6630558-b800-48ff-9a09-5863d6055154"),
					Username:         util.Ptr("ronaldo"),
				},
			},
			nil,
		},
	}

	for _, v := range testTable {
		t.Run(v.title, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"uuid", "content", "num_like", "visibility_type_id", "uuid", "username"}).
				AddRow(v.wantPost[0].UUID, v.wantPost[0].Content, v.wantPost[0].NumLike, v.wantPost[0].VisibilityTypeId, v.wantPost[0].UserUUID, v.wantPost[0].Username))

			postRepo := NewPostRepository(db)
			posts, err := postRepo.GetPostsByUserUUID(v.userUUID)
			assert.Equalf(t, v.wantErr, err, "unexpected error: %v", err)
			assert.Equalf(t, v.wantPost, posts, "Want %v but got %v", v.wantPost, posts)
		})
	}
}

func TestGetPostByUserUUID_SQL_ERR(t *testing.T) {
	testTable := []struct {
		title    string
		userUUID string
		wantPost []PostResponse
		wantErr  error
	}{
		{
			"create success",
			"f6630558-b800-48ff-9a09-5863d6055154",
			nil,
			errors.New("some errors"),
		},
	}

	for _, v := range testTable {
		t.Run(v.title, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			mock.ExpectQuery("SELECT").WillReturnError(errors.New("some errors"))

			postRepo := NewPostRepository(db)
			posts, err := postRepo.GetPostsByUserUUID(v.userUUID)
			assert.Equalf(t, v.wantErr, err, "unexpected error: %v", err)
			assert.Equalf(t, v.wantPost, posts, "Want %v but got %v", v.wantPost, posts)
		})
	}
}

func TestGetPosts(t *testing.T) {
	testTable := []struct {
		title    string
		userUUID string
		wantPost []PostResponse
		wantErr  error
	}{
		{
			"create success",
			"f6630558-b800-48ff-9a09-5863d6055154",
			[]PostResponse{
				{
					UUID:             util.Ptr("f307d2db-d2ea-4ec9-8d31-27b7443d7c72"),
					Content:          util.Ptr("Hello"),
					NumLike:          0,
					VisibilityTypeId: 1,
					UserUUID:         util.Ptr("f6630558-b800-48ff-9a09-5863d6055154"),
					Username:         util.Ptr("ronaldo"),
				},
			},
			nil,
		},
	}

	for _, v := range testTable {
		t.Run(v.title, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"uuid", "content", "num_like", "visibility_type_id", "uuid", "username"}).
				AddRow(v.wantPost[0].UUID, v.wantPost[0].Content, v.wantPost[0].NumLike, v.wantPost[0].VisibilityTypeId, v.wantPost[0].UserUUID, v.wantPost[0].Username))

			postRepo := NewPostRepository(db)
			posts, err := postRepo.GetPosts()
			assert.Equalf(t, v.wantErr, err, "unexpected error: %v", err)
			assert.Equalf(t, v.wantPost, posts, "Want %v but got %v", v.wantPost, posts)
		})
	}
}
