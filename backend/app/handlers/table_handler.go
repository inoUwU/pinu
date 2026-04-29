package handlers

import (
	"inoUwU/pinu/app/domain/table"
	tableUsecase "inoUwU/pinu/app/usecases/table"
	"inoUwU/pinu/app/usecases/table/input"

	"github.com/gofiber/fiber/v3"
	"github.com/samber/do"
)

// TableHandler テーブルハンドラー
type TableHandler struct {
	tableUsecase tableUsecase.TableService
}

// NewTableHandler 新しいテーブルハンドラーを生成
func NewTableHandler(i *do.Injector) (*TableHandler, error) {
	tableUsecase := do.MustInvoke[tableUsecase.TableService](i)

	return &TableHandler{
		tableUsecase: tableUsecase,
	}, nil
}

// CreateTable テーブル作成
func (h *TableHandler) CreateTable(c fiber.Ctx) error {
	var req struct {
		TableID string `json:"table_id"`
		Status  string `json:"status,omitempty"`
	}

	if err := c.Bind().Body(&req); err != nil {
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

	output, err := h.tableUsecase.CreateTable(c.RequestCtx(), createInput)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create table",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(output)
}

// GetTable テーブル取得
func (h *TableHandler) GetTable(c fiber.Ctx) error {
	id := table.TableID(c.Params("id"))

	getInput := &input.GetTableByIDInput{
		TableID: id,
	}

	output, err := h.tableUsecase.GetTableByID(c.RequestCtx(), getInput)
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
func (h *TableHandler) GetAllTables(c fiber.Ctx) error {
	getInput := &input.GetTablesInput{}

	output, err := h.tableUsecase.GetAllTables(c.RequestCtx(), getInput)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get tables",
		})
	}

	return c.JSON(output)
}

// GetTablesByStatus ステータス別テーブル取得
func (h *TableHandler) GetTablesByStatus(c fiber.Ctx) error {
	status := table.TableStatus(c.Params("status"))

	getInput := &input.GetTablesByStatusInput{
		Status: status,
	}

	output, err := h.tableUsecase.GetTablesByStatus(c.RequestCtx(), getInput)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get tables by status",
		})
	}

	return c.JSON(output)
}

// UpdateTableStatus テーブルステータス更新
func (h *TableHandler) UpdateTableStatus(c fiber.Ctx) error {
	id := table.TableID(c.Params("id"))

	var req struct {
		Status string `json:"status"`
	}

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	updateInput := &input.UpdateTableStatusInput{
		TableID: id,
		Status:  table.TableStatus(req.Status),
	}

	output, err := h.tableUsecase.UpdateTableStatus(c.RequestCtx(), updateInput)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update table status",
		})
	}

	return c.Status(fiber.StatusOK).JSON(output)
}

// DeleteTable テーブル削除
func (h *TableHandler) DeleteTable(c fiber.Ctx) error {
	id := table.TableID(c.Params("id"))

	deleteInput := &input.DeleteTableInput{
		TableID: id,
	}

	output, err := h.tableUsecase.DeleteTable(c.RequestCtx(), deleteInput)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete table",
		})
	}

	return c.Status(fiber.StatusOK).JSON(output)
}

// Checkout テーブルの会計処理
func (h *TableHandler) Checkout(c fiber.Ctx) error {
	id := table.TableID(c.Params("id"))

	checkoutInput := &input.CheckoutTableInput{
		TableID: id,
	}

	result, err := h.tableUsecase.CheckoutTable(c.RequestCtx(), checkoutInput)
	if err != nil {
		switch err.Error() {
		case "table not found":
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Table not found",
			})
		case "table is not occupied":
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Table is not occupied",
			})
		case "table has no active session":
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Table has no active session",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   "Failed to checkout table",
				"message": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
