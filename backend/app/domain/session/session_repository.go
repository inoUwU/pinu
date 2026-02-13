package session

import (
	"context"

	"github.com/google/uuid"
)

// SessionStore セッション永続化のポート
type SessionStore interface {
	// Session関連
	CreateSession(ctx context.Context, session *Session) error
	GetSessionByID(ctx context.Context, id string) (*Session, error)
	RevokeSessionByID(ctx context.Context, id string) error
	DeleteSession(ctx context.Context, id string) error

	// TableSession関連
	CreateTableSession(ctx context.Context, tableSession *TableSession) error
	GetTableSessionByID(ctx context.Context, id uuid.UUID) (*TableSession, error)
	GetTableSessionByTableID(ctx context.Context, tableID string) (*TableSession, error)
	UpdateTableSessionLastUsed(ctx context.Context, id uuid.UUID) error
	DeleteTableSession(ctx context.Context, id uuid.UUID) error
	DeleteExpiredTableSessions(ctx context.Context) error
}
