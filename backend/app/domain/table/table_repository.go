package table

import (
	"context"
)

// TableRepository テーブルリポジトリのインターface
type TableRepository interface {
	Create(ctx context.Context, table *Table) error
	GetByID(ctx context.Context, id TableID) (*Table, error)
	GetAll(ctx context.Context) ([]*Table, error)
	GetByStatus(ctx context.Context, status TableStatus) ([]*Table, error)
	Update(ctx context.Context, table *Table) error
	Delete(ctx context.Context, id TableID) error
	UpdateStatus(ctx context.Context, id TableID, status TableStatus) error
}
