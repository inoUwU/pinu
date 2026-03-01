package order

import (
	"context"
	"errors"
	"time"

	domainMenu "inoUwU/pinu/app/domain/menu"
	domainOrder "inoUwU/pinu/app/domain/order"
	"inoUwU/pinu/app/domain/port"
	domainSession "inoUwU/pinu/app/domain/session"
	domainTable "inoUwU/pinu/app/domain/table"
	"inoUwU/pinu/app/usecases/order/input"
	"inoUwU/pinu/app/usecases/order/output"

	"github.com/google/uuid"
	"github.com/samber/do"
)

var (
	ErrInvalidTableSessionID = errors.New("invalid table session id")
	ErrTableSessionNotFound  = errors.New("table session not found")
	ErrTableSessionRevoked   = errors.New("table session is revoked")
	ErrTableSessionExpired   = errors.New("table session is expired")
	ErrOrderItemsRequired    = errors.New("at least one order item is required")
	ErrInvalidOrderItem      = errors.New("order item is invalid")
	ErrMenuNotFound          = errors.New("menu not found")
	ErrMenuSoldOut           = errors.New("menu is sold out")
)

// OrderService 注文ユースケースのポート
type OrderService interface {
	CreateOrder(ctx context.Context, in *input.CreateOrderInput) (*output.CreateOrderOutput, error)
	GetOrdersByTableSession(ctx context.Context, in *input.GetOrdersInput) (*output.GetOrdersOutput, error)
}

type OrderUsecaseImpl struct {
	unitOfWork  port.UnitOfWork
	orderRepo   domainOrder.OrderRepository
	sessionRepo domainSession.SessionRepository
	tableRepo   domainTable.TableRepository
	menuRepo    domainMenu.MenuRepository
	logger      port.Logger
}

func NewOrderUsecase(i *do.Injector) (OrderService, error) {
	uow := do.MustInvokeNamed[port.UnitOfWork](i, "uow")
	orderRepo := do.MustInvoke[domainOrder.OrderRepository](i)
	sessionRepo := do.MustInvoke[domainSession.SessionRepository](i)
	tableRepo := do.MustInvoke[domainTable.TableRepository](i)
	menuRepo := do.MustInvoke[domainMenu.MenuRepository](i)
	logger := do.MustInvokeNamed[port.Logger](i, "logger")

	return &OrderUsecaseImpl{
		unitOfWork:  uow,
		orderRepo:   orderRepo,
		sessionRepo: sessionRepo,
		tableRepo:   tableRepo,
		menuRepo:    menuRepo,
		logger:      logger,
	}, nil
}

func (u *OrderUsecaseImpl) CreateOrder(ctx context.Context, in *input.CreateOrderInput) (*output.CreateOrderOutput, error) {
	tableSessionID, err := uuid.Parse(in.TableSessionID)
	if err != nil {
		return nil, ErrInvalidTableSessionID
	}

	tableSession, err := u.sessionRepo.GetTableSessionByID(ctx, tableSessionID)
	if err != nil {
		return nil, err
	}
	if tableSession == nil {
		return nil, ErrTableSessionNotFound
	}
	if tableSession.IsRevoked {
		return nil, ErrTableSessionRevoked
	}
	if tableSession.IsExpired(time.Now()) {
		return nil, ErrTableSessionExpired
	}
	if len(in.Items) == 0 {
		return nil, ErrOrderItemsRequired
	}

	var targetGroup *domainOrder.OrderGroup

	err = u.unitOfWork.Run(ctx, func(txCtx context.Context) error {
		groups, err := u.orderRepo.GetOrderGroupsByTableSession(txCtx, tableSessionID)
		if err != nil {
			return err
		}

		if len(groups) > 0 {
			targetGroup = groups[0]
		} else {
			targetGroup = &domainOrder.OrderGroup{
				OrdersID:       uuid.New(),
				TableSessionID: tableSessionID,
				Status:         domainOrder.GroupStatusOpen,
				CreatedAt:      time.Now(),
			}
			if err := u.orderRepo.CreateOrderGroup(txCtx, targetGroup); err != nil {
				return err
			}
		}

		for _, item := range in.Items {
			if item.MenuID == "" || item.Quantity <= 0 {
				return ErrInvalidOrderItem
			}

			menu, err := u.menuRepo.GetByID(txCtx, item.MenuID)
			if err != nil {
				return err
			}
			if menu == nil {
				return ErrMenuNotFound
			}
			if !menu.IsOrderable() {
				return ErrMenuSoldOut
			}

			orderItem, err := domainOrder.NewOrderItem(targetGroup.OrdersID, item.MenuID, item.Quantity, menu.Price, time.Now())
			if err != nil {
				return ErrInvalidOrderItem
			}

			if err := u.orderRepo.CreateOrderItem(txCtx, orderItem); err != nil {
				return err
			}

			for _, menuOptionID := range item.MenuOptionIDs {
				if menuOptionID == "" {
					continue
				}
				if err := u.orderRepo.AddOrderItemOption(txCtx, &domainOrder.OrderItemOption{
					OrderItemID:  orderItem.OrderItemID,
					MenuOptionID: menuOptionID,
				}); err != nil {
					return err
				}
			}
		}

		if err := u.sessionRepo.UpdateTableSessionLastUsed(txCtx, tableSessionID); err != nil {
			return err
		}

		if err := u.tableRepo.UpdateStatus(txCtx, domainTable.TableID(tableSession.TableID), domainTable.StatusOccupied); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	groupOutput, err := u.buildOrderGroupOutput(ctx, targetGroup)
	if err != nil {
		return nil, err
	}

	u.logger.Info("order created", "tableSessionID", in.TableSessionID, "ordersID", targetGroup.OrdersID.String())

	return &output.CreateOrderOutput{OrderGroup: *groupOutput}, nil
}

func (u *OrderUsecaseImpl) GetOrdersByTableSession(ctx context.Context, in *input.GetOrdersInput) (*output.GetOrdersOutput, error) {
	tableSessionID, err := uuid.Parse(in.TableSessionID)
	if err != nil {
		return nil, ErrInvalidTableSessionID
	}

	tableSession, err := u.sessionRepo.GetTableSessionByID(ctx, tableSessionID)
	if err != nil {
		return nil, err
	}
	if tableSession == nil {
		return nil, ErrTableSessionNotFound
	}

	groups, err := u.orderRepo.GetOrderGroupsByTableSession(ctx, tableSessionID)
	if err != nil {
		return nil, err
	}

	result := make([]output.OrderGroupOutput, 0, len(groups))
	for _, group := range groups {
		groupOutput, err := u.buildOrderGroupOutput(ctx, group)
		if err != nil {
			return nil, err
		}
		result = append(result, *groupOutput)
	}

	return &output.GetOrdersOutput{OrderGroups: result}, nil
}

func (u *OrderUsecaseImpl) buildOrderGroupOutput(ctx context.Context, group *domainOrder.OrderGroup) (*output.OrderGroupOutput, error) {
	items, err := u.orderRepo.GetOrderItemsByOrderGroup(ctx, group.OrdersID)
	if err != nil {
		return nil, err
	}

	itemOutputs := make([]output.OrderItemOutput, 0, len(items))
	for _, item := range items {
		options, err := u.orderRepo.GetOrderItemOptions(ctx, item.OrderItemID)
		if err != nil {
			return nil, err
		}

		optionOutputs := make([]output.OrderItemOptionOutput, 0, len(options))
		for _, option := range options {
			optionOutputs = append(optionOutputs, output.OrderItemOptionOutput{
				MenuOptionID: option.MenuOptionID,
			})
		}

		itemOutputs = append(itemOutputs, output.OrderItemOutput{
			OrderItemID:  string(item.OrderItemID),
			MenuID:       item.MenuID,
			Quantity:     item.Quantity,
			PriceAtOrder: item.PriceAtOrder,
			Status:       string(item.Status),
			CreatedAt:    item.CreatedAt,
			Options:      optionOutputs,
		})
	}

	return &output.OrderGroupOutput{
		OrdersID:       group.OrdersID.String(),
		TableSessionID: group.TableSessionID.String(),
		Status:         string(group.Status),
		CreatedAt:      group.CreatedAt,
		Items:          itemOutputs,
	}, nil
}
