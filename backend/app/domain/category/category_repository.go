package category

import (
	"context"
)

// CategoryRepository カテゴリリポジトリのインターface
type CategoryRepository interface {
	Create(ctx context.Context, category *Category) error
	GetByID(ctx context.Context, id string) (*Category, error)
	GetAll(ctx context.Context) ([]*Category, error)
	Update(ctx context.Context, category *Category) error
	Delete(ctx context.Context, id string) error
}
