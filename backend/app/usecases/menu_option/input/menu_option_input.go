package input

import "inoUwU/pinu/app/domain/menu_option"

// GetMenuOptionsInput メニューオプション一覧取得の入力
type GetMenuOptionsInput struct {
	// 将来的に検索条件やページング情報を追加
}

// GetMenuOptionByIDInput メニューオプション単一取得の入力
type GetMenuOptionByIDInput struct {
	MenuOptionID menu_option.MenuOptionID `json:"menu_option_id" validate:"required"`
}

// CreateMenuOptionInput メニューオプション作成の入力
type CreateMenuOptionInput struct {
	MenuOptionID menu_option.MenuOptionID `json:"menu_option_id" validate:"required"`
	Name         string                   `json:"name" validate:"required"`
	Price        float64                  `json:"price" validate:"required,min=0"`
}

// UpdateMenuOptionInput メニューオプション更新の入力
type UpdateMenuOptionInput struct {
	MenuOptionID menu_option.MenuOptionID `json:"menu_option_id" validate:"required"`
	Name         string                   `json:"name" validate:"required"`
	Price        float64                  `json:"price" validate:"required,min=0"`
}

// DeleteMenuOptionInput メニューオプション削除の入力
type DeleteMenuOptionInput struct {
	MenuOptionID menu_option.MenuOptionID `json:"menu_option_id" validate:"required"`
}
