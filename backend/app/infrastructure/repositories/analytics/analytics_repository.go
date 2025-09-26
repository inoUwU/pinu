package repositories

import (
	"context"
	"database/sql"
	"inoUwU/pinu/app/domain/analytics"
	"strings"
	"time"

	"github.com/samber/do"
	"github.com/uptrace/bun"
)

// AnalyticsRepositoryImpl 統計データリポジトリの実装（アダプター）
type AnalyticsRepositoryImpl struct {
	db *bun.DB
}

// NewAnalyticsRepository 統計データリポジトリの実装を生成する
func NewAnalyticsRepository(i *do.Injector) (analytics.IAnalyticsRepository, error) {
	db := do.MustInvokeNamed[*bun.DB](i, "db")
	return &AnalyticsRepositoryImpl{db: db}, nil
}

// GetKPISummary KPI指標のサマリーを取得する（v_kpi_orders_menuビューを使用）
func (r *AnalyticsRepositoryImpl) GetKPISummary(ctx context.Context) (*analytics.KPISummary, error) {
	type KPIResult struct {
		TotalOrders      int64           `bun:"total_orders"`
		OrdersToday      int64           `bun:"orders_today"`
		TotalRevenue     sql.NullFloat64 `bun:"total_revenue"`
		AverageOrderValue sql.NullFloat64 `bun:"average_order_value"`
		TotalItemsSold   int64           `bun:"total_items_sold"`
	}

	var result KPIResult
	if err := r.db.NewRaw("SELECT * FROM v_kpi_orders_menu").Scan(ctx, &result); err != nil {
		return nil, err
	}

	return &analytics.KPISummary{
		TotalOrders:      result.TotalOrders,
		OrdersToday:      result.OrdersToday,
		TotalRevenue:     result.TotalRevenue.Float64,
		AverageOrderValue: result.AverageOrderValue.Float64,
		TotalItemsSold:   result.TotalItemsSold,
	}, nil
}

// GetTopMenus 人気メニューランキング（過去30日）を取得する（v_top_menus_30dビューを使用）
func (r *AnalyticsRepositoryImpl) GetTopMenus(ctx context.Context, limit int) ([]analytics.TopMenu, error) {
	type TopMenuResult struct {
		MenuID       string          `bun:"menu_id"`
		Name         string          `bun:"name"`
		QuantitySold int64           `bun:"quantity_sold"`
		Revenue      sql.NullFloat64 `bun:"revenue"`
	}

	var results []TopMenuResult
	query := r.db.NewRaw("SELECT * FROM v_top_menus_30d LIMIT ?", limit)
	if err := query.Scan(ctx, &results); err != nil {
		return nil, err
	}

	topMenus := make([]analytics.TopMenu, len(results))
	for i, result := range results {
		topMenus[i] = analytics.TopMenu{
			MenuID:       result.MenuID,
			Name:         result.Name,
			QuantitySold: result.QuantitySold,
			Revenue:      result.Revenue.Float64,
		}
	}

	return topMenus, nil
}

// GetCategorySales カテゴリ別売上（過去30日）を取得する（v_sales_by_category_30dビューを使用）
func (r *AnalyticsRepositoryImpl) GetCategorySales(ctx context.Context) ([]analytics.CategorySales, error) {
	type CategorySalesResult struct {
		CategoryID      string          `bun:"category_id"`
		CategoryName    string          `bun:"category_name"`
		TotalQuantity   int64           `bun:"total_quantity"`
		TotalRevenue    sql.NullFloat64 `bun:"total_revenue"`
		AvgPricePerItem sql.NullFloat64 `bun:"avg_price_per_item"`
	}

	var results []CategorySalesResult
	if err := r.db.NewRaw("SELECT * FROM v_sales_by_category_30d").Scan(ctx, &results); err != nil {
		return nil, err
	}

	categorySales := make([]analytics.CategorySales, len(results))
	for i, result := range results {
		categorySales[i] = analytics.CategorySales{
			CategoryID:      result.CategoryID,
			CategoryName:    result.CategoryName,
			TotalQuantity:   result.TotalQuantity,
			TotalRevenue:    result.TotalRevenue.Float64,
			AvgPricePerItem: result.AvgPricePerItem.Float64,
		}
	}

	return categorySales, nil
}

// GetMenuPerformance メニューパフォーマンスを取得する（v_menu_performanceビューを使用）
func (r *AnalyticsRepositoryImpl) GetMenuPerformance(ctx context.Context, limit int) ([]analytics.MenuPerformance, error) {
	type MenuPerformanceResult struct {
		MenuID              string          `bun:"menu_id"`
		Name                string          `bun:"name"`
		TimesOrdered        int64           `bun:"times_ordered"`
		Revenue             sql.NullFloat64 `bun:"revenue"`
		AvgPriceAtOrder     sql.NullFloat64 `bun:"avg_price_at_order"`
		AvgQuantityPerOrder sql.NullFloat64 `bun:"avg_quantity_per_order"`
	}

	var results []MenuPerformanceResult
	query := r.db.NewRaw("SELECT * FROM v_menu_performance LIMIT ?", limit)
	if err := query.Scan(ctx, &results); err != nil {
		return nil, err
	}

	menuPerformance := make([]analytics.MenuPerformance, len(results))
	for i, result := range results {
		menuPerformance[i] = analytics.MenuPerformance{
			MenuID:              result.MenuID,
			Name:                result.Name,
			TimesOrdered:        result.TimesOrdered,
			Revenue:             result.Revenue.Float64,
			AvgPriceAtOrder:     result.AvgPriceAtOrder.Float64,
			AvgQuantityPerOrder: result.AvgQuantityPerOrder.Float64,
		}
	}

	return menuPerformance, nil
}

// GetDailySales 日次売上トレンド（過去30日）を取得する（v_orders_daily_30dビューを使用）
func (r *AnalyticsRepositoryImpl) GetDailySales(ctx context.Context) ([]analytics.DailySales, error) {
	type DailySalesResult struct {
		Day         time.Time       `bun:"day"`
		OrdersCount int64           `bun:"orders_count"`
		Revenue     sql.NullFloat64 `bun:"revenue"`
	}

	var results []DailySalesResult
	if err := r.db.NewRaw("SELECT * FROM v_orders_daily_30d").Scan(ctx, &results); err != nil {
		return nil, err
	}

	dailySales := make([]analytics.DailySales, len(results))
	for i, result := range results {		
		dailySales[i] = analytics.DailySales{
			Day:         result.Day,
			OrdersCount: result.OrdersCount,
			Revenue:     result.Revenue.Float64,
		}
	}

	return dailySales, nil
}

// GetMenuDailyTrends メニュー日次トレンド（過去30日）を取得する（v_menu_daily_trend_30dビューを使用）
func (r *AnalyticsRepositoryImpl) GetMenuDailyTrends(ctx context.Context, menuIds []string) ([]analytics.MenuDailyTrend, error) {
	type MenuDailyTrendResult struct {
		Day          time.Time       `bun:"day"`
		MenuID       string          `bun:"menu_id"`
		Name         string          `bun:"name"`
		QuantitySold int64           `bun:"quantity_sold"`
		Revenue      sql.NullFloat64 `bun:"revenue"`
	}

	var results []MenuDailyTrendResult
	var query *bun.RawQuery
	
	if len(menuIds) > 0 {
		// 特定のメニューIDに絞り込み
		// PostgreSQL配列形式に変換: {id1,id2,id3}
		pgArray := "{" + strings.Join(menuIds, ",") + "}"
		query = r.db.NewRaw("SELECT * FROM v_menu_daily_trend_30d WHERE menu_id = ANY(?::text[])", pgArray)
	} else {
		// 全てのメニュー
		query = r.db.NewRaw("SELECT * FROM v_menu_daily_trend_30d")
	}
	
	if err := query.Scan(ctx, &results); err != nil {
		return nil, err
	}

	menuDailyTrends := make([]analytics.MenuDailyTrend, len(results))
	for i, result := range results {
		menuDailyTrends[i] = analytics.MenuDailyTrend{
			Day:          result.Day,
			MenuID:       result.MenuID,
			Name:         result.Name,
			QuantitySold: result.QuantitySold,
			Revenue:      result.Revenue.Float64,
		}
	}

	return menuDailyTrends, nil
}

// GetAnalyticsData 全統計データを取得する
func (r *AnalyticsRepositoryImpl) GetAnalyticsData(ctx context.Context) (*analytics.AnalyticsData, error) {
	// 各統計データを並行取得
	kpiSummary, err := r.GetKPISummary(ctx)
	if err != nil {
		return nil, err
	}

	topMenus, err := r.GetTopMenus(ctx, 10) // トップ10
	if err != nil {
		return nil, err
	}

	categorySales, err := r.GetCategorySales(ctx)
	if err != nil {
		return nil, err
	}

	menuPerformance, err := r.GetMenuPerformance(ctx, 20) // トップ20
	if err != nil {
		return nil, err
	}

	dailySales, err := r.GetDailySales(ctx)
	if err != nil {
		return nil, err
	}

	// トップメニューのIDを取得してメニュー日次トレンドを取得
	maxMenus := 5
	if len(topMenus) < maxMenus {
		maxMenus = len(topMenus)
	}
	topMenuIds := make([]string, maxMenus) // トップ5メニューのトレンド
	for i := 0; i < maxMenus; i++ {
		topMenuIds[i] = topMenus[i].MenuID
	}
	
	menuDailyTrends, err := r.GetMenuDailyTrends(ctx, topMenuIds)
	if err != nil {
		return nil, err
	}

	return &analytics.AnalyticsData{
		KPISummary:      *kpiSummary,
		TopMenus:        topMenus,
		CategorySales:   categorySales,
		MenuPerformance: menuPerformance,
		DailySales:      dailySales,
		MenuDailyTrends: menuDailyTrends,
	}, nil
}