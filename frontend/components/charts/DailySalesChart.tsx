"use client";

import {
  Card,
  CardContent,
  CardHeader,
  CardTitle,
} from "@workspace/ui/components/card";
import {
  Line,
  LineChart,
  ResponsiveContainer,
  Tooltip,
  XAxis,
  YAxis,
} from "recharts";
import type { DailySales } from "../../lib/api/analytics/types";

interface DailySalesChartProps {
  data: DailySales[];
  isLoading?: boolean;
}

export function DailySalesChart({ data, isLoading }: DailySalesChartProps) {
  if (isLoading) {
    return (
      <Card className='col-span-2'>
        <CardHeader>
          <CardTitle>日次売上トレンド</CardTitle>
        </CardHeader>
        <CardContent>
          <output
            className='h-[300px] flex items-center justify-center text-sm text-muted-foreground'
            aria-live='polite'
          >
            読み込み中...
          </output>
        </CardContent>
      </Card>
    );
  }

  if (!data || data.length === 0) {
    return (
      <Card className='col-span-2'>
        <CardHeader>
          <CardTitle>日次売上トレンド</CardTitle>
        </CardHeader>
        <CardContent>
          <div className='h-[300px] flex items-center justify-center text-sm text-muted-foreground'>
            データがありません
          </div>
        </CardContent>
      </Card>
    );
  }

  const chartData = data.map(item => ({
    day: new Date(item.day).toLocaleDateString("ja-JP", {
      month: "short",
      day: "numeric",
    }),
    fullDate: item.day,
    ordersCount: item.ordersCount,
    revenue: Math.round(item.revenue),
  }));

  return (
    <Card className='col-span-2'>
      <CardHeader>
        <CardTitle>日次売上トレンド（過去30日）</CardTitle>
      </CardHeader>
      <CardContent>
        <div
          className='h-[300px]'
          role='img'
          aria-label={`日次売上トレンドグラフ：過去30日間の注文数と売上額の推移を表示。最新の売上は${chartData[chartData.length - 1]?.revenue.toLocaleString()}円です。`}
        >
          <ResponsiveContainer width='100%' height='100%'>
            <LineChart
              data={chartData}
              margin={{
                top: 20,
                left: 12,
                right: 12,
                bottom: 5,
              }}
            >
              <XAxis
                dataKey='day'
                tickLine={false}
                axisLine={false}
                tick={{ fontSize: 12 }}
              />
              <YAxis
                yAxisId='orders'
                orientation='left'
                tickLine={false}
                axisLine={false}
                tick={{ fontSize: 12 }}
              />
              <YAxis
                yAxisId='revenue'
                orientation='right'
                tickLine={false}
                axisLine={false}
                tick={{ fontSize: 12 }}
              />
              <Tooltip
                content={({ active, payload, label }) => {
                  if (active && payload && payload.length) {
                    return (
                      <div className='rounded-lg border bg-background p-2 shadow-md'>
                        <div className='grid grid-cols-1 gap-2'>
                          <div className='flex flex-col'>
                            <span className='text-xs text-muted-foreground'>
                              日付
                            </span>
                            <span className='font-bold'>{label}</span>
                          </div>
                          <div className='flex flex-col'>
                            <span className='text-xs text-muted-foreground'>
                              注文数
                            </span>
                            <span className='font-bold'>
                              {payload[0]?.value}
                            </span>
                          </div>
                          <div className='flex flex-col'>
                            <span className='text-xs text-muted-foreground'>
                              売上
                            </span>
                            <span className='font-bold'>
                              ¥{payload[1]?.value?.toLocaleString()}
                            </span>
                          </div>
                        </div>
                      </div>
                    );
                  }
                  return null;
                }}
              />
              <Line
                yAxisId='orders'
                type='monotone'
                dataKey='ordersCount'
                stroke='hsl(var(--chart-1))'
                strokeWidth={2}
                dot={{ r: 4 }}
                activeDot={{ r: 6 }}
              />
              <Line
                yAxisId='revenue'
                type='monotone'
                dataKey='revenue'
                stroke='hsl(var(--chart-2))'
                strokeWidth={2}
                dot={{ r: 4 }}
                activeDot={{ r: 6 }}
              />
            </LineChart>
          </ResponsiveContainer>
        </div>
      </CardContent>
    </Card>
  );
}
