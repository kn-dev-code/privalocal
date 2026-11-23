package token

import (
	"time"
)

type APIToken struct {
	ID        int64     `json:"id" db:"id"`
	UserID    int64     `json:"user_id" db:"user_id"`
	TokenPrefix string     `json:"token_prefix" db:"token_prefix"`
    TokenHash   string     `json:"-"            db:"token_hash"`
    LastUsedAt  *time.Time `json:"last_used_at" db:"last_used_at"`
    RevokedAt   *time.Time `json:"revoked_at"   db:"revoked_at"`
    CreatedAt   time.Time  `json:"created_at"   db:"created_at"`
}
