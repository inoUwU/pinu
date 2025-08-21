package handlers

import (
	"github.com/gofiber/fiber/v2"

	"inoUwU/pinu/app/domain/menu_option"
)

type MenuOptionHandler struct {
	repo menu_option.MenuOptionRepository
}

// NewMenuOptionHandler メニューオプションハンドラーの新規作成
func NewMenuOptionHandler(repo menu_option.MenuOptionRepository) *MenuOptionHandler {
	return &MenuOptionHandler{repo: repo}
}

// CreateMenuOption メニューオプション作成
func (h *MenuOptionHandler) CreateMenuOption(c *fiber.Ctx) error {
	var req struct {
		MenuOptionID string  `json:"menu_option_id"`
		Name         string  `json:"name"`
		Price        float64 `json:"price"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	menuOption := &menu_option.MenuOption{
		MenuOptionID: menu_option.MenuOptionID(req.MenuOptionID),
		Name:         req.Name,
		Price:        req.Price,
	}

	if err := h.repo.Create(c.Context(), menuOption); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create menu option",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(menuOption)
}

// GetMenuOption メニューオプション取得
func (h *MenuOptionHandler) GetMenuOption(c *fiber.Ctx) error {
	id := menu_option.MenuOptionID(c.Params("id"))

	menuOption, err := h.repo.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get menu option",
		})
	}

	if menuOption == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Menu option not found",
		})
	}

	return c.JSON(menuOption)
}

// GetAllMenuOptions 全メニューオプション取得
func (h *MenuOptionHandler) GetAllMenuOptions(c *fiber.Ctx) error {
	menuOptions, err := h.repo.GetAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get menu options",
		})
	}

	return c.JSON(menuOptions)
}

// UpdateMenuOption メニューオプション更新
func (h *MenuOptionHandler) UpdateMenuOption(c *fiber.Ctx) error {
	id := menu_option.MenuOptionID(c.Params("id"))

	var req struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	menuOption := &menu_option.MenuOption{
		MenuOptionID: id,
		Name:         req.Name,
		Price:        req.Price,
	}

	if err := h.repo.Update(c.Context(), menuOption); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update menu option",
		})
	}

	return c.JSON(menuOption)
}

// DeleteMenuOption メニューオプション削除
func (h *MenuOptionHandler) DeleteMenuOption(c *fiber.Ctx) error {
	id := menu_option.MenuOptionID(c.Params("id"))

	if err := h.repo.Delete(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete menu option",
		})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}
