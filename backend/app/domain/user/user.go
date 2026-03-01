package user

import "time"

// UserID ユーザーIDの値オブジェクト
type UserID string

// LoginID ログインIDの値オブジェクト
type LoginID string

// User ユーザーエンティティ
type User struct {
	UserID       UserID    `json:"user_id"`
	LoginID      LoginID   `json:"login_id"`
	PasswordHash string    `json:"password_hash"`
	PasswordSalt string    `json:"password_salt"`
	Name         string    `json:"name"`
	IsAdmin      bool      `json:"is_admin"`
	CreatedAt    time.Time `json:"created_at"`
}
