package models

import (
	"time"

	"github.com/uptrace/bun"
)

// MenuOptionModel メニューオプションのインフラストラクチャーモデル
type MenuOptionModel struct {
	bun.BaseModel `bun:"table:menu_options"`

	MenuOptionID string  `bun:"menu_option_id,pk"`
	Name         string  `bun:"name,notnull"`
	Price        float64 `bun:"price,notnull"`
}

// TableModel テーブルのインフラストラクチャーモデル
type TableModel struct {
	bun.BaseModel `bun:"table:tables"`

	TableID         string    `bun:"table_id,pk"`
	Status          string    `bun:"status,notnull"`
	CurrentOrdersID *string   `bun:"current_table_session_id"`
	LastUpdated     time.Time `bun:"last_updated,notnull,default:current_timestamp"`
}

// OrderGroupModel 注文グループのインフラストラクチャーモデル
type OrderGroupModel struct {
	bun.BaseModel `bun:"table:order_groups"`

	OrdersID       string    `bun:"orders_id,pk"`
	TableSessionID string    `bun:"table_session_id,notnull"`
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

// SessionModel セッションのインフラストラクチャーモデル
type SessionModel struct {
	bun.BaseModel `bun:"table:sessions"`

	SessionID string    `bun:"session_id,pk"`
	UserID    string    `bun:"user_id,notnull"`
	CreatedAt time.Time `bun:"created_at,notnull,default:now()"`
	ExpiresAt time.Time `bun:"expires_at,notnull"`
	IPAddress *string   `bun:"ip_address"`
	UserAgent *string   `bun:"user_agent"`
}

// TableSessionModel テーブルセッションのインフラストラクチャーモデル
type TableSessionModel struct {
	bun.BaseModel `bun:"table:table_sessions"`

	TableSessionID string    `bun:"table_session_id,pk"`
	TableID        string    `bun:"table_id,notnull"`
	IsRevoked      bool      `bun:"is_revoked,notnull,default:false"`
	CreatedAt      time.Time `bun:"created_at,notnull,default:current_timestamp"`
	LastUsed       time.Time `bun:"last_used,notnull,default:current_timestamp"`
	ExpiresAt      time.Time `bun:"expires_at,notnull"`
}

// SettingsModel 設定のインフラストラクチャーモデル
type SettingsModel struct {
	bun.BaseModel `bun:"table:settings"`

	Key   string `bun:"key,pk"`
	Value string `bun:"value,notnull"`
}
