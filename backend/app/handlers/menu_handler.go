package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/samber/do"

	menuUsecase "inoUwU/pinu/app/usecases/menu"
	"inoUwU/pinu/app/usecases/menu/input"
)

type MenuHandler struct {
	menuUsecase menuUsecase.MenuService
}

// NewMenuHandler メニューハンドラーの新規作成
func NewMenuHandler(i *do.Injector) (*MenuHandler, error) {
	menuUC := do.MustInvoke[menuUsecase.MenuService](i)
	return &MenuHandler{
		menuUsecase: menuUC,
	}, nil
}

// CreateMenu メニュー作成
func (h *MenuHandler) CreateMenu(c fiber.Ctx) error {
	req := new(input.CreateMenuInput)

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	result, err := h.menuUsecase.CreateMenu(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create menu",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"result":  result,
		"message": "Menu created successfully",
	})
}

// GetMenu メニュー取得
func (h *MenuHandler) GetMenu(c fiber.Ctx) error {
	id := c.Params("id")

	req := &input.GetMenuByIDInput{
		MenuID: id,
	}

	result, err := h.menuUsecase.GetMenuByID(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get menu",
		})
	}

	if result.Menu == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Menu not found",
		})
	}

	return c.Status(http.StatusOK).JSON(result)
}

// GetAllMenus 全メニュー取得
func (h *MenuHandler) GetAllMenus(c fiber.Ctx) error {
	req := &input.GetMenusInput{}

	result, err := h.menuUsecase.GetAllMenus(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get menus",
		})
	}

	return c.Status(http.StatusOK).JSON(result)
}

// GetMenusByCategory カテゴリ別メニュー取得
func (h *MenuHandler) GetMenusByCategory(c fiber.Ctx) error {
	categoryID := c.Params("categoryId")

	req := &input.GetMenusByCategoryInput{
		CategoryID: categoryID,
	}

	result, err := h.menuUsecase.GetMenusByCategory(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get menus by category",
		})
	}

	return c.Status(http.StatusOK).JSON(result)
}

// UpdateMenu メニュー更新
func (h *MenuHandler) UpdateMenu(c fiber.Ctx) error {
	id := c.Params("id")

	req := new(input.UpdateMenuInput)
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// URLパラメータからIDを設定
	req.MenuID = id

	result, err := h.menuUsecase.UpdateMenu(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update menu",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"result":  result,
		"message": "Menu updated successfully",
	})
}

// DeleteMenu メニュー削除
func (h *MenuHandler) DeleteMenu(c fiber.Ctx) error {
	id := c.Params("id")

	req := &input.DeleteMenuInput{
		MenuID: id,
	}

	result, err := h.menuUsecase.DeleteMenu(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete menu",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"result":  result,
		"message": "Menu deleted successfully",
	})
}
