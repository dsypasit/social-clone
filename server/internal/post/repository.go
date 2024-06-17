package post

import "database/sql"

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db}
}

func (r *PostRepository) CreatePost(p PostCreated) (int64, error) {
	query := `INSERT INTO post (uuid, content, num_like, visibility_type_id, app_user_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        (SELECT id FROM app_user WHERE uuid = $5)
    ) RETURNING id`

	var id int64
	err := r.db.QueryRow(query, p.UUID, p.Content, 0,
		p.VisibilityTypeId, p.UserUUID).Scan(&id)
	return id, err
}

func (r *PostRepository) GetPostsByUserUUID(userUUID string) ([]PostResponse, error) {
	query := `
  SELECT p.uuid, p.content, p.num_like, p.visibility_type_id, u.uuid, u.username, u.updated_at
  FROM post as p
  LEFT JOIN app_user AS u ON u.id = p.app_user_id
  WHERE u.uuid=$1
  ORDER BY u.updated_at DESC
  `

	var posts []PostResponse
	rows, err := r.db.Query(query, userUUID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var post PostResponse
		err := rows.Scan(&post.UUID, &post.Content, &post.NumLike, &post.VisibilityTypeId, &post.UserUUID, &post.Username, &post.UpdateAt)
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}

	return posts, err
}

func (r *PostRepository) GetPosts() ([]PostResponse, error) {
	query := `
  SELECT p.uuid, p.content, p.num_like, p.visibility_type_id, u.uuid, u.username, u.updated_at
  FROM post as p
  LEFT JOIN app_user AS u ON u.id = p.app_user_id
  ORDER BY u.updated_at DESC
  `

	var posts []PostResponse
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var post PostResponse
		err := rows.Scan(&post.UUID, &post.Content, &post.NumLike, &post.VisibilityTypeId, &post.UserUUID, &post.Username, &post.UpdateAt)
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}

	return posts, err
}
