"use client";

import {
  Card,
  CardContent,
  CardHeader,
  CardTitle,
} from "@workspace/ui/components/card";
import type { KPISummary } from "@/lib/api/analytics/types";

interface KPISummaryCardsProps {
  data?: KPISummary;
  isLoading?: boolean;
}

interface StatCardProps {
  title: string;
  value: string | number;
  description?: string;
  isLoading?: boolean;
}

function StatCard({ title, value, description, isLoading }: StatCardProps) {
  if (isLoading) {
    return (
      <Card>
        <CardHeader className='flex flex-row items-center justify-between space-y-0 pb-2'>
          <CardTitle className='text-sm font-medium'>{title}</CardTitle>
        </CardHeader>
        <CardContent>
          <div className='text-2xl font-bold'>
            <div className='h-8 bg-muted animate-pulse rounded' />
          </div>
          {description && (
            <p className='text-xs text-muted-foreground mt-1'>{description}</p>
          )}
        </CardContent>
      </Card>
    );
  }

  return (
    <Card>
      <CardHeader className='flex flex-row items-center justify-between space-y-0 pb-2'>
        <CardTitle className='text-sm font-medium'>{title}</CardTitle>
      </CardHeader>
      <CardContent>
        <div className='text-2xl font-bold'>{value}</div>
        {description && (
          <p className='text-xs text-muted-foreground mt-1'>{description}</p>
        )}
      </CardContent>
    </Card>
  );
}

export function KPISummaryCards({ data, isLoading }: KPISummaryCardsProps) {
  return (
    <div className='grid gap-4 md:grid-cols-2 lg:grid-cols-5'>
      <StatCard
        title='総注文数'
        value={isLoading ? "" : data?.totalOrders.toLocaleString() || "0"}
        description='全期間'
        isLoading={isLoading}
      />

      <StatCard
        title='本日の注文'
        value={isLoading ? "" : data?.ordersToday.toLocaleString() || "0"}
        description='今日の注文数'
        isLoading={isLoading}
      />

      <StatCard
        title='総売上'
        value={
          isLoading
            ? ""
            : data
              ? `¥${Math.round(data.totalRevenue).toLocaleString()}`
              : "¥0"
        }
        description='全期間'
        isLoading={isLoading}
      />

      <StatCard
        title='平均注文額'
        value={
          isLoading
            ? ""
            : data
              ? `¥${Math.round(data.averageOrderValue).toLocaleString()}`
              : "¥0"
        }
        description='1注文あたり'
        isLoading={isLoading}
      />

      <StatCard
        title='総販売アイテム数'
        value={isLoading ? "" : data?.totalItemsSold.toLocaleString() || "0"}
        description='全期間'
        isLoading={isLoading}
      />
    </div>
  );
}
