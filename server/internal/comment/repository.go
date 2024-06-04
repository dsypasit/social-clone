package comment

import "database/sql"

type commentRepository struct {
	db *sql.DB
}

func NewcommentRepository(db *sql.DB) *commentRepository {
	return &commentRepository{db}
}
