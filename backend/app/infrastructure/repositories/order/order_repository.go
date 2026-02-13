package order

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/samber/do"
	"github.com/uptrace/bun"
	"inoUwU/pinu/app/domain/order"
	"inoUwU/pinu/app/infrastructure/ctx"
	"inoUwU/pinu/app/infrastructure/models"
)

func NewOrderRepository(i *do.Injector) (order.OrderStore, error) {
	db := do.MustInvokeNamed[*bun.DB](i, "db")
	return &OrderRepositoryImpl{
		db: db,
	}, nil
}

// OrderRepositoryImpl 注文リポジトリの実装（アダプター）
type OrderRepositoryImpl struct {
	db *bun.DB
}

type orderItemOptionModel struct {
	bun.BaseModel `bun:"table:order_item_options"`

	OrderItemID  string `bun:"order_item_id,pk"`
	MenuOptionID string `bun:"menu_option_id,pk"`
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

// AddOrderItemOption implements order.OrderStore.
func (o *OrderRepositoryImpl) AddOrderItemOption(ctx context.Context, orderItemOption *order.OrderItemOption) error {
	model := &orderItemOptionModel{
		OrderItemID:  string(orderItemOption.OrderItemID),
		MenuOptionID: orderItemOption.MenuOptionID,
	}

	var inserter *bun.InsertQuery
	if tx, ok := ctx.Value(ctxkey.TxCtxKey).(bun.Tx); ok {
		inserter = tx.NewInsert()
	} else {
		inserter = o.db.NewInsert()
	}

	if _, err := inserter.Model(model).Exec(ctx); err != nil {
		return err
	}

	return nil
}

// CreateOrderItem implements order.OrderStore.
func (o *OrderRepositoryImpl) CreateOrderItem(ctx context.Context, orderItem *order.OrderItem) error {
	model := &models.OrderItemModel{
		OrderItemID:  string(orderItem.OrderItemID),
		OrdersID:     orderItem.OrdersID.String(),
		MenuID:       orderItem.MenuID,
		Quantity:     orderItem.Quantity,
		PriceAtOrder: orderItem.PriceAtOrder,
		Status:       string(orderItem.Status),
		CreatedAt:    orderItem.CreatedAt,
	}

	var inserter *bun.InsertQuery
	if tx, ok := ctx.Value(ctxkey.TxCtxKey).(bun.Tx); ok {
		inserter = tx.NewInsert()
	} else {
		inserter = o.db.NewInsert()
	}

	if _, err := inserter.Model(model).Exec(ctx); err != nil {
		return err
	}

	return nil
}

// GetOrderGroupByID implements order.OrderStore.
func (o *OrderRepositoryImpl) GetOrderGroupByID(ctx context.Context, id uuid.UUID) (*order.OrderGroup, error) {
	model := &models.OrderGroupModel{}

	var selector *bun.SelectQuery
	if tx, ok := ctx.Value(ctxkey.TxCtxKey).(bun.Tx); ok {
		selector = tx.NewSelect()
	} else {
		selector = o.db.NewSelect()
	}

	err := selector.Model(model).Where("orders_id = ?", id.String()).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	tableSessionID, err := uuid.Parse(model.TableSessionID)
	if err != nil {
		return nil, err
	}

	ordersID, err := uuid.Parse(model.OrdersID)
	if err != nil {
		return nil, err
	}

	return &order.OrderGroup{
		OrdersID:       ordersID,
		TableSessionID: tableSessionID,
		CreatedAt:      model.CreatedAt,
	}, nil
}

// GetOrderGroupsByTableSession implements order.OrderStore.
func (o *OrderRepositoryImpl) GetOrderGroupsByTableSession(ctx context.Context, tableSessionID uuid.UUID) ([]*order.OrderGroup, error) {
	modelsGroup := make([]models.OrderGroupModel, 0)

	var selector *bun.SelectQuery
	if tx, ok := ctx.Value(ctxkey.TxCtxKey).(bun.Tx); ok {
		selector = tx.NewSelect()
	} else {
		selector = o.db.NewSelect()
	}

	if err := selector.Model(&modelsGroup).
		Where("table_session_id = ?", tableSessionID.String()).
		Order("created_at DESC").
		Scan(ctx); err != nil {
		return nil, err
	}

	groups := make([]*order.OrderGroup, 0, len(modelsGroup))
	for _, model := range modelsGroup {
		ordersID, err := uuid.Parse(model.OrdersID)
		if err != nil {
			return nil, err
		}

		parsedTableSessionID, err := uuid.Parse(model.TableSessionID)
		if err != nil {
			return nil, err
		}

		groups = append(groups, &order.OrderGroup{
			OrdersID:       ordersID,
			TableSessionID: parsedTableSessionID,
			CreatedAt:      model.CreatedAt,
		})
	}

	return groups, nil
}

// GetOrderItemByID implements order.OrderStore.
func (o *OrderRepositoryImpl) GetOrderItemByID(ctx context.Context, id order.OrderItemID) (*order.OrderItem, error) {
	model := &models.OrderItemModel{}

	var selector *bun.SelectQuery
	if tx, ok := ctx.Value(ctxkey.TxCtxKey).(bun.Tx); ok {
		selector = tx.NewSelect()
	} else {
		selector = o.db.NewSelect()
	}

	err := selector.Model(model).Where("order_item_id = ?", string(id)).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	ordersID, err := uuid.Parse(model.OrdersID)
	if err != nil {
		return nil, err
	}

	return &order.OrderItem{
		OrderItemID:  order.OrderItemID(model.OrderItemID),
		OrdersID:     ordersID,
		MenuID:       model.MenuID,
		Quantity:     model.Quantity,
		PriceAtOrder: model.PriceAtOrder,
		Status:       order.OrderStatus(model.Status),
		CreatedAt:    model.CreatedAt,
	}, nil
}

// GetOrderItemOptions implements order.OrderStore.
func (o *OrderRepositoryImpl) GetOrderItemOptions(ctx context.Context, orderItemID order.OrderItemID) ([]*order.OrderItemOption, error) {
	optionModels := make([]orderItemOptionModel, 0)

	var selector *bun.SelectQuery
	if tx, ok := ctx.Value(ctxkey.TxCtxKey).(bun.Tx); ok {
		selector = tx.NewSelect()
	} else {
		selector = o.db.NewSelect()
	}

	if err := selector.Model(&optionModels).Where("order_item_id = ?", string(orderItemID)).Scan(ctx); err != nil {
		return nil, err
	}

	options := make([]*order.OrderItemOption, 0, len(optionModels))
	for _, model := range optionModels {
		options = append(options, &order.OrderItemOption{
			OrderItemID:  order.OrderItemID(model.OrderItemID),
			MenuOptionID: model.MenuOptionID,
		})
	}

	return options, nil
}

// GetOrderItemsByOrderGroup implements order.OrderStore.
func (o *OrderRepositoryImpl) GetOrderItemsByOrderGroup(ctx context.Context, ordersID uuid.UUID) ([]*order.OrderItem, error) {
	itemModels := make([]models.OrderItemModel, 0)

	var selector *bun.SelectQuery
	if tx, ok := ctx.Value(ctxkey.TxCtxKey).(bun.Tx); ok {
		selector = tx.NewSelect()
	} else {
		selector = o.db.NewSelect()
	}

	if err := selector.Model(&itemModels).
		Where("orders_id = ?", ordersID.String()).
		Order("created_at ASC").
		Scan(ctx); err != nil {
		return nil, err
	}

	items := make([]*order.OrderItem, 0, len(itemModels))
	for _, model := range itemModels {
		parsedOrdersID, err := uuid.Parse(model.OrdersID)
		if err != nil {
			return nil, err
		}

		items = append(items, &order.OrderItem{
			OrderItemID:  order.OrderItemID(model.OrderItemID),
			OrdersID:     parsedOrdersID,
			MenuID:       model.MenuID,
			Quantity:     model.Quantity,
			PriceAtOrder: model.PriceAtOrder,
			Status:       order.OrderStatus(model.Status),
			CreatedAt:    model.CreatedAt,
		})
	}

	return items, nil
}

// RemoveOrderItemOption implements order.OrderStore.
func (o *OrderRepositoryImpl) RemoveOrderItemOption(ctx context.Context, orderItemID order.OrderItemID, menuOptionID string) error {
	var deleter *bun.DeleteQuery
	if tx, ok := ctx.Value(ctxkey.TxCtxKey).(bun.Tx); ok {
		deleter = tx.NewDelete()
	} else {
		deleter = o.db.NewDelete()
	}

	if _, err := deleter.Model((*orderItemOptionModel)(nil)).
		Where("order_item_id = ?", string(orderItemID)).
		Where("menu_option_id = ?", menuOptionID).
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

// UpdateOrderItemStatus implements order.OrderStore.
func (o *OrderRepositoryImpl) UpdateOrderItemStatus(ctx context.Context, id order.OrderItemID, status order.OrderStatus) error {
	var updater *bun.UpdateQuery
	if tx, ok := ctx.Value(ctxkey.TxCtxKey).(bun.Tx); ok {
		updater = tx.NewUpdate()
	} else {
		updater = o.db.NewUpdate()
	}

	if _, err := updater.Model((*models.OrderItemModel)(nil)).
		Set("status = ?", string(status)).
		Set("created_at = ?", time.Now()).
		Where("order_item_id = ?", string(id)).
		Exec(ctx); err != nil {
		return err
	}

	return nil
}
