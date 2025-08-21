package handlers

import (
	"github.com/gofiber/fiber/v2"

	"inoUwU/pinu/app/domain/category"
)

type CategoryHandler struct {
	repo category.CategoryRepository
}

// NewCategoryHandler カテゴリハンドラーの新規作成
func NewCategoryHandler(repo category.CategoryRepository) *CategoryHandler {
	return &CategoryHandler{repo: repo}
}

// CreateCategory カテゴリ作成
func (h *CategoryHandler) CreateCategory(c *fiber.Ctx) error {
	var req struct {
		CategoryID   string `json:"category_id"`
		Name         string `json:"name"`
		DisplayOrder int    `json:"display_order,omitempty"`
		ImageURL     string `json:"image_url,omitempty"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	cat := &category.Category{
		CategoryID:   req.CategoryID,
		Name:         req.Name,
		DisplayOrder: req.DisplayOrder,
		ImageURL:     req.ImageURL,
	}

	if err := h.repo.Create(c.Context(), cat); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create category",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(cat)
}

// GetCategory カテゴリ取得
func (h *CategoryHandler) GetCategory(c *fiber.Ctx) error {
	id := c.Params("id")

	cat, err := h.repo.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get category",
		})
	}

	if cat == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Category not found",
		})
	}

	return c.JSON(cat)
}

// GetAllCategories 全カテゴリ取得
func (h *CategoryHandler) GetAllCategories(c *fiber.Ctx) error {
	categories, err := h.repo.GetAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get categories",
		})
	}

	return c.JSON(categories)
}

// UpdateCategory カテゴリ更新
func (h *CategoryHandler) UpdateCategory(c *fiber.Ctx) error {
	id := c.Params("id")

	var req struct {
		Name         string `json:"name"`
		DisplayOrder int    `json:"display_order"`
		ImageURL     string `json:"image_url"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	cat := &category.Category{
		CategoryID:   id,
		Name:         req.Name,
		DisplayOrder: req.DisplayOrder,
		ImageURL:     req.ImageURL,
	}

	if err := h.repo.Update(c.Context(), cat); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update category",
		})
	}

	return c.JSON(cat)
}

// DeleteCategory カテゴリ削除
func (h *CategoryHandler) DeleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.repo.Delete(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete category",
		})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}
