package store

import (
	"context"
	"database/sql"
)

type User struct {
	Id        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"` // Exclude password from JSON responses
	CreatedAt string `json:"created_at"`
}

type UsersStore struct {
	db *sql.DB
}

// Create inserts a new user into the database and returns the generated ID and creation timestamp.
func (s *UsersStore) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users (username, email, password)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`

	err := s.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password).Scan(
		&user.Id,
		&user.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
