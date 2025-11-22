package store

import (
	"context"
	"database/sql"
)

type PostsStore struct {
	db *sql.DB
}

type Post struct {
	ID      int64
	Title   string
	Content string
	Author  string
}

func (s *PostsStore) Create(ctx context.Context) error {
	return nil
}
