package session

import (
	"time"

	"inoUwU/pinu/app/domain/user"

	"github.com/google/uuid"
)

// Session セッションエンティティ
type Session struct {
	SessionID    string
	UserID       user.UserID
	RefreshToken string
	IsRevoked    bool
	CreatedAt    time.Time
	ExpiresAt    time.Time
}

// TableSession テーブルセッションエンティティ
type TableSession struct {
	TableSessionID uuid.UUID
	TableID        string
	IsRevoked      bool
	CreatedAt      time.Time
	LastUsed       time.Time
	ExpiresAt      time.Time
}
