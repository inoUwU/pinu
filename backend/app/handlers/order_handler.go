package handlers

import (
	"errors"
	orderUsecase "inoUwU/pinu/app/usecases/order"
	"inoUwU/pinu/app/usecases/order/input"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/samber/do"
)

type OrderHandler struct {
	orderService orderUsecase.OrderService
}

// NewOrderHandler 注文ハンドラーを生成する
func NewOrderHandler(i *do.Injector) (*OrderHandler, error) {
	service := do.MustInvoke[orderUsecase.OrderService](i)
	return &OrderHandler{orderService: service}, nil
}

// GetAllOrders すべての注文を取得します
func (h OrderHandler) GetAllOrders(c fiber.Ctx) error {
	tableSessionID := c.Query("table_session_id")
	if tableSessionID == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "table_session_id is required"})
	}

	result, err := h.orderService.GetOrdersByTableSession(c.Context(), &input.GetOrdersInput{TableSessionID: tableSessionID})
	if err != nil {
		if errors.Is(err, orderUsecase.ErrInvalidTableSessionID) {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		if errors.Is(err, orderUsecase.ErrTableSessionNotFound) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "failed to get orders", "message": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(result)
}

// Order 注文を作成します
func (h OrderHandler) Order(c fiber.Ctx) error {
	request := new(input.CreateOrderInput)
	if err := c.Bind().Body(request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse request"})
	}

	result, err := h.orderService.CreateOrder(c.Context(), request)
	if err != nil {
		switch {
		case errors.Is(err, orderUsecase.ErrInvalidTableSessionID),
			errors.Is(err, orderUsecase.ErrOrderItemsRequired),
			errors.Is(err, orderUsecase.ErrInvalidOrderItem):
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		case errors.Is(err, orderUsecase.ErrTableSessionNotFound),
			errors.Is(err, orderUsecase.ErrMenuNotFound):
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
		case errors.Is(err, orderUsecase.ErrTableSessionRevoked),
			errors.Is(err, orderUsecase.ErrTableSessionExpired),
			errors.Is(err, orderUsecase.ErrMenuSoldOut):
			return c.Status(http.StatusConflict).JSON(fiber.Map{"error": err.Error()})
		default:
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "failed to create order", "message": err.Error()})
		}
	}

	return c.Status(http.StatusCreated).JSON(result)
}
