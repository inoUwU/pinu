package input

// GetMenusInput メニュー一覧取得の入力
type GetMenusInput struct {
	// 将来的に検索条件やページング情報を追加
}

// GetMenuByIDInput メニュー単一取得の入力
type GetMenuByIDInput struct {
	MenuID string `json:"menu_id" validate:"required"`
}

// GetMenusByCategoryInput カテゴリ別メニュー取得の入力
type GetMenusByCategoryInput struct {
	CategoryID string `json:"category_id" validate:"required"`
}

// CreateMenuInput メニュー作成の入力
type CreateMenuInput struct {
	MenuID      string  `json:"menu_id" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"required,min=0"`
	ImageURL    string  `json:"image_url"`
	IsSoldOut   bool    `json:"is_sold_out"`
	CategoryID  string  `json:"category_id" validate:"required"`
}

// UpdateMenuInput メニュー更新の入力
type UpdateMenuInput struct {
	MenuID      string  `json:"menu_id" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"required,min=0"`
	ImageURL    string  `json:"image_url"`
	IsSoldOut   bool    `json:"is_sold_out"`
	CategoryID  string  `json:"category_id" validate:"required"`
}

// DeleteMenuInput メニュー削除の入力
type DeleteMenuInput struct {
	MenuID string `json:"menu_id" validate:"required"`
}
