package handlers

import (
	"inoUwU/pinu/app/domain/table"
	tableUsecase "inoUwU/pinu/app/usecases/table"
	"inoUwU/pinu/app/usecases/table/input"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/do"
)

// TableHandler テーブルハンドラー
type TableHandler struct {
	tableUsecase tableUsecase.ITableUsecase
}

// NewTableHandler 新しいテーブルハンドラーを生成
func NewTableHandler(i *do.Injector) (*TableHandler, error) {
	tableUsecase := do.MustInvoke[tableUsecase.ITableUsecase](i)

	return &TableHandler{
		tableUsecase: tableUsecase,
	}, nil
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

	createInput := &input.CreateTableInput{
		TableID: table.TableID(req.TableID),
		Status:  status,
	}

	output, err := h.tableUsecase.CreateTable(c.Context(), createInput)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create table",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(output)
}

// GetTable テーブル取得
func (h *TableHandler) GetTable(c *fiber.Ctx) error {
	id := table.TableID(c.Params("id"))

	getInput := &input.GetTableByIDInput{
		TableID: id,
	}

	output, err := h.tableUsecase.GetTableByID(c.Context(), getInput)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get table",
		})
	}

	if output.Table == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Table not found",
		})
	}

	return c.JSON(output)
}

// GetAllTables 全テーブル取得
func (h *TableHandler) GetAllTables(c *fiber.Ctx) error {
	getInput := &input.GetTablesInput{}

	output, err := h.tableUsecase.GetAllTables(c.Context(), getInput)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get tables",
		})
	}

	return c.JSON(output)
}

// GetTablesByStatus ステータス別テーブル取得
func (h *TableHandler) GetTablesByStatus(c *fiber.Ctx) error {
	status := table.TableStatus(c.Params("status"))

	getInput := &input.GetTablesByStatusInput{
		Status: status,
	}

	output, err := h.tableUsecase.GetTablesByStatus(c.Context(), getInput)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get tables by status",
		})
	}

	return c.JSON(output)
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

	updateInput := &input.UpdateTableStatusInput{
		TableID: id,
		Status:  table.TableStatus(req.Status),
	}

	output, err := h.tableUsecase.UpdateTableStatus(c.Context(), updateInput)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update table status",
		})
	}

	return c.Status(fiber.StatusOK).JSON(output)
}

// DeleteTable テーブル削除
func (h *TableHandler) DeleteTable(c *fiber.Ctx) error {
	id := table.TableID(c.Params("id"))

	deleteInput := &input.DeleteTableInput{
		TableID: id,
	}

	output, err := h.tableUsecase.DeleteTable(c.Context(), deleteInput)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete table",
		})
	}

	return c.Status(fiber.StatusOK).JSON(output)
}
