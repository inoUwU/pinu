package analytics

import (
	"time"
)

// KPIサマリー - 主要業績指標
type KPISummary struct {
	TotalOrders      int64   `json:"totalOrders"`      // 総注文数
	OrdersToday      int64   `json:"ordersToday"`      // 本日の注文数
	TotalRevenue     float64 `json:"totalRevenue"`     // 総売上
	AverageOrderValue float64 `json:"averageOrderValue"` // 平均注文額
	TotalItemsSold   int64   `json:"totalItemsSold"`   // 総販売アイテム数
}

// トップメニュー（過去30日）
type TopMenu struct {
	MenuID       string  `json:"menuId"`       // メニューID
	Name         string  `json:"name"`         // メニュー名
	QuantitySold int64   `json:"quantitySold"` // 販売数量
	Revenue      float64 `json:"revenue"`      // 売上
}

// カテゴリ別売上（過去30日）
type CategorySales struct {
	CategoryID      string  `json:"categoryId"`      // カテゴリID
	CategoryName    string  `json:"categoryName"`    // カテゴリ名
	TotalQuantity   int64   `json:"totalQuantity"`   // 総販売数量
	TotalRevenue    float64 `json:"totalRevenue"`    // 総売上
	AvgPricePerItem float64 `json:"avgPricePerItem"` // アイテム単価平均
}

// メニューパフォーマンス
type MenuPerformance struct {
	MenuID               string  `json:"menuId"`               // メニューID
	Name                 string  `json:"name"`                 // メニュー名
	TimesOrdered         int64   `json:"timesOrdered"`         // 注文回数
	AvgPriceAtOrder      float64 `json:"avgPriceAtOrder"`      // 注文時平均価格
	AvgQuantityPerOrder  float64 `json:"avgQuantityPerOrder"`  // 注文あたり平均数量
	Revenue              float64 `json:"revenue"`              // 売上
}

// 日次売上トレンド
type DailySales struct {
	Day          time.Time `json:"day"`          // 日付
	OrdersCount  int64     `json:"ordersCount"`  // 注文数
	Revenue      float64   `json:"revenue"`      // 売上
}

// メニュー日次トレンド
type MenuDailyTrend struct {
	Day          time.Time `json:"day"`          // 日付
	MenuID       string    `json:"menuId"`       // メニューID
	Name         string    `json:"name"`         // メニュー名
	QuantitySold int64     `json:"quantitySold"` // 販売数量
	Revenue      float64   `json:"revenue"`      // 売上
}

// 統計データ全体
type AnalyticsData struct {
	KPISummary        KPISummary         `json:"kpiSummary"`        // KPI要約
	TopMenus          []TopMenu          `json:"topMenus"`          // トップメニュー（過去30日）
	CategorySales     []CategorySales    `json:"categorySales"`     // カテゴリ別売上（過去30日）
	MenuPerformance   []MenuPerformance  `json:"menuPerformance"`   // メニューパフォーマンス
	DailySales        []DailySales       `json:"dailySales"`        // 日次売上トレンド（過去30日）
	MenuDailyTrends   []MenuDailyTrend   `json:"menuDailyTrends"`   // メニュー日次トレンド（過去30日）
}