package table

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/samber/do"
	"github.com/uptrace/bun"

	"inoUwU/pinu/app/domain/table"
	ctxkey "inoUwU/pinu/app/infrastructure/ctx"
	"inoUwU/pinu/app/infrastructure/models"
)

type tableRepository struct {
	db *bun.DB
}

// NewTableRepository テーブルリポジトリの新規作成
func NewTableRepository(i *do.Injector) table.TableRepository {
	db := do.MustInvokeNamed[*bun.DB](i, "db")
	return &tableRepository{db: db}
}

// getIDB トランザクション context がある場合は Tx、なければ DB を返す
func (r *tableRepository) getIDB(ctx context.Context) bun.IDB {
	if tx, ok := ctx.Value(ctxkey.TxCtxKey).(bun.Tx); ok {
		return tx
	}
	return r.db
}

func (r *tableRepository) Create(ctx context.Context, tbl *table.Table) error {
	var currentTableSessionID *string
	if tbl.CurrentTableSessionID != nil {
		id := tbl.CurrentTableSessionID.String()
		currentTableSessionID = &id
	}

	model := &models.TableModel{
		TableID:               string(tbl.TableID),
		Status:                string(tbl.Status),
		CurrentTableSessionID: currentTableSessionID,
		LastUpdated:           tbl.LastUpdated,
	}

	_, err := r.getIDB(ctx).NewInsert().Model(model).Exec(ctx)
	return err
}

func (r *tableRepository) GetByID(ctx context.Context, id table.TableID) (*table.Table, error) {
	model := &models.TableModel{}
	err := r.getIDB(ctx).NewSelect().
		Model(model).
		Where("table_id = ?", string(id)).
		Scan(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	var currentTableSessionID *uuid.UUID
	if model.CurrentTableSessionID != nil {
		if parsedUUID, err := uuid.Parse(*model.CurrentTableSessionID); err == nil {
			currentTableSessionID = &parsedUUID
		}
	}

	return &table.Table{
		TableID:               table.TableID(model.TableID),
		Status:                table.TableStatus(model.Status),
		CurrentTableSessionID: currentTableSessionID,
		LastUpdated:           model.LastUpdated,
	}, nil
}

func (r *tableRepository) GetAll(ctx context.Context) ([]*table.Table, error) {
	var models []*models.TableModel
	err := r.getIDB(ctx).NewSelect().
		Model(&models).
		Order("table_id ASC").
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	tables := make([]*table.Table, len(models))
	for i, model := range models {
		var currentTableSessionID *uuid.UUID
		if model.CurrentTableSessionID != nil {
			if parsedUUID, err := uuid.Parse(*model.CurrentTableSessionID); err == nil {
				currentTableSessionID = &parsedUUID
			}
		}

		tables[i] = &table.Table{
			TableID:               table.TableID(model.TableID),
			Status:                table.TableStatus(model.Status),
			CurrentTableSessionID: currentTableSessionID,
			LastUpdated:           model.LastUpdated,
		}
	}

	return tables, nil
}

func (r *tableRepository) GetByStatus(ctx context.Context, status table.TableStatus) ([]*table.Table, error) {
	var models []*models.TableModel
	err := r.getIDB(ctx).NewSelect().
		Model(&models).
		Where("status = ?", string(status)).
		Order("table_id ASC").
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	tables := make([]*table.Table, len(models))
	for i, model := range models {
		var currentTableSessionID *uuid.UUID
		if model.CurrentTableSessionID != nil {
			if parsedUUID, err := uuid.Parse(*model.CurrentTableSessionID); err == nil {
				currentTableSessionID = &parsedUUID
			}
		}

		tables[i] = &table.Table{
			TableID:               table.TableID(model.TableID),
			Status:                table.TableStatus(model.Status),
			CurrentTableSessionID: currentTableSessionID,
			LastUpdated:           model.LastUpdated,
		}
	}

	return tables, nil
}

func (r *tableRepository) Update(ctx context.Context, tbl *table.Table) error {
	var currentTableSessionID *string
	if tbl.CurrentTableSessionID != nil {
		id := tbl.CurrentTableSessionID.String()
		currentTableSessionID = &id
	}

	model := &models.TableModel{
		TableID:               string(tbl.TableID),
		Status:                string(tbl.Status),
		CurrentTableSessionID: currentTableSessionID,
		LastUpdated:           time.Now(),
	}

	_, err := r.getIDB(ctx).NewUpdate().
		Model(model).
		Where("table_id = ?", string(tbl.TableID)).
		Exec(ctx)

	return err
}

func (r *tableRepository) Delete(ctx context.Context, id table.TableID) error {
	_, err := r.getIDB(ctx).NewDelete().
		Model((*models.TableModel)(nil)).
		Where("table_id = ?", string(id)).
		Exec(ctx)

	return err
}

func (r *tableRepository) UpdateStatus(ctx context.Context, id table.TableID, status table.TableStatus) error {
	_, err := r.getIDB(ctx).NewUpdate().
		Model((*models.TableModel)(nil)).
		Set("status = ?, last_updated = ?", string(status), time.Now()).
		Where("table_id = ?", string(id)).
		Exec(ctx)

	return err
}

// ClearTableSession テーブルの current_table_session_id をクリアする
func (r *tableRepository) ClearTableSession(ctx context.Context, id table.TableID) error {
	_, err := r.getIDB(ctx).NewUpdate().
		Model((*models.TableModel)(nil)).
		Set("current_table_session_id = NULL, last_updated = ?", time.Now()).
		Where("table_id = ?", string(id)).
		Exec(ctx)

	return err
}
