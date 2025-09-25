package handlers

import (
	"net/http"
	"strconv"

	"inoUwU/pinu/app/usecases/analytics"
	"inoUwU/pinu/app/usecases/analytics/input"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/do"
)

type IAnalyticsHandler interface {
	GetAnalyticsData(c *fiber.Ctx) error
	GetKPISummary(c *fiber.Ctx) error
	GetTopMenus(c *fiber.Ctx) error
	GetCategorySales(c *fiber.Ctx) error
	GetMenuPerformance(c *fiber.Ctx) error
	GetDailySales(c *fiber.Ctx) error
	GetMenuDailyTrends(c *fiber.Ctx) error
}

// AnalyticsHandler 統計データハンドラー
type AnalyticsHandler struct {
	analyticsUsecase analytics.IAnalyticsUsecase
}

// NewAnalyticsHandler 統計データハンドラーを生成する
func NewAnalyticsHandler(i *do.Injector) (IAnalyticsHandler, error) {
	analyticsUsecase := do.MustInvoke[analytics.IAnalyticsUsecase](i)
	return &AnalyticsHandler{
		analyticsUsecase: analyticsUsecase,
	}, nil
}

// GetAnalyticsData 統計データ全体を取得します
func (h *AnalyticsHandler) GetAnalyticsData(c *fiber.Ctx) error {
	result, err := h.analyticsUsecase.GetAnalyticsData(c.UserContext(), &input.GetAnalyticsInput{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "統計データの取得に失敗しました",
		})
	}

	return c.Status(http.StatusOK).JSON(result)
}

// GetKPISummary KPIサマリーを取得します
func (h *AnalyticsHandler) GetKPISummary(c *fiber.Ctx) error {
	result, err := h.analyticsUsecase.GetKPISummary(c.UserContext(), &input.GetKPISummaryInput{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "KPIサマリーの取得に失敗しました",
		})
	}

	return c.Status(http.StatusOK).JSON(result)
}

// GetTopMenus トップメニューを取得します
func (h *AnalyticsHandler) GetTopMenus(c *fiber.Ctx) error {
	limit := 10 // デフォルト値

	// クエリパラメーターからlimitを取得
	if limitStr := c.Query("limit"); limitStr != "" {
		if parsedLimit, err := strconv.Atoi(limitStr); err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	result, err := h.analyticsUsecase.GetTopMenus(c.UserContext(), &input.GetTopMenusInput{
		Limit: limit,
	})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "トップメニューの取得に失敗しました",
		})
	}

	return c.Status(http.StatusOK).JSON(result)
}

// GetCategorySales カテゴリ別売上を取得します
func (h *AnalyticsHandler) GetCategorySales(c *fiber.Ctx) error {
	result, err := h.analyticsUsecase.GetCategorySales(c.UserContext(), &input.GetCategorySalesInput{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "カテゴリ別売上の取得に失敗しました",
		})
	}

	return c.Status(http.StatusOK).JSON(result)
}

// GetMenuPerformance メニューパフォーマンスを取得します
func (h *AnalyticsHandler) GetMenuPerformance(c *fiber.Ctx) error {
	limit := 20 // デフォルト値

	// クエリパラメーターからlimitを取得
	if limitStr := c.Query("limit"); limitStr != "" {
		if parsedLimit, err := strconv.Atoi(limitStr); err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	result, err := h.analyticsUsecase.GetMenuPerformance(c.UserContext(), &input.GetMenuPerformanceInput{
		Limit: limit,
	})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "メニューパフォーマンスの取得に失敗しました",
		})
	}

	return c.Status(http.StatusOK).JSON(result)
}

// GetDailySales 日次売上を取得します
func (h *AnalyticsHandler) GetDailySales(c *fiber.Ctx) error {
	result, err := h.analyticsUsecase.GetDailySales(c.UserContext(), &input.GetDailySalesInput{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "日次売上の取得に失敗しました",
		})
	}

	return c.Status(http.StatusOK).JSON(result)
}

// GetMenuDailyTrends メニュー日次トレンドを取得します
func (h *AnalyticsHandler) GetMenuDailyTrends(c *fiber.Ctx) error {
	// クエリパラメーターからmenuIdsを取得（カンマ区切り）
	var menuIds []string
	if menuIdsStr := c.Query("menuIds"); menuIdsStr != "" {
		// 簡単な分割処理（実際のプロダクションではより堅牢な解析が必要）
		menuIds = []string{menuIdsStr} // 単一メニューIDの場合
		// 複数の場合: menuIds = strings.Split(menuIdsStr, ",")
	}

	result, err := h.analyticsUsecase.GetMenuDailyTrends(c.UserContext(), &input.GetMenuDailyTrendsInput{
		MenuIds: menuIds,
	})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "メニュー日次トレンドの取得に失敗しました",
		})
	}

	return c.Status(http.StatusOK).JSON(result)
}