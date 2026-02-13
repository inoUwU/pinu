package input

// CreateOrderItemInput 注文アイテム入力
// MenuOptionIDs は任意。
type CreateOrderItemInput struct {
	MenuID        string   `json:"menu_id"`
	Quantity      int      `json:"quantity"`
	MenuOptionIDs []string `json:"menu_option_ids,omitempty"`
}

// CreateOrderInput 注文作成入力
type CreateOrderInput struct {
	TableSessionID string                 `json:"table_session_id"`
	Items          []CreateOrderItemInput `json:"items"`
}

// GetOrdersInput 注文一覧取得入力
type GetOrdersInput struct {
	TableSessionID string `json:"table_session_id"`
}
