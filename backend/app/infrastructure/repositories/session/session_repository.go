package session

import (
	"context"
	"database/sql"
	ctxkey "inoUwU/pinu/app/infrastructure/ctx"
	"time"

	"inoUwU/pinu/app/domain/session"
	"inoUwU/pinu/app/domain/user"
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
	newSession := &models.SessionModel{
		SessionID:    session.SessionID,
		UserID:       string(session.UserID),
		RefreshToken: session.RefreshToken,
		IsRevoked:    session.IsRevoked,
		CreatedAt:    session.CreatedAt,
		ExpiresAt:    session.ExpiresAt,
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
	model := new(models.SessionModel)

	var selector *bun.SelectQuery

	if tx, ok := ctx.Value(ctxkey.TxCtxKey).(bun.Tx); ok {
		selector = tx.NewSelect()
	} else {
		selector = r.db.NewSelect()
	}

	selector.Model(model).Where("session_id = ?", id)
	if err := selector.Scan(ctx); err != nil {
		return nil, nil // セッションが見つからない場合はnilを返す
	}

	return mapSessionModelToDomain(model), nil
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
	currentModel := new(models.SessionModel)
	if err := selector.Model(currentModel).Where("session_id = ?", id).Scan(ctx); err != nil {
		return err
	}

	// 無効化
	currentModel.IsRevoked = true

	if _, err := updater.Model(currentModel).WherePK().Exec(ctx); err != nil {
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

	if _, err := deleter.Model((*models.SessionModel)(nil)).Where("session_id = ?", id).Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (r SessionRepositoryImpl) CreateTableSession(ctx context.Context, tableSession *session.TableSession) error {
	newTableSession := &models.TableSessionModel{
		TableSessionID: tableSession.TableSessionID,
		TableID:        tableSession.TableID,
		IsRevoked:      tableSession.IsRevoked,
		CreatedAt:      tableSession.CreatedAt,
		LastUsed:       tableSession.LastUsed,
		ExpiresAt:      tableSession.ExpiresAt,
	}

	var inserter *bun.InsertQuery
	if tx, ok := ctx.Value(ctxkey.TxCtxKey).(bun.Tx); ok {
		inserter = tx.NewInsert()
	} else {
		inserter = r.db.NewInsert()
	}

	if _, err := inserter.Model(newTableSession).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (r SessionRepositoryImpl) GetTableSessionByID(ctx context.Context, id uuid.UUID) (*session.TableSession, error) {
	tableSession := &models.TableSessionModel{}

	var selector *bun.SelectQuery
	if tx, ok := ctx.Value(ctxkey.TxCtxKey).(bun.Tx); ok {
		selector = tx.NewSelect()
	} else {
		selector = r.db.NewSelect()
	}

	err := selector.Model(tableSession).Where("table_session_id = ?", id).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &session.TableSession{
		TableSessionID: tableSession.TableSessionID,
		TableID:        tableSession.TableID,
		IsRevoked:      tableSession.IsRevoked,
		CreatedAt:      tableSession.CreatedAt,
		LastUsed:       tableSession.LastUsed,
		ExpiresAt:      tableSession.ExpiresAt,
	}, nil
}

func (r SessionRepositoryImpl) GetTableSessionByTableID(ctx context.Context, tableID string) (*session.TableSession, error) {
	tableSession := &models.TableSessionModel{}

	var selector *bun.SelectQuery
	if tx, ok := ctx.Value(ctxkey.TxCtxKey).(bun.Tx); ok {
		selector = tx.NewSelect()
	} else {
		selector = r.db.NewSelect()
	}

	err := selector.Model(tableSession).
		Where("table_id = ?", tableID).
		Where("is_revoked = ?", false).
		Order("created_at DESC").
		Limit(1).
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &session.TableSession{
		TableSessionID: tableSession.TableSessionID,
		TableID:        tableSession.TableID,
		IsRevoked:      tableSession.IsRevoked,
		CreatedAt:      tableSession.CreatedAt,
		LastUsed:       tableSession.LastUsed,
		ExpiresAt:      tableSession.ExpiresAt,
	}, nil
}

func (r SessionRepositoryImpl) UpdateTableSessionLastUsed(ctx context.Context, id uuid.UUID) error {
	var updater *bun.UpdateQuery
	if tx, ok := ctx.Value(ctxkey.TxCtxKey).(bun.Tx); ok {
		updater = tx.NewUpdate()
	} else {
		updater = r.db.NewUpdate()
	}

	if _, err := updater.Model((*models.TableSessionModel)(nil)).
		Set("last_used = ?", time.Now()).
		Where("table_session_id = ?", id).
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (r SessionRepositoryImpl) DeleteTableSession(ctx context.Context, id uuid.UUID) error {
	var deleter *bun.DeleteQuery
	if tx, ok := ctx.Value(ctxkey.TxCtxKey).(bun.Tx); ok {
		deleter = tx.NewDelete()
	} else {
		deleter = r.db.NewDelete()
	}

	if _, err := deleter.Model((*models.TableSessionModel)(nil)).Where("table_session_id = ?", id).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (r SessionRepositoryImpl) DeleteExpiredTableSessions(ctx context.Context) error {
	var deleter *bun.DeleteQuery
	if tx, ok := ctx.Value(ctxkey.TxCtxKey).(bun.Tx); ok {
		deleter = tx.NewDelete()
	} else {
		deleter = r.db.NewDelete()
	}

	if _, err := deleter.Model((*models.TableSessionModel)(nil)).Where("expires_at < ?", time.Now()).Exec(ctx); err != nil {
		return err
	}

	return nil
}

// RevokeTableSessionByTableID テーブルIDに紐づく全テーブルセッションを無効化する
func (r SessionRepositoryImpl) RevokeTableSessionByTableID(ctx context.Context, tableID string) error {
	var updater *bun.UpdateQuery
	if tx, ok := ctx.Value(ctxkey.TxCtxKey).(bun.Tx); ok {
		updater = tx.NewUpdate()
	} else {
		updater = r.db.NewUpdate()
	}

	if _, err := updater.Model((*models.TableSessionModel)(nil)).
		Set("is_revoked = ?", true).
		Where("table_id = ?", tableID).
		Where("is_revoked = ?", false).
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

// mapSessionModelToDomain SessionModelからドメイン型へのマッピング
func mapSessionModelToDomain(m *models.SessionModel) *session.Session {
	return &session.Session{
		SessionID:    m.SessionID,
		UserID:       user.UserID(m.UserID),
		RefreshToken: m.RefreshToken,
		IsRevoked:    m.IsRevoked,
		CreatedAt:    m.CreatedAt,
		ExpiresAt:    m.ExpiresAt,
	}
}
