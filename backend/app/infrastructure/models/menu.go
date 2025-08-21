package models

import (
	"github.com/uptrace/bun"
)

// MenuModel メニューのインフラストラクチャーモデル
type MenuModel struct {
	bun.BaseModel `bun:"table:menus"`

	MenuID      string  `bun:"menu_id,pk"`
	Name        string  `bun:"name,notnull"`
	Description string  `bun:"description"`
	Price       float64 `bun:"price,notnull"`
	ImageURL    string  `bun:"image_url"`
	IsSoldOut   bool    `bun:"is_sold_out,notnull,default:false"`
	CategoryID  string  `bun:"category_id,notnull"`
}
