package output

import "time"

// OrderItemOptionOutput 注文アイテムオプション出力
type OrderItemOptionOutput struct {
	MenuOptionID string `json:"menu_option_id"`
}

// OrderItemOutput 注文アイテム出力
type OrderItemOutput struct {
	OrderItemID  string                  `json:"order_item_id"`
	MenuID       string                  `json:"menu_id"`
	Quantity     int                     `json:"quantity"`
	PriceAtOrder float64                 `json:"price_at_order"`
	Status       string                  `json:"status"`
	CreatedAt    time.Time               `json:"created_at"`
	Options      []OrderItemOptionOutput `json:"options"`
}

// OrderGroupOutput 注文グループ出力
type OrderGroupOutput struct {
	OrdersID       string            `json:"orders_id"`
	TableSessionID string            `json:"table_session_id"`
	CreatedAt      time.Time         `json:"created_at"`
	Items          []OrderItemOutput `json:"items"`
}

// CreateOrderOutput 注文作成結果
type CreateOrderOutput struct {
	OrderGroup OrderGroupOutput `json:"order_group"`
}

// GetOrdersOutput 注文一覧取得結果
type GetOrdersOutput struct {
	OrderGroups []OrderGroupOutput `json:"order_groups"`
}
