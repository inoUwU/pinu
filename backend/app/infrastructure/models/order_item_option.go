package models

import (
	"github.com/uptrace/bun"
)

// OrderItemOptionModel 注文アイテムオプションのインフラストラクチャーモデル
type OrderItemOptionModel struct {
	bun.BaseModel `bun:"table:order_item_options"`

	OrderItemID  string `bun:"order_item_id,pk"`
	MenuOptionID string `bun:"menu_option_id,pk"`
}
