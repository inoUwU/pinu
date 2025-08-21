package models

import (
	"github.com/uptrace/bun"
)

// CategoryModel カテゴリのインフラストラクチャーモデル
type CategoryModel struct {
	bun.BaseModel `bun:"table:categories"`

	CategoryID   string `bun:"category_id,pk"`
	Name         string `bun:"name,notnull"`
	DisplayOrder int    `bun:"display_order,notnull,default:0"`
	ImageURL     string `bun:"image_url"`
}
