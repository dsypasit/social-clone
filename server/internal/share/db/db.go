package db

import (
	"context"
	"database/sql"
	"log"
	"time"
)

var DB *sql.DB

func Init(DBConnection string) error {
	var err error
	DB, err = sql.Open("postgres", DBConnection)

	return err
}

func Close() {
	DB.Close()
}

func Ping(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if err := DB.PingContext(ctx); err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
}
