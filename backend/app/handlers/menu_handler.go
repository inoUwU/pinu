package handlers

import (
	"github.com/gofiber/fiber/v2"

	"inoUwU/pinu/app/domain/menu"
)

type MenuHandler struct {
	repo menu.MenuRepository
}

// NewMenuHandler メニューハンドラーの新規作成
func NewMenuHandler(repo menu.MenuRepository) *MenuHandler {
	return &MenuHandler{repo: repo}
}

// CreateMenu メニュー作成
func (h *MenuHandler) CreateMenu(c *fiber.Ctx) error {
	var req struct {
		MenuID      string  `json:"menu_id"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		ImageURL    string  `json:"image_url"`
		IsSoldOut   bool    `json:"is_sold_out"`
		CategoryID  string  `json:"category_id"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	menuItem := &menu.Menu{
		MENU_ID:     req.MenuID,
		NAME:        req.Name,
		DESCRIPTION: req.Description,
		PRICE:       req.Price,
		IMAGE_URL:   req.ImageURL,
		IS_SOLD_OUT: req.IsSoldOut,
		CATEGORY_ID: req.CategoryID,
	}

	if err := h.repo.Create(c.Context(), menuItem); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create menu",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(menuItem)
}

// GetMenu メニュー取得
func (h *MenuHandler) GetMenu(c *fiber.Ctx) error {
	id := c.Params("id")

	menuItem, err := h.repo.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get menu",
		})
	}

	if menuItem == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Menu not found",
		})
	}

	return c.JSON(menuItem)
}

// GetAllMenus 全メニュー取得
func (h *MenuHandler) GetAllMenus(c *fiber.Ctx) error {
	menus, err := h.repo.GetAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get menus",
		})
	}

	return c.JSON(menus)
}

// GetMenusByCategory カテゴリ別メニュー取得
func (h *MenuHandler) GetMenusByCategory(c *fiber.Ctx) error {
	categoryID := c.Params("categoryId")

	menus, err := h.repo.GetByCategory(c.Context(), categoryID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get menus by category",
		})
	}

	return c.JSON(menus)
}

// UpdateMenu メニュー更新
func (h *MenuHandler) UpdateMenu(c *fiber.Ctx) error {
	id := c.Params("id")

	var req struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		ImageURL    string  `json:"image_url"`
		IsSoldOut   bool    `json:"is_sold_out"`
		CategoryID  string  `json:"category_id"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	menuItem := &menu.Menu{
		MENU_ID:     id,
		NAME:        req.Name,
		DESCRIPTION: req.Description,
		PRICE:       req.Price,
		IMAGE_URL:   req.ImageURL,
		IS_SOLD_OUT: req.IsSoldOut,
		CATEGORY_ID: req.CategoryID,
	}

	if err := h.repo.Update(c.Context(), menuItem); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update menu",
		})
	}

	return c.JSON(menuItem)
}

// DeleteMenu メニュー削除
func (h *MenuHandler) DeleteMenu(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.repo.Delete(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete menu",
		})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}
