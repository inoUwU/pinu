package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samber/do"
)

type IOrderHandler interface {
	GetAllOrders(c *fiber.Ctx) error
	Order(c *fiber.Ctx) error
}

type OrderHandler struct {
	// ここに必要なサービスを追加
}

// NewOrderHandler 注文ハンドラーを生成する
func NewOrderHandler(i *do.Injector) (IOrderHandler, error) {
	return &OrderHandler{}, nil
}

// GetAllOrders すべての注文を取得します
func (h OrderHandler) GetAllOrders(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

// Order 注文を作成します
func (h OrderHandler) Order(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
