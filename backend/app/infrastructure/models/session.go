package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// Session represents a user session.
type Session struct {
	bun.BaseModel `bun:"table:sessions,alias:ss"`

	SESSION_ID    string    `bun:"session_id,pk"`
	USER_ID       string    `bun:"user_id,"`
	REFRESH_TOKEN string    `bun:"refresh_token,notnull"`
	IS_REVOKED    bool      `bun:"is_revoked,notnull,default:false"`
	CREATED_AT    time.Time `bun:"created_at,notnull"`
	EXPIRES_AT    time.Time `bun:"expires_at,notnull"`
}

// TableSession represents a table session.
type TableSession struct {
	bun.BaseModel `bun:"table:table_sessions,alias:ts"`

	TableSessionID uuid.UUID `bun:"table_session_id,pk"`
	TableID        string    `bun:"table_id"`
	IsRevoked      bool      `bun:"is_revoked"`
	CreatedAt      time.Time `bun:"created_at"`
	LastUsed       time.Time `bun:"last_used"`
	ExpiresAt      time.Time `bun:"expires_at"`
}
