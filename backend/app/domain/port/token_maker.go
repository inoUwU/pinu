package port

import (
	"inoUwU/pinu/app/domain/user"
	"time"
)

// Claims トークンのクレーム情報
type Claims struct {
	ID        string
	UserID    string
	LoginID   string
	IsAdmin   bool
	IssuedAt  time.Time
	ExpiresAt time.Time
}

// TokenMaker トークン生成・検証のポート
type TokenMaker interface {
	// GenerateToken ユーザー情報からトークンを生成する
	GenerateToken(user *user.User, duration time.Duration) (string, *Claims, error)

	// VerifyToken トークンを検証しクレーム情報を返す
	VerifyToken(tokenString string) (*Claims, error)
}
