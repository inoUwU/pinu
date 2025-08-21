package handlers

import (
	"github.com/gofiber/fiber/v2"

	"inoUwU/pinu/app/domain/table"
)

type TableHandler struct {
	repo table.TableRepository
}

// NewTableHandler テーブルハンドラーの新規作成
func NewTableHandler(repo table.TableRepository) *TableHandler {
	return &TableHandler{repo: repo}
}

// CreateTable テーブル作成
func (h *TableHandler) CreateTable(c *fiber.Ctx) error {
	var req struct {
		TableID string `json:"table_id"`
		Status  string `json:"status,omitempty"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	status := table.StatusAvailable
	if req.Status != "" {
		status = table.TableStatus(req.Status)
	}

	tbl := &table.Table{
		TableID: table.TableID(req.TableID),
		Status:  status,
	}

	if err := h.repo.Create(c.Context(), tbl); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create table",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(tbl)
}

// GetTable テーブル取得
func (h *TableHandler) GetTable(c *fiber.Ctx) error {
	id := table.TableID(c.Params("id"))

	tbl, err := h.repo.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get table",
		})
	}

	if tbl == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Table not found",
		})
	}

	return c.JSON(tbl)
}

// GetAllTables 全テーブル取得
func (h *TableHandler) GetAllTables(c *fiber.Ctx) error {
	tables, err := h.repo.GetAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get tables",
		})
	}

	return c.JSON(tables)
}

// GetTablesByStatus ステータス別テーブル取得
func (h *TableHandler) GetTablesByStatus(c *fiber.Ctx) error {
	status := table.TableStatus(c.Params("status"))

	tables, err := h.repo.GetByStatus(c.Context(), status)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get tables by status",
		})
	}

	return c.JSON(tables)
}

// UpdateTableStatus テーブルステータス更新
func (h *TableHandler) UpdateTableStatus(c *fiber.Ctx) error {
	id := table.TableID(c.Params("id"))

	var req struct {
		Status string `json:"status"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.repo.UpdateStatus(c.Context(), id, table.TableStatus(req.Status)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update table status",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Table status updated successfully",
	})
}

// DeleteTable テーブル削除
func (h *TableHandler) DeleteTable(c *fiber.Ctx) error {
	id := table.TableID(c.Params("id"))

	if err := h.repo.Delete(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete table",
		})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}
