package order

import (
	"context"

	"github.com/google/uuid"
)

// OrderRepository 注文永続化のポート
type OrderRepository interface {
	// OrderGroup関連
	CreateOrderGroup(ctx context.Context, orderGroup *OrderGroup) error
	GetOrderGroupByID(ctx context.Context, id uuid.UUID) (*OrderGroup, error)
	GetOrderGroupsByTableSession(ctx context.Context, tableSessionID uuid.UUID) ([]*OrderGroup, error)
	// CloseOrderGroupsByTableSession テーブルセッションに紐づく全オーダーグループを closed にする
	CloseOrderGroupsByTableSession(ctx context.Context, tableSessionID uuid.UUID) error

	// OrderItem関連
	CreateOrderItem(ctx context.Context, orderItem *OrderItem) error
	GetOrderItemByID(ctx context.Context, id OrderItemID) (*OrderItem, error)
	GetOrderItemsByOrderGroup(ctx context.Context, ordersID uuid.UUID) ([]*OrderItem, error)
	UpdateOrderItemStatus(ctx context.Context, id OrderItemID, status OrderStatus) error

	// OrderItemOption関連
	AddOrderItemOption(ctx context.Context, orderItemOption *OrderItemOption) error
	GetOrderItemOptions(ctx context.Context, orderItemID OrderItemID) ([]*OrderItemOption, error)
	RemoveOrderItemOption(ctx context.Context, orderItemID OrderItemID, menuOptionID string) error
}
