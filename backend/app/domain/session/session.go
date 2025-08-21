package session

import (
	"time"

	"github.com/google/uuid"
	"inoUwU/pinu/app/domain/user"
)

// Session セッションエンティティ
type Session struct {
	SessionID uuid.UUID   `json:"session_id" bun:",pk"`
	UserID    user.UserID `json:"user_id"`
	CreatedAt time.Time   `json:"created_at"`
	ExpiresAt time.Time   `json:"expires_at"`
	IPAddress *string     `json:"ip_address,omitempty"`
	UserAgent *string     `json:"user_agent,omitempty"`
}

// TableSession テーブルセッションエンティティ
type TableSession struct {
	TableSessionID uuid.UUID `json:"table_session_id" bun:",pk"`
	TableID        string    `json:"table_id"`
	OrdersID       uuid.UUID `json:"orders_id"`
	CreatedAt      time.Time `json:"created_at"`
	LastUsed       time.Time `json:"last_used"`
	ExpiresAt      time.Time `json:"expires_at"`
}
