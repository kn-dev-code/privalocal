package store

import "time"

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type User struct {
	ID        int64     `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password"`
	Role      Role      `json:"role" db:"role"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type APIToken struct {
	ID        int64     `json:"id" db:"id"`
	UserID    int64     `json:"user_id" db:"user_id"`
	Token     string    `json:"api_token,omitempty" db:"api_token"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type ServerState string

const (
	ServerStateIdle      ServerState = "IDLE"
	ServerStateConnected ServerState = "CONNECTED"
	ServerStateError     ServerState = "ERROR"
	ServerStateRetrying  ServerState = "RETRYING"
)