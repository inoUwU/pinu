"use client";

import { CategorySalesChart } from "@/components/charts/CategorySalesChart";
import { DailySalesChart } from "@/components/charts/DailySalesChart";
import { KPISummaryCards } from "@/components/charts/KPISummaryCards";
import { TopMenusChart } from "@/components/charts/TopMenusChart";
import { useAnalyticsData, useKPISummary } from "@/hooks/useAnalytics";

export default function AdminPage() {
  // SWRを使用してデータを取得
  const {
    data: analyticsData,
    isLoading: analyticsLoading,
    error: analyticsError,
  } = useAnalyticsData();
  const {
    data: kpiData,
    isLoading: kpiLoading,
    error: kpiError,
  } = useKPISummary();

  // エラーハンドリング
  if (analyticsError || kpiError) {
    return (
      <div className='flex min-h-screen items-center justify-center'>
        <div className='text-center'>
          <h2 className='text-lg font-semibold text-red-600'>
            データの取得に失敗しました
          </h2>
          <p className='text-sm text-muted-foreground mt-2'>
            サーバーとの接続を確認してページを再読み込みしてください。
          </p>
        </div>
      </div>
    );
  }

  return (
    <div className='flex-1 space-y-6 p-6'>
      <div className='flex items-center justify-between'>
        <h1 className='text-3xl font-bold tracking-tight'>ダッシュボード</h1>
        <div className='text-sm text-muted-foreground'>
          {analyticsLoading
            ? "更新中..."
            : `最終更新: ${new Date().toLocaleTimeString("ja-JP")}`}
        </div>
      </div>

      {/* KPI サマリーカード */}
      <KPISummaryCards data={kpiData} isLoading={kpiLoading} />

      {/* チャートセクション */}
      <div className='grid gap-6 md:grid-cols-2'>
        {/* トップメニューチャート */}
        <TopMenusChart
          data={analyticsData?.topMenus || []}
          isLoading={analyticsLoading}
        />

        {/* カテゴリ別売上チャート */}
        <CategorySalesChart
          data={analyticsData?.categorySales || []}
          isLoading={analyticsLoading}
        />
      </div>

      {/* 日次売上トレンドチャート（全幅） */}
      <div className='grid gap-6'>
        <DailySalesChart
          data={analyticsData?.dailySales || []}
          isLoading={analyticsLoading}
        />
      </div>

      {/* データがない場合のメッセージ */}
      {!analyticsLoading &&
        (!analyticsData ||
          (analyticsData.topMenus.length === 0 &&
            analyticsData.categorySales.length === 0 &&
            analyticsData.dailySales.length === 0)) && (
          <div className='text-center py-12'>
            <h3 className='text-lg font-semibold text-muted-foreground'>
              統計データがありません
            </h3>
            <p className='text-sm text-muted-foreground mt-2'>
              注文データが蓄積されると統計が表示されます。
            </p>
          </div>
        )}
    </div>
  );
}
