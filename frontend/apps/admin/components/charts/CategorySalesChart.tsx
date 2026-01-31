"use client";

import {
  Card,
  CardContent,
  CardHeader,
  CardTitle,
} from "@workspace/ui/components/card";
import { Cell, Pie, PieChart, ResponsiveContainer, Tooltip } from "recharts";
import type { CategorySales } from "@/lib/api/analytics/types";

interface CategorySalesChartProps {
  data: CategorySales[];
  isLoading?: boolean;
}

const CHART_COLORS = [
  "hsl(var(--chart-1))",
  "hsl(var(--chart-2))",
  "hsl(var(--chart-3))",
  "hsl(var(--chart-4))",
  "hsl(var(--chart-5))",
];

export function CategorySalesChart({
  data,
  isLoading,
}: CategorySalesChartProps) {
  if (isLoading) {
    return (
      <Card>
        <CardHeader>
          <CardTitle>カテゴリ別売上</CardTitle>
        </CardHeader>
        <CardContent>
          <div
            className='h-[300px] flex items-center justify-center text-sm text-muted-foreground'
            role='status'
            aria-live='polite'
          >
            読み込み中...
          </div>
        </CardContent>
      </Card>
    );
  }

  if (!data || data.length === 0) {
    return (
      <Card>
        <CardHeader>
          <CardTitle>カテゴリ別売上</CardTitle>
        </CardHeader>
        <CardContent>
          <div className='h-[300px] flex items-center justify-center text-sm text-muted-foreground'>
            データがありません
          </div>
        </CardContent>
      </Card>
    );
  }

  const chartData = data.map((category, index) => ({
    name: category.categoryName,
    value: Math.round(category.totalRevenue),
    quantity: category.totalQuantity,
    fill: CHART_COLORS[index % CHART_COLORS.length],
  }));

  const totalRevenue = chartData.reduce((sum, item) => sum + item.value, 0);

  return (
    <Card>
      <CardHeader>
        <CardTitle>カテゴリ別売上（過去30日）</CardTitle>
      </CardHeader>
      <CardContent>
        <div
          className='h-[300px]'
          role='img'
          aria-label={`カテゴリ別売上円グラフ：過去30日間のカテゴリ別売上構成比。合計売上は${totalRevenue.toLocaleString()}円です。`}
        >
          <ResponsiveContainer width='100%' height='100%'>
            <PieChart>
              <Tooltip
                content={({ active, payload }) => {
                  if (active && payload && payload.length) {
                    const data = payload[0].payload;
                    return (
                      <div className='rounded-lg border bg-background p-2 shadow-md'>
                        <div className='grid grid-cols-1 gap-2'>
                          <div className='flex flex-col'>
                            <span className='text-xs text-muted-foreground'>
                              カテゴリ
                            </span>
                            <span className='font-bold'>{data.name}</span>
                          </div>
                          <div className='flex flex-col'>
                            <span className='text-xs text-muted-foreground'>
                              売上
                            </span>
                            <span className='font-bold'>
                              ¥{data.value.toLocaleString()}
                            </span>
                          </div>
                        </div>
                      </div>
                    );
                  }
                  return null;
                }}
              />
              <Pie
                data={chartData}
                cx='50%'
                cy='50%'
                innerRadius={60}
                outerRadius={120}
                paddingAngle={5}
                dataKey='value'
              >
                {chartData.map(entry => (
                  <Cell key={`cell-${entry.name}`} fill={entry.fill} />
                ))}
              </Pie>
            </PieChart>
          </ResponsiveContainer>
        </div>

        <div className='mt-4 space-y-2'>
          {chartData.map(item => (
            <div
              key={item.name}
              className='flex items-center justify-between text-sm'
            >
              <div className='flex items-center'>
                <div
                  className='w-3 h-3 rounded-full mr-2'
                  style={{ backgroundColor: item.fill }}
                />
                <span>{item.name}</span>
              </div>
              <div className='text-right'>
                <div className='font-medium'>
                  ¥{item.value.toLocaleString()}
                </div>
                <div className='text-muted-foreground text-xs'>
                  {((item.value / totalRevenue) * 100).toFixed(1)}%
                </div>
              </div>
            </div>
          ))}
        </div>
      </CardContent>
    </Card>
  );
}
