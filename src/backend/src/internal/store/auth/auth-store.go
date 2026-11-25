package auth

import (
	"context"
	"errors"
	"fmt"
	"time"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type User struct {
	ID        int64     `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	PasswordHash  string    `json:"-" db:"password"`
	Role      Role      `json:"role" db:"role"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Store struct {
	pool *pgxpool.Pool
}


func (s *Store) GetUserById(ctx context.Context, id int64) (User, error) {
	const query = `
		SELECT ID, username, email, password_hash, role, created_at
		WHERE ID = $1
	`

	var u User
	err := s.pool.QueryRow(ctx, query, id).Scan(
		&u.ID,
		&u.Username,
		&u.Email,
		&u.PasswordHash,
		&u.Role,
		&u.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return User{}, fmt.Errorf("Get user %d: %w", id, pgx.ErrNoRows)
		}
		return User{}, fmt.Errorf("query user %d: %w", id, err)
	}
	return u, nil
}







