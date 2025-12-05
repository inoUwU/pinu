package order

import (
	"context"

	"github.com/google/uuid"
	"github.com/samber/do"
	"github.com/uptrace/bun"
	"inoUwU/pinu/app/domain/order"
	"inoUwU/pinu/app/infrastructure/ctx"
	"inoUwU/pinu/app/infrastructure/models"
)

func NewOrderRepository(i *do.Injector) (order.OrderRepository, error) {
	db := do.MustInvokeNamed[*bun.DB](i, "db")
	return &OrderRepositoryImpl{
		db: db,
	}, nil
}

// OrderRepositoryImpl 注文リポジトリの実装（アダプター）
type OrderRepositoryImpl struct {
	db *bun.DB
}

// CreateOrderGroup オーダーグループを作成します
func (r *OrderRepositoryImpl) CreateOrderGroup(ctx context.Context, orderGroup *order.OrderGroup) error {
	modelOrderGroup := &models.OrderGroupModel{
		OrdersID:       orderGroup.OrdersID.String(),
		TableSessionID: orderGroup.TableSessionID.String(),
		CreatedAt:      orderGroup.CreatedAt,
	}

	var inserter *bun.InsertQuery
	// contextからトランザクションオブジェクトを取得する
	if tx, ok := ctx.Value(ctxkey.TxCtxKey).(bun.Tx); ok {
		// トランザクションオブジェクトが存在する場合はそれを使用
		inserter = tx.NewInsert()
	} else {
		// トランザクションオブジェクトが存在しない場合はDBオブジェクトを使用
		inserter = r.db.NewInsert()
	}

	if _, err := inserter.Model(modelOrderGroup).Exec(ctx); err != nil {
		return err
	}
	return nil
}

// AddOrderItemOption implements order.OrderRepository.
func (o *OrderRepositoryImpl) AddOrderItemOption(ctx context.Context, orderItemOption *order.OrderItemOption) error {
	panic("unimplemented")
}

// CreateOrderItem implements order.OrderRepository.
func (o *OrderRepositoryImpl) CreateOrderItem(ctx context.Context, orderItem *order.OrderItem) error {
	panic("unimplemented")
}

// GetOrderGroupByID implements order.OrderRepository.
func (o *OrderRepositoryImpl) GetOrderGroupByID(ctx context.Context, id uuid.UUID) (*order.OrderGroup, error) {
	panic("unimplemented")
}

// GetOrderGroupsByTableSession implements order.OrderRepository.
func (o *OrderRepositoryImpl) GetOrderGroupsByTableSession(ctx context.Context, tableSessionID uuid.UUID) ([]*order.OrderGroup, error) {
	panic("unimplemented")
}

// GetOrderItemByID implements order.OrderRepository.
func (o *OrderRepositoryImpl) GetOrderItemByID(ctx context.Context, id order.OrderItemID) (*order.OrderItem, error) {
	panic("unimplemented")
}

// GetOrderItemOptions implements order.OrderRepository.
func (o *OrderRepositoryImpl) GetOrderItemOptions(ctx context.Context, orderItemID order.OrderItemID) ([]*order.OrderItemOption, error) {
	panic("unimplemented")
}

// GetOrderItemsByOrderGroup implements order.OrderRepository.
func (o *OrderRepositoryImpl) GetOrderItemsByOrderGroup(ctx context.Context, ordersID uuid.UUID) ([]*order.OrderItem, error) {
	panic("unimplemented")
}

// RemoveOrderItemOption implements order.OrderRepository.
func (o *OrderRepositoryImpl) RemoveOrderItemOption(ctx context.Context, orderItemID order.OrderItemID, menuOptionID string) error {
	panic("unimplemented")
}

// UpdateOrderItemStatus implements order.OrderRepository.
func (o *OrderRepositoryImpl) UpdateOrderItemStatus(ctx context.Context, id order.OrderItemID, status order.OrderStatus) error {
	panic("unimplemented")
}
