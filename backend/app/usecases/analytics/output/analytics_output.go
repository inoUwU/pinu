package output

import (
	"inoUwU/pinu/app/domain/analytics"
)

// GetAnalyticsOutput 統計データ取得の出力
type GetAnalyticsOutput struct {
	Data analytics.AnalyticsData `json:"data"`
}

// GetKPISummaryOutput KPIサマリー取得の出力
type GetKPISummaryOutput struct {
	KPISummary analytics.KPISummary `json:"kpiSummary"`
}

// GetTopMenusOutput トップメニュー取得の出力
type GetTopMenusOutput struct {
	TopMenus []analytics.TopMenu `json:"topMenus"`
}

// GetCategorySalesOutput カテゴリ別売上取得の出力
type GetCategorySalesOutput struct {
	CategorySales []analytics.CategorySales `json:"categorySales"`
}

// GetMenuPerformanceOutput メニューパフォーマンス取得の出力
type GetMenuPerformanceOutput struct {
	MenuPerformance []analytics.MenuPerformance `json:"menuPerformance"`
}

// GetDailySalesOutput 日次売上取得の出力
type GetDailySalesOutput struct {
	DailySales []analytics.DailySales `json:"dailySales"`
}

// GetMenuDailyTrendsOutput メニュー日次トレンド取得の出力
type GetMenuDailyTrendsOutput struct {
	MenuDailyTrends []analytics.MenuDailyTrend `json:"menuDailyTrends"`
}