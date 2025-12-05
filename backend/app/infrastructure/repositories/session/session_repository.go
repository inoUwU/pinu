package session

import (
	"context"
	ctxkey "inoUwU/pinu/app/infrastructure/ctx"

	"inoUwU/pinu/app/domain/session"
	"inoUwU/pinu/app/infrastructure/models"

	"github.com/google/uuid"
	"github.com/samber/do"
	"github.com/uptrace/bun"
)

// SessionRepositoryImpl セッションリポジトリの実装（アダプター）
type SessionRepositoryImpl struct {
	db *bun.DB
}

// NewSessionRepository セッションリポジトリの実装を生成する
func NewSessionRepository(i *do.Injector) (session.SessionRepository, error) {
	db := do.MustInvokeNamed[*bun.DB](i, "db")
	return &SessionRepositoryImpl{
		db: db,
	}, nil
}

// CreateSession セッションを作成する
func (r SessionRepositoryImpl) CreateSession(ctx context.Context, session *session.Session) error {
	newSession := &models.Session{
		SESSION_ID:    session.SessionID,
		USER_ID:       string(session.UserID),
		REFRESH_TOKEN: session.RefreshToken,
		IS_REVOKED:    session.IsRevoked,
		CREATED_AT:    session.CreatedAt,
		EXPIRES_AT:    session.ExpiresAt,
	}

	var inserter *bun.InsertQuery
	if tx, ok := ctx.Value(ctxkey.TxCtxKey).(bun.Tx); ok {
		inserter = tx.NewInsert()
	} else {
		inserter = r.db.NewInsert()
	}

	if _, err := inserter.Model(newSession).Exec(ctx); err != nil {
		return err
	}

	return nil
}

// GetSessionByID セッションIDでセッションを取得する
func (r SessionRepositoryImpl) GetSessionByID(ctx context.Context, id string) (*session.Session, error) {
	res := new(session.Session)

	var selector *bun.SelectQuery

	if tx, ok := ctx.Value(ctxkey.TxCtxKey).(bun.Tx); ok {
		selector = tx.NewSelect()
	} else {
		selector = r.db.NewSelect()
	}

	selector.Model(res).Where("session_id = ?", id)
	if err := selector.Scan(ctx); err != nil {
		return nil, nil // ユーザーが見つからない場合はnilを返す
	}
	return res, nil
}

// RevokeSessionByID セッションIDでセッションを無効化する
func (r SessionRepositoryImpl) RevokeSessionByID(ctx context.Context, id string) error {

	var selector *bun.SelectQuery
	var updater *bun.UpdateQuery
	if tx, ok := ctx.Value(ctxkey.TxCtxKey).(bun.Tx); ok {
		selector = tx.NewSelect()
		updater = tx.NewUpdate()
	} else {
		selector = r.db.NewSelect()
		updater = r.db.NewUpdate()
	}

	// 更新対象を取得
	currentSession := new(session.Session)
	if err := selector.Model(currentSession).Where("session_id = ?", id).Scan(ctx); err != nil {
		return err
	}

	// 無効化
	currentSession.IsRevoked = true
	newSession := &models.Session{
		SESSION_ID:    currentSession.SessionID,
		USER_ID:       string(currentSession.UserID),
		REFRESH_TOKEN: currentSession.RefreshToken,
		IS_REVOKED:    currentSession.IsRevoked,
		CREATED_AT:    currentSession.CreatedAt,
		EXPIRES_AT:    currentSession.ExpiresAt,
	}

	if _, err := updater.Model(newSession).WherePK().Exec(ctx); err != nil {
		return err
	}
	return nil
}

// DeleteSession セッションIDでセッションを削除する
func (r SessionRepositoryImpl) DeleteSession(ctx context.Context, id string) error {
	var deleter *bun.DeleteQuery
	if tx, ok := ctx.Value(ctxkey.TxCtxKey).(bun.Tx); ok {
		deleter = tx.NewDelete()
	} else {
		deleter = r.db.NewDelete()
	}

	if _, err := deleter.Where("session_id = ?", id).Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (r SessionRepositoryImpl) CreateTableSession(ctx context.Context, tableSession *session.TableSession) error {
	//TODO implement me
	panic("implement me")
}

func (r SessionRepositoryImpl) GetTableSessionByID(ctx context.Context, id uuid.UUID) (*session.TableSession, error) {
	//TODO implement me
	panic("implement me")
}

func (r SessionRepositoryImpl) GetTableSessionByTableID(ctx context.Context, tableID string) (*session.TableSession, error) {
	//TODO implement me
	panic("implement me")
}

func (r SessionRepositoryImpl) UpdateTableSessionLastUsed(ctx context.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (r SessionRepositoryImpl) DeleteTableSession(ctx context.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (r SessionRepositoryImpl) DeleteExpiredTableSessions(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
