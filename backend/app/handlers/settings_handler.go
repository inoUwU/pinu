package handlers

import (
	"github.com/gofiber/fiber/v2"

	"inoUwU/pinu/app/domain/settings"
)

type SettingsHandler struct {
	repo settings.SettingsRepository
}

// NewSettingsHandler 設定ハンドラーの新規作成
func NewSettingsHandler(repo settings.SettingsRepository) *SettingsHandler {
	return &SettingsHandler{repo: repo}
}

// GetSetting 設定取得
func (h *SettingsHandler) GetSetting(c *fiber.Ctx) error {
	key := c.Params("key")

	setting, err := h.repo.Get(c.Context(), key)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get setting",
		})
	}

	if setting == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Setting not found",
		})
	}

	return c.JSON(setting)
}

// GetAllSettings 全設定取得
func (h *SettingsHandler) GetAllSettings(c *fiber.Ctx) error {
	settings, err := h.repo.GetAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get settings",
		})
	}

	return c.JSON(settings)
}

// SetSetting 設定保存
func (h *SettingsHandler) SetSetting(c *fiber.Ctx) error {
	var req struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.repo.Set(c.Context(), req.Key, req.Value); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to set setting",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Setting saved successfully",
	})
}

// DeleteSetting 設定削除
func (h *SettingsHandler) DeleteSetting(c *fiber.Ctx) error {
	key := c.Params("key")

	if err := h.repo.Delete(c.Context(), key); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete setting",
		})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}
