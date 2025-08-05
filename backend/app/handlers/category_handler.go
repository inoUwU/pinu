package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/do"
)

type ICategoryHandler interface {
	Handler
	GetAllCategorys(c *fiber.Ctx) error
}

type CategoryHandler struct {
}

// NewCategoryHandler カテゴリハンドラーを生成する
func NewCategoryHandler(i *do.Injector) (ICategoryHandler, error) {
	return &CategoryHandler{}, nil
}

func (h *CategoryHandler) Route(router fiber.Router) error {
	fmt.Println("Handling category request")
	category := router.Group("/category")
	category.Get("/", h.GetAllCategorys)
	return nil
}

func (h *CategoryHandler) GetAllCategorys(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "GetAllCategorys called",
	})
}
