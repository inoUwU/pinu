package analytics

import (
	"context"
	"inoUwU/pinu/app/domain/analytics"
	"inoUwU/pinu/app/domain/port"
	"inoUwU/pinu/app/usecases/analytics/input"
	"inoUwU/pinu/app/usecases/analytics/output"

	"github.com/samber/do"
)

// IAnalyticsUsecase 統計データユースケースのインターフェース
type IAnalyticsUsecase interface {
	// GetAnalyticsData 統計データ全体を取得する
	GetAnalyticsData(ctx context.Context, input *input.GetAnalyticsInput) (*output.GetAnalyticsOutput, error)

	// GetKPISummary KPIサマリーを取得する
	GetKPISummary(ctx context.Context, input *input.GetKPISummaryInput) (*output.GetKPISummaryOutput, error)

	// GetTopMenus トップメニューを取得する
	GetTopMenus(ctx context.Context, input *input.GetTopMenusInput) (*output.GetTopMenusOutput, error)

	// GetCategorySales カテゴリ別売上を取得する
	GetCategorySales(ctx context.Context, input *input.GetCategorySalesInput) (*output.GetCategorySalesOutput, error)

	// GetMenuPerformance メニューパフォーマンスを取得する
	GetMenuPerformance(ctx context.Context, input *input.GetMenuPerformanceInput) (*output.GetMenuPerformanceOutput, error)

	// GetDailySales 日次売上を取得する
	GetDailySales(ctx context.Context, input *input.GetDailySalesInput) (*output.GetDailySalesOutput, error)

	// GetMenuDailyTrends メニュー日次トレンドを取得する
	GetMenuDailyTrends(ctx context.Context, input *input.GetMenuDailyTrendsInput) (*output.GetMenuDailyTrendsOutput, error)
}

// AnalyticsUsecaseImpl 統計データユースケースの実装
type AnalyticsUsecaseImpl struct {
	analyticsRepo analytics.AnalyticsRepository
	logger        port.Logger
	txManager     port.TransactionManager
}

// NewAnalyticsUsecase 統計データユースケースを生成する
func NewAnalyticsUsecase(i *do.Injector) (IAnalyticsUsecase, error) {
	repository := do.MustInvoke[analytics.AnalyticsRepository](i)
	logger := do.MustInvokeNamed[port.Logger](i, "logger")
	txManager := do.MustInvokeNamed[port.TransactionManager](i, "tx")

	return &AnalyticsUsecaseImpl{
		analyticsRepo: repository,
		logger:        logger,
		txManager:     txManager,
	}, nil
}

// GetAnalyticsData 統計データ全体を取得する
func (u *AnalyticsUsecaseImpl) GetAnalyticsData(ctx context.Context, input *input.GetAnalyticsInput) (*output.GetAnalyticsOutput, error) {
	data, err := u.analyticsRepo.GetAnalyticsData(ctx)
	if err != nil {
		u.logger.Error("failed to get analytics data", "error", err)
		return nil, err
	}

	u.logger.Info("analytics data retrieved successfully")

	return &output.GetAnalyticsOutput{
		Data: *data,
	}, nil
}

// GetKPISummary KPIサマリーを取得する
func (u *AnalyticsUsecaseImpl) GetKPISummary(ctx context.Context, input *input.GetKPISummaryInput) (*output.GetKPISummaryOutput, error) {
	kpiSummary, err := u.analyticsRepo.GetKPISummary(ctx)
	if err != nil {
		u.logger.Error("failed to get KPI summary", "error", err)
		return nil, err
	}

	u.logger.Info("KPI summary retrieved successfully")

	return &output.GetKPISummaryOutput{
		KPISummary: *kpiSummary,
	}, nil
}

// GetTopMenus トップメニューを取得する
func (u *AnalyticsUsecaseImpl) GetTopMenus(ctx context.Context, input *input.GetTopMenusInput) (*output.GetTopMenusOutput, error) {
	limit := input.Limit
	if limit <= 0 {
		limit = 10 // デフォルト値
	}

	topMenus, err := u.analyticsRepo.GetTopMenus(ctx, limit)
	if err != nil {
		u.logger.Error("failed to get top menus", "error", err, "limit", limit)
		return nil, err
	}

	u.logger.Info("top menus retrieved successfully", "count", len(topMenus), "limit", limit)

	return &output.GetTopMenusOutput{
		TopMenus: topMenus,
	}, nil
}

// GetCategorySales カテゴリ別売上を取得する
func (u *AnalyticsUsecaseImpl) GetCategorySales(ctx context.Context, input *input.GetCategorySalesInput) (*output.GetCategorySalesOutput, error) {
	categorySales, err := u.analyticsRepo.GetCategorySales(ctx)
	if err != nil {
		u.logger.Error("failed to get category sales", "error", err)
		return nil, err
	}

	u.logger.Info("category sales retrieved successfully", "count", len(categorySales))

	return &output.GetCategorySalesOutput{
		CategorySales: categorySales,
	}, nil
}

// GetMenuPerformance メニューパフォーマンスを取得する
func (u *AnalyticsUsecaseImpl) GetMenuPerformance(ctx context.Context, input *input.GetMenuPerformanceInput) (*output.GetMenuPerformanceOutput, error) {
	limit := input.Limit
	if limit <= 0 {
		limit = 20 // デフォルト値
	}

	menuPerformance, err := u.analyticsRepo.GetMenuPerformance(ctx, limit)
	if err != nil {
		u.logger.Error("failed to get menu performance", "error", err, "limit", limit)
		return nil, err
	}

	u.logger.Info("menu performance retrieved successfully", "count", len(menuPerformance), "limit", limit)

	return &output.GetMenuPerformanceOutput{
		MenuPerformance: menuPerformance,
	}, nil
}

// GetDailySales 日次売上を取得する
func (u *AnalyticsUsecaseImpl) GetDailySales(ctx context.Context, input *input.GetDailySalesInput) (*output.GetDailySalesOutput, error) {
	dailySales, err := u.analyticsRepo.GetDailySales(ctx)
	if err != nil {
		u.logger.Error("failed to get daily sales", "error", err)
		return nil, err
	}

	u.logger.Info("daily sales retrieved successfully", "count", len(dailySales))

	return &output.GetDailySalesOutput{
		DailySales: dailySales,
	}, nil
}

// GetMenuDailyTrends メニュー日次トレンドを取得する
func (u *AnalyticsUsecaseImpl) GetMenuDailyTrends(ctx context.Context, input *input.GetMenuDailyTrendsInput) (*output.GetMenuDailyTrendsOutput, error) {
	menuDailyTrends, err := u.analyticsRepo.GetMenuDailyTrends(ctx, input.MenuIds)
	if err != nil {
		u.logger.Error("failed to get menu daily trends", "error", err, "menuIds", input.MenuIds)
		return nil, err
	}

	u.logger.Info("menu daily trends retrieved successfully", "count", len(menuDailyTrends), "menuIds", input.MenuIds)

	return &output.GetMenuDailyTrendsOutput{
		MenuDailyTrends: menuDailyTrends,
	}, nil
}