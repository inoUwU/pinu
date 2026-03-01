package models

import (
	"github.com/uptrace/bun"
)

// MenuOptionModel メニューオプションのインフラストラクチャーモデル
type MenuOptionModel struct {
	bun.BaseModel `bun:"table:menu_options"`

	MenuOptionID string  `bun:"menu_option_id,pk"`
	Name         string  `bun:"name,notnull"`
	Price        float64 `bun:"price,notnull"`
}
