package session

import (
	"context"

	"github.com/google/uuid"
	"inoUwU/pinu/app/domain/user"
)

// SessionRepository セッションリポジトリのインターface
type SessionRepository interface {
	// Session関連
	CreateSession(ctx context.Context, session *Session) error
	GetSessionByID(ctx context.Context, id uuid.UUID) (*Session, error)
	GetSessionsByUserID(ctx context.Context, userID user.UserID) ([]*Session, error)
	DeleteSession(ctx context.Context, id uuid.UUID) error
	DeleteExpiredSessions(ctx context.Context) error

	// TableSession関連
	CreateTableSession(ctx context.Context, tableSession *TableSession) error
	GetTableSessionByID(ctx context.Context, id uuid.UUID) (*TableSession, error)
	GetTableSessionByTableID(ctx context.Context, tableID string) (*TableSession, error)
	UpdateTableSessionLastUsed(ctx context.Context, id uuid.UUID) error
	DeleteTableSession(ctx context.Context, id uuid.UUID) error
	DeleteExpiredTableSessions(ctx context.Context) error
}
