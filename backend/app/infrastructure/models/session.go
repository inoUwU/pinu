package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// SessionModel ユーザーセッションのインフラストラクチャーモデル
type SessionModel struct {
	bun.BaseModel `bun:"table:sessions,alias:ss"`

	SessionID    string    `bun:"session_id,pk"`
	UserID       string    `bun:"user_id,notnull"`
	RefreshToken string    `bun:"refresh_token,notnull"`
	IsRevoked    bool      `bun:"is_revoked,notnull,default:false"`
	CreatedAt    time.Time `bun:"created_at,notnull"`
	ExpiresAt    time.Time `bun:"expires_at,notnull"`
}

// TableSessionModel テーブルセッションのインフラストラクチャーモデル
type TableSessionModel struct {
	bun.BaseModel `bun:"table:table_sessions,alias:ts"`

	TableSessionID uuid.UUID `bun:"table_session_id,pk"`
	TableID        string    `bun:"table_id"`
	IsRevoked      bool      `bun:"is_revoked"`
	CreatedAt      time.Time `bun:"created_at"`
	LastUsed       time.Time `bun:"last_used"`
	ExpiresAt      time.Time `bun:"expires_at"`
}
