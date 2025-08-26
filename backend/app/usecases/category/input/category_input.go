package input

// GetCategoriesInput カテゴリ一覧取得の入力
type GetCategoriesInput struct {
	// 将来的に検索条件やページング情報を追加
}

// GetCategoryByIDInput カテゴリ単一取得の入力
type GetCategoryByIDInput struct {
	CategoryID string `json:"category_id" validate:"required"`
}

// CreateCategoryInput カテゴリ作成の入力
type CreateCategoryInput struct {
	CategoryID   string `json:"category_id" validate:"required"`
	Name         string `json:"name" validate:"required"`
	DisplayOrder int    `json:"display_order"`
	ImageURL     string `json:"image_url"`
}

// UpdateCategoryInput カテゴリ更新の入力
type UpdateCategoryInput struct {
	CategoryID   string `json:"category_id" validate:"required"`
	Name         string `json:"name" validate:"required"`
	DisplayOrder int    `json:"display_order"`
	ImageURL     string `json:"image_url"`
}

// DeleteCategoryInput カテゴリ削除の入力
type DeleteCategoryInput struct {
	CategoryID string `json:"category_id" validate:"required"`
}
