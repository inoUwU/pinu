package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/samber/do"

	"inoUwU/pinu/app/domain/menu_option"
	menuOptionUsecase "inoUwU/pinu/app/usecases/menu_option"
	"inoUwU/pinu/app/usecases/menu_option/input"
)

type MenuOptionHandler struct {
	menuOptionUsecase menuOptionUsecase.MenuOptionService
}

// NewMenuOptionHandler メニューオプションハンドラーの新規作成
func NewMenuOptionHandler(i *do.Injector) (*MenuOptionHandler, error) {
	menuOptionUC := do.MustInvoke[menuOptionUsecase.MenuOptionService](i)
	return &MenuOptionHandler{
		menuOptionUsecase: menuOptionUC,
	}, nil
}

// CreateMenuOption メニューオプション作成
func (h *MenuOptionHandler) CreateMenuOption(c fiber.Ctx) error {
	req := new(input.CreateMenuOptionInput)

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	result, err := h.menuOptionUsecase.CreateMenuOption(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create menu option",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"result":  result,
		"message": "Menu option created successfully",
	})
}

// GetMenuOption メニューオプション取得
func (h *MenuOptionHandler) GetMenuOption(c fiber.Ctx) error {
	id := c.Params("id")

	req := &input.GetMenuOptionByIDInput{
		MenuOptionID: menu_option.MenuOptionID(id),
	}

	result, err := h.menuOptionUsecase.GetMenuOptionByID(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get menu option",
		})
	}

	if result.MenuOption == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Menu option not found",
		})
	}

	return c.Status(http.StatusOK).JSON(result)
}

// GetAllMenuOptions 全メニューオプション取得
func (h *MenuOptionHandler) GetAllMenuOptions(c fiber.Ctx) error {
	req := &input.GetMenuOptionsInput{}

	result, err := h.menuOptionUsecase.GetAllMenuOptions(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get menu options",
		})
	}

	return c.Status(http.StatusOK).JSON(result)
}

// UpdateMenuOption メニューオプション更新
func (h *MenuOptionHandler) UpdateMenuOption(c fiber.Ctx) error {
	id := c.Params("id")

	req := new(input.UpdateMenuOptionInput)
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// URLパラメータからIDを設定
	req.MenuOptionID = menu_option.MenuOptionID(id)

	result, err := h.menuOptionUsecase.UpdateMenuOption(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update menu option",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"result":  result,
		"message": "Menu option updated successfully",
	})
}

// DeleteMenuOption メニューオプション削除
func (h *MenuOptionHandler) DeleteMenuOption(c fiber.Ctx) error {
	id := c.Params("id")

	req := &input.DeleteMenuOptionInput{
		MenuOptionID: menu_option.MenuOptionID(id),
	}

	result, err := h.menuOptionUsecase.DeleteMenuOption(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete menu option",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"result":  result,
		"message": "Menu option deleted successfully",
	})
}
