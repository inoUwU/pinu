package models

import (
	"time"

	"github.com/uptrace/bun"
)

// UserModel ユーザーのインフラストラクチャーモデル
type UserModel struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	UserID       string    `bun:"user_id,pk"`
	LoginID      string    `bun:"login_id"`
	Name         string    `bun:"name"`
	PasswordHash string    `bun:"password_hash"`
	PasswordSalt string    `bun:"password_salt"`
	IsAdmin      bool      `bun:"is_admin"`
	CreatedAt    time.Time `bun:"created_at"`
}
