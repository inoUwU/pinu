package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/samber/do"

	categoryUsecase "inoUwU/pinu/app/usecases/category"
	"inoUwU/pinu/app/usecases/category/input"
)

type CategoryHandler struct {
	categoryUsecase categoryUsecase.CategoryService
}

// NewCategoryHandler カテゴリハンドラーの新規作成
func NewCategoryHandler(i *do.Injector) (*CategoryHandler, error) {
	categoryUC := do.MustInvoke[categoryUsecase.CategoryService](i)
	return &CategoryHandler{
		categoryUsecase: categoryUC,
	}, nil
}

// CreateCategory カテゴリ作成
func (h *CategoryHandler) CreateCategory(c fiber.Ctx) error {
	req := new(input.CreateCategoryInput)

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	result, err := h.categoryUsecase.CreateCategory(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create category",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"result":  result,
		"message": "Category created successfully",
	})
}

// GetCategory カテゴリ取得
func (h *CategoryHandler) GetCategory(c fiber.Ctx) error {
	id := c.Params("id")

	req := &input.GetCategoryByIDInput{
		CategoryID: id,
	}

	result, err := h.categoryUsecase.GetCategoryByID(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get category",
		})
	}

	if result.Category == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Category not found",
		})
	}

	return c.Status(http.StatusOK).JSON(result)
}

// GetAllCategories 全カテゴリ取得
func (h *CategoryHandler) GetAllCategories(c fiber.Ctx) error {
	req := &input.GetCategoriesInput{}

	result, err := h.categoryUsecase.GetAllCategories(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get categories",
		})
	}

	return c.Status(http.StatusOK).JSON(result)
}

// UpdateCategory カテゴリ更新
func (h *CategoryHandler) UpdateCategory(c fiber.Ctx) error {
	id := c.Params("id")

	req := new(input.UpdateCategoryInput)
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// URLパラメータからIDを設定
	req.CategoryID = id

	result, err := h.categoryUsecase.UpdateCategory(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update category",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"result":  result,
		"message": "Category updated successfully",
	})
}

// DeleteCategory カテゴリ削除
func (h *CategoryHandler) DeleteCategory(c fiber.Ctx) error {
	id := c.Params("id")

	req := &input.DeleteCategoryInput{
		CategoryID: id,
	}

	result, err := h.categoryUsecase.DeleteCategory(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete category",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"result":  result,
		"message": "Category deleted successfully",
	})
}
