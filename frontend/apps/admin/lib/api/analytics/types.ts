// 統計データの型定義
export interface KPISummary {
  totalOrders: number;
  ordersToday: number;
  totalRevenue: number;
  averageOrderValue: number;
  totalItemsSold: number;
}

export interface TopMenu {
  menuId: string;
  name: string;
  quantitySold: number;
  revenue: number;
}

export interface CategorySales {
  categoryId: string;
  categoryName: string;
  totalQuantity: number;
  totalRevenue: number;
  avgPricePerItem: number;
}

export interface MenuPerformance {
  menuId: string;
  name: string;
  timesOrdered: number;
  revenue: number;
  avgPriceAtOrder: number;
  avgQuantityPerOrder: number;
}

export interface DailySales {
  day: string;
  ordersCount: number;
  revenue: number;
}

export interface MenuDailyTrend {
  day: string;
  menuId: string;
  name: string;
  quantitySold: number;
  revenue: number;
}

export interface AnalyticsData {
  kpiSummary: KPISummary;
  topMenus: TopMenu[];
  categorySales: CategorySales[];
  menuPerformance: MenuPerformance[];
  dailySales: DailySales[];
  menuDailyTrends: MenuDailyTrend[];
}

// API レスポンス型
export interface GetAnalyticsResponse {
  data: AnalyticsData;
}

export interface GetKPISummaryResponse {
  kpiSummary: KPISummary;
}

export interface GetTopMenusResponse {
  topMenus: TopMenu[];
}

export interface GetCategorySalesResponse {
  categorySales: CategorySales[];
}

export interface GetMenuPerformanceResponse {
  menuPerformance: MenuPerformance[];
}

export interface GetDailySalesResponse {
  dailySales: DailySales[];
}

export interface GetMenuDailyTrendsResponse {
  menuDailyTrends: MenuDailyTrend[];
}
