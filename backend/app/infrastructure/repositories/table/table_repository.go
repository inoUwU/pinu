package table

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/samber/do"
	"github.com/uptrace/bun"

	"inoUwU/pinu/app/domain/table"
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

func (r *tableRepository) Create(ctx context.Context, tbl *table.Table) error {
	var currentOrdersID *string
	if tbl.CurrentOrdersID != nil {
		id := tbl.CurrentOrdersID.String()
		currentOrdersID = &id
	}

	model := &models.TableModel{
		TableID:         string(tbl.TableID),
		Status:          string(tbl.Status),
		CurrentOrdersID: currentOrdersID,
		LastUpdated:     tbl.LastUpdated,
	}

	_, err := r.db.NewInsert().Model(model).Exec(ctx)
	return err
}

func (r *tableRepository) GetByID(ctx context.Context, id table.TableID) (*table.Table, error) {
	model := &models.TableModel{}
	err := r.db.NewSelect().
		Model(model).
		Where("table_id = ?", string(id)).
		Scan(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	var currentOrdersID *uuid.UUID
	if model.CurrentOrdersID != nil {
		if parsedUUID, err := uuid.Parse(*model.CurrentOrdersID); err == nil {
			currentOrdersID = &parsedUUID
		}
	}

	return &table.Table{
		TableID:         table.TableID(model.TableID),
		Status:          table.TableStatus(model.Status),
		CurrentOrdersID: currentOrdersID,
		LastUpdated:     model.LastUpdated,
	}, nil
}

func (r *tableRepository) GetAll(ctx context.Context) ([]*table.Table, error) {
	var models []*models.TableModel
	err := r.db.NewSelect().
		Model(&models).
		Order("table_id ASC").
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	tables := make([]*table.Table, len(models))
	for i, model := range models {
		var currentOrdersID *uuid.UUID
		if model.CurrentOrdersID != nil {
			if parsedUUID, err := uuid.Parse(*model.CurrentOrdersID); err == nil {
				currentOrdersID = &parsedUUID
			}
		}

		tables[i] = &table.Table{
			TableID:         table.TableID(model.TableID),
			Status:          table.TableStatus(model.Status),
			CurrentOrdersID: currentOrdersID,
			LastUpdated:     model.LastUpdated,
		}
	}

	return tables, nil
}

func (r *tableRepository) GetByStatus(ctx context.Context, status table.TableStatus) ([]*table.Table, error) {
	var models []*models.TableModel
	err := r.db.NewSelect().
		Model(&models).
		Where("status = ?", string(status)).
		Order("table_id ASC").
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	tables := make([]*table.Table, len(models))
	for i, model := range models {
		var currentOrdersID *uuid.UUID
		if model.CurrentOrdersID != nil {
			if parsedUUID, err := uuid.Parse(*model.CurrentOrdersID); err == nil {
				currentOrdersID = &parsedUUID
			}
		}

		tables[i] = &table.Table{
			TableID:         table.TableID(model.TableID),
			Status:          table.TableStatus(model.Status),
			CurrentOrdersID: currentOrdersID,
			LastUpdated:     model.LastUpdated,
		}
	}

	return tables, nil
}

func (r *tableRepository) Update(ctx context.Context, tbl *table.Table) error {
	var currentOrdersID *string
	if tbl.CurrentOrdersID != nil {
		id := tbl.CurrentOrdersID.String()
		currentOrdersID = &id
	}

	model := &models.TableModel{
		TableID:         string(tbl.TableID),
		Status:          string(tbl.Status),
		CurrentOrdersID: currentOrdersID,
		LastUpdated:     time.Now(),
	}

	_, err := r.db.NewUpdate().
		Model(model).
		Where("table_id = ?", string(tbl.TableID)).
		Exec(ctx)

	return err
}

func (r *tableRepository) Delete(ctx context.Context, id table.TableID) error {
	_, err := r.db.NewDelete().
		Model((*models.TableModel)(nil)).
		Where("table_id = ?", string(id)).
		Exec(ctx)

	return err
}

func (r *tableRepository) UpdateStatus(ctx context.Context, id table.TableID, status table.TableStatus) error {
	_, err := r.db.NewUpdate().
		Model((*models.TableModel)(nil)).
		Set("status = ?, last_updated = ?", string(status), time.Now()).
		Where("table_id = ?", string(id)).
		Exec(ctx)

	return err
}
