package category

import (
	"context"
	"github.com/google/uuid"
)

type Repository interface {
	// カテゴリーが存在するか確認する
	Exists(ctx context.Context, id uuid.UUID) (bool, error)

	// IDでカテゴリーを取得する
	GetCategoryByID(ctx context.Context, id uuid.UUID) (*Category, error)

	// カテゴリーを作成する
	CreateCategory(ctx context.Context, category *Category) error

	// カテゴリーを更新する
	UpdateCategory(ctx context.Context, category *Category) error

	// カテゴリーを削除する
	DeleteCategory(ctx context.Context, id uuid.UUID) error
}
