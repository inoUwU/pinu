"use client";

import {
  Card,
  CardContent,
  CardHeader,
  CardTitle,
} from "@workspace/ui/components/card";
import {
  Bar,
  BarChart,
  ResponsiveContainer,
  Tooltip,
  XAxis,
  YAxis,
} from "recharts";
import type { TopMenu } from "../../lib/api/analytics/types";

interface TopMenusChartProps {
  data: TopMenu[];
  isLoading?: boolean;
}

export function TopMenusChart({ data, isLoading }: TopMenusChartProps) {
  if (isLoading) {
    return (
      <Card>
        <CardHeader>
          <CardTitle>トップメニュー</CardTitle>
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
      <Card>
        <CardHeader>
          <CardTitle>トップメニュー</CardTitle>
        </CardHeader>
        <CardContent>
          <div className='h-[300px] flex items-center justify-center text-sm text-muted-foreground'>
            データがありません
          </div>
        </CardContent>
      </Card>
    );
  }

  const chartData = data.map(menu => ({
    name:
      menu.name.length > 10 ? `${menu.name.substring(0, 10)}...` : menu.name,
    fullName: menu.name,
    quantitySold: menu.quantitySold,
    revenue: Math.round(menu.revenue),
  }));

  return (
    <Card>
      <CardHeader>
        <CardTitle>トップメニュー（過去30日）</CardTitle>
      </CardHeader>
      <CardContent>
        <div
          className='h-[300px]'
          role='img'
          aria-label='トップメニュー売上棒グラフ：過去30日間で最も売れたメニューの販売数と売上を表示。'
        >
          <ResponsiveContainer width='100%' height='100%'>
            <BarChart
              data={chartData}
              margin={{
                top: 20,
                left: 12,
                right: 12,
                bottom: 20,
              }}
            >
              <XAxis
                dataKey='name'
                tickLine={false}
                axisLine={false}
                tick={{ fontSize: 12 }}
                angle={-45}
                textAnchor='end'
                height={60}
              />
              <YAxis
                tickLine={false}
                axisLine={false}
                tick={{ fontSize: 12 }}
              />
              <Tooltip
                content={({ active, payload }) => {
                  if (active && payload && payload.length) {
                    const data = payload[0].payload;
                    return (
                      <div className='rounded-lg border bg-background p-2 shadow-md'>
                        <div className='grid grid-cols-2 gap-2'>
                          <div className='flex flex-col'>
                            <span className='text-xs text-muted-foreground'>
                              メニュー
                            </span>
                            <span className='font-bold'>{data.fullName}</span>
                          </div>
                          <div className='flex flex-col'>
                            <span className='text-xs text-muted-foreground'>
                              販売数量
                            </span>
                            <span className='font-bold'>
                              {data.quantitySold}
                            </span>
                          </div>
                        </div>
                      </div>
                    );
                  }
                  return null;
                }}
              />
              <Bar
                dataKey='quantitySold'
                fill='hsl(var(--chart-1))'
                radius={[4, 4, 0, 0]}
              />
            </BarChart>
          </ResponsiveContainer>
        </div>
      </CardContent>
    </Card>
  );
}
