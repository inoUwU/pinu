package menu

import (
	"context"
)

// MenuRepository メニューリポジトリのインターフェース
type MenuRepository interface {
	// 全てのメニューを取得する
	GetAll(ctx context.Context) ([]*Menu, error)

	// IDでメニューを取得する
	GetByID(ctx context.Context, id string) (*Menu, error)

	// カテゴリIDでメニューを取得する
	GetByCategory(ctx context.Context, categoryID string) ([]*Menu, error)

	// メニューを作成する
	Create(ctx context.Context, menu *Menu) error

	// メニューを更新する
	Update(ctx context.Context, menu *Menu) error

	// メニューを削除する
	Delete(ctx context.Context, id string) error
}
