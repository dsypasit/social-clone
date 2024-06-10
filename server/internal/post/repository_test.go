package post

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreatePost(t *testing.T) {
	testTable := []struct {
		title   string
		post    Post
		wantId  int64
		wantErr error
	}{
		{
			"create success",
			Post{
				UUID:             "f307d2db-d2ea-4ec9-8d31-27b7443d7c72",
				Content:          "Hello",
				NumLike:          0,
				UserId:           1,
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
					UUID:             "f307d2db-d2ea-4ec9-8d31-27b7443d7c72",
					Content:          "Hello",
					NumLike:          0,
					VisibilityTypeId: 1,
					UserUUID:         "f6630558-b800-48ff-9a09-5863d6055154",
				},
			},
			nil,
		},
	}

	for _, v := range testTable {
		t.Run(v.title, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"uuid", "content", "num_like", "visibility_type_id", "uuid"}).
				AddRow(v.wantPost[0].UUID, v.wantPost[0].Content, v.wantPost[0].NumLike, v.wantPost[0].VisibilityTypeId, v.wantPost[0].UserUUID))

			postRepo := NewPostRepository(db)
			posts, err := postRepo.GetPostByUserUUID(v.userUUID)
			assert.Equalf(t, v.wantErr, err, "unexpected error: %v", err)
			assert.Equalf(t, v.wantPost, posts, "Want %v but got %v", v.wantPost, posts)
		})
	}
}
