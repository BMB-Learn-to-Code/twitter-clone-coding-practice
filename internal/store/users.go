package store

import (
	"context"
	"database/sql"
)

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UsersStore struct {
	db *sql.DB
}

func (s *UsersStore) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users (username, email, password)
		VALUES ($1, $2, $3)
		RETURNING id, username, email, password
	`

	err := s.db.QueryRowContext(ctx, query).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		return err
	}

	return nil
}
