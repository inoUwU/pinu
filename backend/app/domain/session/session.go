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

func (s *Session) IsExpired(now time.Time) bool {
	if s == nil {
		return true
	}
	return !now.Before(s.ExpiresAt)
}

func (s *Session) IsValid(now time.Time) bool {
	if s == nil {
		return false
	}
	return !s.IsRevoked && !s.IsExpired(now)
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

func (s *TableSession) IsExpired(now time.Time) bool {
	if s == nil {
		return true
	}
	return !now.Before(s.ExpiresAt)
}

func (s *TableSession) IsActive(now time.Time) bool {
	if s == nil {
		return false
	}
	return !s.IsRevoked && !s.IsExpired(now)
}
