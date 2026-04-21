import useSWR from "swr";
import {
  getAnalyticsData,
  getCategorySales,
  getDailySales,
  getKPISummary,
  getMenuDailyTrends,
  getMenuPerformance,
  getTopMenus,
} from "@/lib/api/analytics/analytics";

// SWRキー定数
const SWR_KEYS = {
  analytics: "analytics",
  kpiSummary: "analytics/kpi",
  topMenus: (limit: number) => `analytics/top-menus/${limit}`,
  categorySales: "analytics/category-sales",
  menuPerformance: (limit: number) => `analytics/menu-performance/${limit}`,
  dailySales: "analytics/daily-sales",
  menuDailyTrends: (menuIds?: string[]) =>
    `analytics/menu-daily-trends${menuIds ? `/${menuIds.join(",")}` : ""}`,
} as const;

// リフレッシュ間隔（ミリ秒）
const REFRESH_INTERVAL = 60000; // 1分

/**
 * 統計データ全体を取得するフック
 */
export const useAnalyticsData = () => {
  return useSWR(SWR_KEYS.analytics, () => getAnalyticsData(), {
    refreshInterval: REFRESH_INTERVAL,
    revalidateOnFocus: false,
    dedupingInterval: 30000, // 30秒間のデータ重複排除
  });
};

/**
 * KPIサマリーを取得するフック
 */
export const useKPISummary = () => {
  return useSWR(SWR_KEYS.kpiSummary, () => getKPISummary(), {
    refreshInterval: REFRESH_INTERVAL,
    revalidateOnFocus: false,
    dedupingInterval: 30000,
  });
};

/**
 * トップメニューを取得するフック
 * @param limit 取得件数制限（デフォルト: 10）
 */
export const useTopMenus = (limit = 10) => {
  return useSWR(SWR_KEYS.topMenus(limit), () => getTopMenus(limit), {
    refreshInterval: REFRESH_INTERVAL,
    revalidateOnFocus: false,
    dedupingInterval: 30000,
  });
};

/**
 * カテゴリ別売上を取得するフック
 */
export const useCategorySales = () => {
  return useSWR(SWR_KEYS.categorySales, () => getCategorySales(), {
    refreshInterval: REFRESH_INTERVAL,
    revalidateOnFocus: false,
    dedupingInterval: 30000,
  });
};

/**
 * メニューパフォーマンスを取得するフック
 * @param limit 取得件数制限（デフォルト: 20）
 */
export const useMenuPerformance = (limit = 20) => {
  return useSWR(
    SWR_KEYS.menuPerformance(limit),
    () => getMenuPerformance(limit),
    {
      refreshInterval: REFRESH_INTERVAL,
      revalidateOnFocus: false,
      dedupingInterval: 30000,
    },
  );
};

/**
 * 日次売上トレンドを取得するフック
 */
export const useDailySales = () => {
  return useSWR(SWR_KEYS.dailySales, () => getDailySales(), {
    refreshInterval: REFRESH_INTERVAL,
    revalidateOnFocus: false,
    dedupingInterval: 30000,
  });
};

/**
 * メニュー日次トレンドを取得するフック
 * @param menuIds 対象メニューIDの配列（省略時は全メニュー）
 */
export const useMenuDailyTrends = (menuIds?: string[]) => {
  return useSWR(
    SWR_KEYS.menuDailyTrends(menuIds),
    () => getMenuDailyTrends(menuIds),
    {
      refreshInterval: REFRESH_INTERVAL,
      revalidateOnFocus: false,
      dedupingInterval: 30000,
    },
  );
};
