package comment

import "database/sql"

type CommentRepository struct {
	db *sql.DB
}

func NewcommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{db}
}
