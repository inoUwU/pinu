package table

import (
	"context"
	"time"

	"github.com/samber/do"

	"inoUwU/pinu/app/domain/port"
	"inoUwU/pinu/app/domain/table"
	"inoUwU/pinu/app/usecases/table/input"
	"inoUwU/pinu/app/usecases/table/output"
)

// ITableUsecase テーブルユースケースのインターフェース
type ITableUsecase interface {
	GetAllTables(ctx context.Context, input *input.GetTablesInput) (*output.GetTablesOutput, error)
	GetTableByID(ctx context.Context, input *input.GetTableByIDInput) (*output.GetTableByIDOutput, error)
	GetTablesByStatus(ctx context.Context, input *input.GetTablesByStatusInput) (*output.GetTablesByStatusOutput, error)
	CreateTable(ctx context.Context, input *input.CreateTableInput) (*output.CreateTableOutput, error)
	UpdateTable(ctx context.Context, input *input.UpdateTableInput) (*output.UpdateTableOutput, error)
	UpdateTableStatus(ctx context.Context, input *input.UpdateTableStatusInput) (*output.UpdateTableStatusOutput, error)
	DeleteTable(ctx context.Context, input *input.DeleteTableInput) (*output.DeleteTableOutput, error)
}

// TableUsecaseImpl テーブルユースケースの実装
type TableUsecaseImpl struct {
	tableRepo table.TableRepository
	logger    port.Logger
}

// NewTableUsecase テーブルユースケースを生成する
func NewTableUsecase(i *do.Injector) (ITableUsecase, error) {
	repository := do.MustInvoke[table.TableRepository](i)
	logger := do.MustInvokeNamed[port.Logger](i, "logger")

	return &TableUsecaseImpl{
		tableRepo: repository,
		logger:    logger,
	}, nil
}

// GetAllTables 全てのテーブルを取得する
func (u *TableUsecaseImpl) GetAllTables(ctx context.Context, input *input.GetTablesInput) (*output.GetTablesOutput, error) {
	u.logger.Info("getting all tables")

	tables, err := u.tableRepo.GetAll(ctx)
	if err != nil {
		u.logger.Error("failed to get all tables", "error", err)
		return nil, err
	}

	u.logger.Info("successfully retrieved tables", "count", len(tables))

	return &output.GetTablesOutput{
		Tables: tables,
		Count:  len(tables),
	}, nil
}

// GetTableByID 指定IDのテーブルを取得する
func (u *TableUsecaseImpl) GetTableByID(ctx context.Context, input *input.GetTableByIDInput) (*output.GetTableByIDOutput, error) {
	u.logger.Info("getting table by ID", "tableID", input.TableID)

	table, err := u.tableRepo.GetByID(ctx, input.TableID)
	if err != nil {
		u.logger.Error("failed to get table by ID", "tableID", input.TableID, "error", err)
		return nil, err
	}

	if table == nil {
		u.logger.Warn("table not found", "tableID", input.TableID)
		return &output.GetTableByIDOutput{
			Table: nil,
		}, nil
	}

	u.logger.Info("successfully retrieved table", "tableID", input.TableID)

	return &output.GetTableByIDOutput{
		Table: table,
	}, nil
}

// GetTablesByStatus 指定ステータスのテーブルを取得する
func (u *TableUsecaseImpl) GetTablesByStatus(ctx context.Context, input *input.GetTablesByStatusInput) (*output.GetTablesByStatusOutput, error) {
	u.logger.Info("getting tables by status", "status", input.Status)

	tables, err := u.tableRepo.GetByStatus(ctx, input.Status)
	if err != nil {
		u.logger.Error("failed to get tables by status", "status", input.Status, "error", err)
		return nil, err
	}

	u.logger.Info("successfully retrieved tables by status", "status", input.Status, "count", len(tables))

	return &output.GetTablesByStatusOutput{
		Tables: tables,
		Count:  len(tables),
		Status: input.Status,
	}, nil
}

// CreateTable テーブルを作成する
func (u *TableUsecaseImpl) CreateTable(ctx context.Context, input *input.CreateTableInput) (*output.CreateTableOutput, error) {
	u.logger.Info("creating table", "tableID", input.TableID, "status", input.Status)

	// TODO: バリデーションを追加
	// TODO: トランザクション処理を追加

	tableEntity := &table.Table{
		TableID:         input.TableID,
		Status:          input.Status,
		CurrentOrdersID: input.CurrentOrdersID,
		LastUpdated:     time.Now(),
	}

	if err := u.tableRepo.Create(ctx, tableEntity); err != nil {
		u.logger.Error("failed to create table", "tableID", input.TableID, "error", err)
		return nil, err
	}

	// 作成されたテーブルを取得
	createdTable, err := u.tableRepo.GetByID(ctx, input.TableID)
	if err != nil {
		u.logger.Error("failed to get created table", "tableID", input.TableID, "error", err)
		return nil, err
	}

	u.logger.Info("table created successfully", "tableID", input.TableID)

	return &output.CreateTableOutput{
		Table:     createdTable,
		Message:   "Table created successfully",
		CreatedAt: time.Now(),
		TableID:   createdTable.TableID,
		Status:    createdTable.Status,
	}, nil
}

// UpdateTable テーブルを更新する
func (u *TableUsecaseImpl) UpdateTable(ctx context.Context, input *input.UpdateTableInput) (*output.UpdateTableOutput, error) {
	u.logger.Info("updating table", "tableID", input.TableID)

	// TODO: バリデーションを追加
	// TODO: トランザクション処理を追加

	// 既存のテーブルが存在するかチェック
	existingTable, err := u.tableRepo.GetByID(ctx, input.TableID)
	if err != nil {
		u.logger.Error("failed to get table for update", "tableID", input.TableID, "error", err)
		return nil, err
	}

	if existingTable == nil {
		u.logger.Warn("table not found for update", "tableID", input.TableID)
		return nil, err // TODO: カスタムエラーに変更
	}

	tableEntity := &table.Table{
		TableID:         input.TableID,
		Status:          input.Status,
		CurrentOrdersID: input.CurrentOrdersID,
		LastUpdated:     time.Now(),
	}

	if err := u.tableRepo.Update(ctx, tableEntity); err != nil {
		u.logger.Error("failed to update table", "tableID", input.TableID, "error", err)
		return nil, err
	}

	// 更新されたテーブルを取得
	updatedTable, err := u.tableRepo.GetByID(ctx, input.TableID)
	if err != nil {
		u.logger.Error("failed to get updated table", "tableID", input.TableID, "error", err)
		return nil, err
	}

	u.logger.Info("table updated successfully", "tableID", input.TableID)

	return &output.UpdateTableOutput{
		Table:     updatedTable,
		Message:   "Table updated successfully",
		UpdatedAt: time.Now(),
	}, nil
}

// UpdateTableStatus テーブルのステータスを更新する
func (u *TableUsecaseImpl) UpdateTableStatus(ctx context.Context, input *input.UpdateTableStatusInput) (*output.UpdateTableStatusOutput, error) {
	u.logger.Info("updating table status", "tableID", input.TableID, "status", input.Status)

	// TODO: バリデーションを追加
	// TODO: トランザクション処理を追加

	// 既存のテーブルが存在するかチェック
	existingTable, err := u.tableRepo.GetByID(ctx, input.TableID)
	if err != nil {
		u.logger.Error("failed to get table for status update", "tableID", input.TableID, "error", err)
		return nil, err
	}

	if existingTable == nil {
		u.logger.Warn("table not found for status update", "tableID", input.TableID)
		return nil, err // TODO: カスタムエラーに変更
	}

	if err := u.tableRepo.UpdateStatus(ctx, input.TableID, input.Status); err != nil {
		u.logger.Error("failed to update table status", "tableID", input.TableID, "status", input.Status, "error", err)
		return nil, err
	}

	// 更新されたテーブルを取得
	updatedTable, err := u.tableRepo.GetByID(ctx, input.TableID)
	if err != nil {
		u.logger.Error("failed to get updated table", "tableID", input.TableID, "error", err)
		return nil, err
	}

	u.logger.Info("table status updated successfully", "tableID", input.TableID, "status", input.Status)

	return &output.UpdateTableStatusOutput{
		Table:     updatedTable,
		Message:   "Table status updated successfully",
		UpdatedAt: time.Now(),
		TableID:   updatedTable.TableID,
		Status:    updatedTable.Status,
	}, nil
}

// DeleteTable テーブルを削除する
func (u *TableUsecaseImpl) DeleteTable(ctx context.Context, input *input.DeleteTableInput) (*output.DeleteTableOutput, error) {
	u.logger.Info("deleting table", "tableID", input.TableID)

	// TODO: 関連データの存在チェック
	// TODO: トランザクション処理を追加

	// 既存のテーブルが存在するかチェック
	existingTable, err := u.tableRepo.GetByID(ctx, input.TableID)
	if err != nil {
		u.logger.Error("failed to get table for delete", "tableID", input.TableID, "error", err)
		return nil, err
	}

	if existingTable == nil {
		u.logger.Warn("table not found for delete", "tableID", input.TableID)
		return nil, err // TODO: カスタムエラーに変更
	}

	if err := u.tableRepo.Delete(ctx, input.TableID); err != nil {
		u.logger.Error("failed to delete table", "tableID", input.TableID, "error", err)
		return nil, err
	}

	u.logger.Info("table deleted successfully", "tableID", input.TableID)

	return &output.DeleteTableOutput{
		TableID:   input.TableID,
		Message:   "Table deleted successfully",
		DeletedAt: time.Now(),
	}, nil
}
