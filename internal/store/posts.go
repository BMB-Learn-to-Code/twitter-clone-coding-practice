package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type PostsStore struct {
	db *sql.DB
}

type Post struct {
	ID        int64    `json:"id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Tags      []string `json:"tags"`
	UserID    int64    `json:"user_id"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

func (s *PostsStore) Create(ctx context.Context, post *Post) error {
	query := `
		INSERT INTO posts (title, content, tags, user_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`
	s.db.QueryContext(ctx, query, post.Title, post.Content, pq.Array(post.Tags), post.UserID)

	return nil
}
