package order

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type OrderItemID string
type OrderStatus string
type OrderGroupStatus string

const (
	StatusPending   OrderStatus = "pending"
	StatusPreparing OrderStatus = "preparing"
	StatusServed    OrderStatus = "served"
	StatusCancelled OrderStatus = "cancelled"
)

const (
	GroupStatusOpen      OrderGroupStatus = "open"
	GroupStatusClosed    OrderGroupStatus = "closed"
	GroupStatusCancelled OrderGroupStatus = "cancelled"
)

var (
	ErrInvalidOrderStatusTransition = errors.New("invalid order status transition")
	ErrInvalidOrderQuantity         = errors.New("invalid order quantity")
	ErrInvalidOrderMenuID           = errors.New("invalid order menu id")
	ErrInvalidOrderPrice            = errors.New("invalid order price")
)

// OrderGroup 注文グループエンティティ
type OrderGroup struct {
	OrdersID       uuid.UUID        `json:"orders_id"`
	TableSessionID uuid.UUID        `json:"table_session_id"`
	Status         OrderGroupStatus `json:"status"`
	CreatedAt      time.Time        `json:"created_at"`
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

func NewOrderItem(ordersID uuid.UUID, menuID string, quantity int, priceAtOrder float64, createdAt time.Time) (*OrderItem, error) {
	if menuID == "" {
		return nil, ErrInvalidOrderMenuID
	}
	if quantity <= 0 {
		return nil, ErrInvalidOrderQuantity
	}
	if priceAtOrder < 0 {
		return nil, ErrInvalidOrderPrice
	}

	return &OrderItem{
		OrderItemID:  OrderItemID(uuid.NewString()),
		OrdersID:     ordersID,
		MenuID:       menuID,
		Quantity:     quantity,
		PriceAtOrder: priceAtOrder,
		Status:       StatusPending,
		CreatedAt:    createdAt,
	}, nil
}

func (o *OrderItem) CanTransitionTo(nextStatus OrderStatus) bool {
	if o == nil {
		return false
	}

	if o.Status == nextStatus {
		return true
	}

	switch o.Status {
	case StatusPending:
		return nextStatus == StatusPreparing || nextStatus == StatusCancelled
	case StatusPreparing:
		return nextStatus == StatusServed || nextStatus == StatusCancelled
	case StatusServed, StatusCancelled:
		return false
	default:
		return false
	}
}

func (o *OrderItem) UpdateStatus(nextStatus OrderStatus) error {
	if !o.CanTransitionTo(nextStatus) {
		return ErrInvalidOrderStatusTransition
	}

	o.Status = nextStatus
	return nil
}

// OrderItemOption 注文アイテムオプションエンティティ
type OrderItemOption struct {
	OrderItemID  OrderItemID `json:"order_item_id"`
	MenuOptionID string      `json:"menu_option_id"`
}
