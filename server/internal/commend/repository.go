package commend

import "database/sql"

type CommendRepository struct {
	db *sql.DB
}

func NewCommendRepository(db *sql.DB) *CommendRepository {
	return &CommendRepository{db}
}
