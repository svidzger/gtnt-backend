package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int            `json:"id"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	About     sql.NullString `json:"about"`
	Password  string         `json:"password"`
	AvatarURL sql.NullString `json:"avatar_url"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	Role      string         `json:"role"`
	IsActive  bool           `json:"is_active"`
	LastLogin sql.NullTime   `json:"last_login"`
}
