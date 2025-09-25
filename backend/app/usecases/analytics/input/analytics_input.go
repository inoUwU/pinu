package input

// GetAnalyticsInput 統計データ取得の入力
type GetAnalyticsInput struct {
	// 現在は特別な入力パラメーターは不要
	// 将来的に日付範囲の指定などを追加可能
}

// GetKPISummaryInput KPIサマリー取得の入力
type GetKPISummaryInput struct {
	// 現在は特別な入力パラメーターは不要
}

// GetTopMenusInput トップメニュー取得の入力
type GetTopMenusInput struct {
	Limit int `json:"limit"` // 取得件数制限
}

// GetCategorySalesInput カテゴリ別売上取得の入力
type GetCategorySalesInput struct {
	// 現在は特別な入力パラメーターは不要
}

// GetMenuPerformanceInput メニューパフォーマンス取得の入力
type GetMenuPerformanceInput struct {
	Limit int `json:"limit"` // 取得件数制限
}

// GetDailySalesInput 日次売上取得の入力
type GetDailySalesInput struct {
	// 現在は特別な入力パラメーターは不要
}

// GetMenuDailyTrendsInput メニュー日次トレンド取得の入力
type GetMenuDailyTrendsInput struct {
	MenuIds []string `json:"menuIds"` // 対象メニューID配列（空の場合は全メニュー）
}