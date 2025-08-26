package output

import (
	"time"

	"inoUwU/pinu/app/domain/menu_option"
)

// GetMenuOptionsOutput メニューオプション一覧取得の出力
type GetMenuOptionsOutput struct {
	MenuOptions []*menu_option.MenuOption `json:"menu_options"`
	Count       int                       `json:"count"`
}

// GetMenuOptionByIDOutput メニューオプション単一取得の出力
type GetMenuOptionByIDOutput struct {
	MenuOption *menu_option.MenuOption `json:"menu_option"`
}

// CreateMenuOptionOutput メニューオプション作成の出力
type CreateMenuOptionOutput struct {
	MenuOption   *menu_option.MenuOption  `json:"menu_option"`
	Message      string                   `json:"message"`
	CreatedAt    time.Time                `json:"created_at"`
	MenuOptionID menu_option.MenuOptionID `json:"menu_option_id"`
	Name         string                   `json:"name"`
	Price        float64                  `json:"price"`
}

// UpdateMenuOptionOutput メニューオプション更新の出力
type UpdateMenuOptionOutput struct {
	MenuOption *menu_option.MenuOption `json:"menu_option"`
	Message    string                  `json:"message"`
	UpdatedAt  time.Time               `json:"updated_at"`
}

// DeleteMenuOptionOutput メニューオプション削除の出力
type DeleteMenuOptionOutput struct {
	MenuOptionID menu_option.MenuOptionID `json:"menu_option_id"`
	Message      string                   `json:"message"`
	DeletedAt    time.Time                `json:"deleted_at"`
}
