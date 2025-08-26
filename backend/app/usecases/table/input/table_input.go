package input

import (
	"github.com/google/uuid"
	"inoUwU/pinu/app/domain/table"
)

// GetTablesInput テーブル一覧取得の入力
type GetTablesInput struct {
	// 将来的に検索条件やページング情報を追加
}

// GetTableByIDInput テーブル単一取得の入力
type GetTableByIDInput struct {
	TableID table.TableID `json:"table_id" validate:"required"`
}

// GetTablesByStatusInput ステータス別テーブル取得の入力
type GetTablesByStatusInput struct {
	Status table.TableStatus `json:"status" validate:"required"`
}

// CreateTableInput テーブル作成の入力
type CreateTableInput struct {
	TableID         table.TableID     `json:"table_id" validate:"required"`
	Status          table.TableStatus `json:"status" validate:"required"`
	CurrentOrdersID *uuid.UUID        `json:"current_orders_id"`
}

// UpdateTableInput テーブル更新の入力
type UpdateTableInput struct {
	TableID         table.TableID     `json:"table_id" validate:"required"`
	Status          table.TableStatus `json:"status" validate:"required"`
	CurrentOrdersID *uuid.UUID        `json:"current_orders_id"`
}

// UpdateTableStatusInput テーブルステータス更新の入力
type UpdateTableStatusInput struct {
	TableID table.TableID     `json:"table_id" validate:"required"`
	Status  table.TableStatus `json:"status" validate:"required"`
}

// DeleteTableInput テーブル削除の入力
type DeleteTableInput struct {
	TableID table.TableID `json:"table_id" validate:"required"`
}
