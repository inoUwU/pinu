package output

import (
	"time"

	"inoUwU/pinu/app/domain/table"
)

// GetTablesOutput テーブル一覧取得の出力
type GetTablesOutput struct {
	Tables []*table.Table `json:"tables"`
	Count  int            `json:"count"`
}

// GetTableByIDOutput テーブル単一取得の出力
type GetTableByIDOutput struct {
	Table *table.Table `json:"table"`
}

// GetTablesByStatusOutput ステータス別テーブル取得の出力
type GetTablesByStatusOutput struct {
	Tables []*table.Table    `json:"tables"`
	Count  int               `json:"count"`
	Status table.TableStatus `json:"status"`
}

// CreateTableOutput テーブル作成の出力
type CreateTableOutput struct {
	Table     *table.Table      `json:"table"`
	Message   string            `json:"message"`
	CreatedAt time.Time         `json:"created_at"`
	TableID   table.TableID     `json:"table_id"`
	Status    table.TableStatus `json:"status"`
}

// UpdateTableOutput テーブル更新の出力
type UpdateTableOutput struct {
	Table     *table.Table `json:"table"`
	Message   string       `json:"message"`
	UpdatedAt time.Time    `json:"updated_at"`
}

// UpdateTableStatusOutput テーブルステータス更新の出力
type UpdateTableStatusOutput struct {
	Table     *table.Table      `json:"table"`
	Message   string            `json:"message"`
	UpdatedAt time.Time         `json:"updated_at"`
	TableID   table.TableID     `json:"table_id"`
	Status    table.TableStatus `json:"status"`
}

// DeleteTableOutput テーブル削除の出力
type DeleteTableOutput struct {
	TableID   table.TableID `json:"table_id"`
	Message   string        `json:"message"`
	DeletedAt time.Time     `json:"deleted_at"`
}

// CheckoutTableOutput テーブル会計処理の出力
type CheckoutTableOutput struct {
	Table     *table.Table      `json:"table"`
	Message   string            `json:"message"`
	UpdatedAt time.Time         `json:"updated_at"`
	TableID   table.TableID     `json:"table_id"`
	Status    table.TableStatus `json:"status"`
}
