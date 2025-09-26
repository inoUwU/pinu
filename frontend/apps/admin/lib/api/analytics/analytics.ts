import api from "../api";
import type {
  AnalyticsData,
  GetAnalyticsResponse,
  GetKPISummaryResponse,
  GetTopMenusResponse,
  GetCategorySalesResponse,
  GetMenuPerformanceResponse,
  GetDailySalesResponse,
  GetMenuDailyTrendsResponse,
} from "./types";

const ANALYTICS_ENDPOINTS = {
  analytics: "analytics",
  kpi: "analytics/kpi",
  topMenus: "analytics/top-menus",
  categorySales: "analytics/category-sales",
  menuPerformance: "analytics/menu-performance", 
  dailySales: "analytics/daily-sales",
  menuDailyTrends: "analytics/menu-daily-trends",
} as const;

/**
 * 統計データ全体を取得する
 */
export const getAnalyticsData = async (): Promise<AnalyticsData> => {
  const response: GetAnalyticsResponse = await api.get(ANALYTICS_ENDPOINTS.analytics).json();
  return response.data;
};

/**
 * KPIサマリーを取得する
 */
export const getKPISummary = async () => {
  const response: GetKPISummaryResponse = await api.get(ANALYTICS_ENDPOINTS.kpi).json();
  return response.kpiSummary;
};

/**
 * トップメニューを取得する
 * @param limit 取得件数制限（デフォルト: 10）
 */
export const getTopMenus = async (limit = 10) => {
  const response: GetTopMenusResponse = await api
    .get(ANALYTICS_ENDPOINTS.topMenus, {
      searchParams: { limit: limit.toString() },
    })
    .json();
  return response.topMenus;
};

/**
 * カテゴリ別売上を取得する
 */
export const getCategorySales = async () => {
  const response: GetCategorySalesResponse = await api.get(ANALYTICS_ENDPOINTS.categorySales).json();
  return response.categorySales;
};

/**
 * メニューパフォーマンスを取得する
 * @param limit 取得件数制限（デフォルト: 20）
 */
export const getMenuPerformance = async (limit = 20) => {
  const response: GetMenuPerformanceResponse = await api
    .get(ANALYTICS_ENDPOINTS.menuPerformance, {
      searchParams: { limit: limit.toString() },
    })
    .json();
  return response.menuPerformance;
};

/**
 * 日次売上トレンドを取得する
 */
export const getDailySales = async () => {
  const response: GetDailySalesResponse = await api.get(ANALYTICS_ENDPOINTS.dailySales).json();
  return response.dailySales;
};

/**
 * メニュー日次トレンドを取得する
 * @param menuIds 対象メニューIDの配列（省略時は全メニュー）
 */
export const getMenuDailyTrends = async (menuIds?: string[]) => {
  const searchParams: Record<string, string> = {};
  if (menuIds && menuIds.length > 0) {
    searchParams.menuIds = menuIds.join(",");
  }

  const response: GetMenuDailyTrendsResponse = await api
    .get(ANALYTICS_ENDPOINTS.menuDailyTrends, { searchParams })
    .json();
  return response.menuDailyTrends;
};