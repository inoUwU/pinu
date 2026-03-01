package models

import (
	"time"

	"github.com/uptrace/bun"
)

// OrderGroupModel 注文グループのインフラストラクチャーモデル
type OrderGroupModel struct {
	bun.BaseModel `bun:"table:order_groups"`

	OrdersID       string    `bun:"orders_id,pk"`
	TableSessionID string    `bun:"table_session_id,notnull"`
	Status         string    `bun:"status,notnull,default:'open'"`
	CreatedAt      time.Time `bun:"created_at,notnull,default:current_timestamp"`
}

// OrderItemModel 注文アイテムのインフラストラクチャーモデル
type OrderItemModel struct {
	bun.BaseModel `bun:"table:order_items"`

	OrderItemID  string    `bun:"order_item_id,pk"`
	OrdersID     string    `bun:"orders_id,notnull"`
	MenuID       string    `bun:"menu_id,notnull"`
	Quantity     int       `bun:"quantity,notnull"`
	PriceAtOrder float64   `bun:"price_at_order,notnull"`
	Status       string    `bun:"status,notnull,default:'pending'"`
	CreatedAt    time.Time `bun:"created_at,notnull,default:current_timestamp"`
}
