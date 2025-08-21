package menu_option

import (
	"context"
)

// MenuOptionRepository メニューオプションリポジトリのインターface
type MenuOptionRepository interface {
	Create(ctx context.Context, menuOption *MenuOption) error
	GetByID(ctx context.Context, id MenuOptionID) (*MenuOption, error)
	GetAll(ctx context.Context) ([]*MenuOption, error)
	Update(ctx context.Context, menuOption *MenuOption) error
	Delete(ctx context.Context, id MenuOptionID) error
}
