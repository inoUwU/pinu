package output

import (
	"time"

	"inoUwU/pinu/app/domain/category"
)

// GetCategoriesOutput カテゴリ一覧取得の出力
type GetCategoriesOutput struct {
	Categories []*category.Category `json:"categories"`
	Count      int                  `json:"count"`
}

// GetCategoryByIDOutput カテゴリ単一取得の出力
type GetCategoryByIDOutput struct {
	Category *category.Category `json:"category"`
}

// CreateCategoryOutput カテゴリ作成の出力
type CreateCategoryOutput struct {
	Category     *category.Category `json:"category"`
	Message      string             `json:"message"`
	CreatedAt    time.Time          `json:"created_at"`
	CategoryID   string             `json:"category_id"`
	Name         string             `json:"name"`
	DisplayOrder int                `json:"display_order"`
	ImageURL     string             `json:"image_url"`
}

// UpdateCategoryOutput カテゴリ更新の出力
type UpdateCategoryOutput struct {
	Category  *category.Category `json:"category"`
	Message   string             `json:"message"`
	UpdatedAt time.Time          `json:"updated_at"`
}

// DeleteCategoryOutput カテゴリ削除の出力
type DeleteCategoryOutput struct {
	CategoryID string    `json:"category_id"`
	Message    string    `json:"message"`
	DeletedAt  time.Time `json:"deleted_at"`
}
