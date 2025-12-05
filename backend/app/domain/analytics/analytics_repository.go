package analytics

import (
	"context"
)

// AnalyticsRepository 統計データリポジトリのインターフェース（ポート）
type AnalyticsRepository interface {
	// GetKPISummary KPI指標のサマリーを取得する
	GetKPISummary(ctx context.Context) (*KPISummary, error)

	// GetTopMenus 人気メニューランキング（過去30日）を取得する
	GetTopMenus(ctx context.Context, limit int) ([]TopMenu, error)

	// GetCategorySales カテゴリ別売上（過去30日）を取得する
	GetCategorySales(ctx context.Context) ([]CategorySales, error)

	// GetMenuPerformance メニューパフォーマンスを取得する
	GetMenuPerformance(ctx context.Context, limit int) ([]MenuPerformance, error)

	// GetDailySales 日次売上トレンド（過去30日）を取得する
	GetDailySales(ctx context.Context) ([]DailySales, error)

	// GetMenuDailyTrends メニュー日次トレンド（過去30日）を取得する
	GetMenuDailyTrends(ctx context.Context, menuIds []string) ([]MenuDailyTrend, error)

	// GetAnalyticsData 全統計データを取得する
	GetAnalyticsData(ctx context.Context) (*AnalyticsData, error)
}