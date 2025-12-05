package order

import (
	"time"

	"github.com/google/uuid"
)

type OrderItemID string
type OrderStatus string

const (
	StatusPending   OrderStatus = "pending"
	StatusPreparing OrderStatus = "preparing"
	StatusServed    OrderStatus = "served"
	StatusCancelled OrderStatus = "cancelled"
)

// OrderGroup 注文グループエンティティ
type OrderGroup struct {
	OrdersID       uuid.UUID `json:"orders_id"`
	TableSessionID uuid.UUID `json:"table_session_id"`
	CreatedAt      time.Time `json:"created_at"`
}

// OrderItem 注文アイテムエンティティ
type OrderItem struct {
	OrderItemID  OrderItemID `json:"order_item_id"`
	OrdersID     uuid.UUID   `json:"orders_id"`
	MenuID       string      `json:"menu_id"`
	Quantity     int         `json:"quantity"`
	PriceAtOrder float64     `json:"price_at_order"`
	Status       OrderStatus `json:"status"`
	CreatedAt    time.Time   `json:"created_at"`
}

// OrderItemOption 注文アイテムオプションエンティティ
type OrderItemOption struct {
	OrderItemID  OrderItemID `json:"order_item_id"`
	MenuOptionID string      `json:"menu_option_id"`
}
