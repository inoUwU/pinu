package handlers

import (
	settingsUsecase "inoUwU/pinu/app/usecases/settings"
	"inoUwU/pinu/app/usecases/settings/input"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/do"
)

// SettingsHandler 設定ハンドラー
type SettingsHandler struct {
	settingsUsecase settingsUsecase.SettingsService
}

// NewSettingsHandler 新しい設定ハンドラーを生成
func NewSettingsHandler(i *do.Injector) (*SettingsHandler, error) {
	settingsUsecase := do.MustInvoke[settingsUsecase.SettingsService](i)

	return &SettingsHandler{
		settingsUsecase: settingsUsecase,
	}, nil
}

// GetSetting 設定取得
func (h *SettingsHandler) GetSetting(c *fiber.Ctx) error {
	key := c.Params("key")

	getInput := &input.GetSettingByKeyInput{
		Key: key,
	}

	output, err := h.settingsUsecase.GetSettingByKey(c.Context(), getInput)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get setting",
		})
	}

	if !output.Found {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Setting not found",
		})
	}

	return c.JSON(output)
}

// GetAllSettings 全設定取得
func (h *SettingsHandler) GetAllSettings(c *fiber.Ctx) error {
	getInput := &input.GetSettingsInput{}

	output, err := h.settingsUsecase.GetAllSettings(c.Context(), getInput)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get settings",
		})
	}

	return c.JSON(output)
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

	setInput := &input.SetSettingInput{
		Key:   req.Key,
		Value: req.Value,
	}

	output, err := h.settingsUsecase.SetSetting(c.Context(), setInput)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to set setting",
		})
	}

	return c.Status(fiber.StatusOK).JSON(output)
}

// DeleteSetting 設定削除
func (h *SettingsHandler) DeleteSetting(c *fiber.Ctx) error {
	key := c.Params("key")

	deleteInput := &input.DeleteSettingInput{
		Key: key,
	}

	output, err := h.settingsUsecase.DeleteSetting(c.Context(), deleteInput)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete setting",
		})
	}

	return c.Status(fiber.StatusOK).JSON(output)
}
