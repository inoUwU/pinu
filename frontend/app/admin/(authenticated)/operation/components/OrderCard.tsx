"use client";
import { Button } from "@workspace/ui/components/button";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@workspace/ui/components/card";
import { useEffect, useState } from "react";
import type { TableOrder } from "../types/order";

// 時間関連の定数
const TIME_THRESHOLDS = {
  GREEN_MINUTES: 5,
  YELLOW_MINUTES: 10,
  RED_MINUTES: 15,
} as const;

const UPDATE_INTERVAL_SECONDS = 1; // 1秒間隔で更新

// 時間色分けの型定義
type TimeColorClass = "text-green-600" | "text-yellow-600" | "text-red-600";

export default function OrderCard({ order }: { order: TableOrder }) {
  // 現在時刻をnumber型で管理（パフォーマンス向上）
  const [currentTimeMs, setCurrentTimeMs] = useState<number>(Date.now());

  // リアルタイム時間更新
  useEffect(() => {
    const intervalId = setInterval(() => {
      setCurrentTimeMs(Date.now());
    }, UPDATE_INTERVAL_SECONDS * 1000);

    // クリーンアップ関数
    return () => clearInterval(intervalId);
  }, []); // 依存配列が空なので、マウント時に1回だけ実行

  // 経過時間（分）を計算
  const calculateElapsedMinutes = (): number => {
    return Math.floor((currentTimeMs - order.elapsedMinutes.getTime()) / 60000);
  };

  // 経過時間に基づく色分け
  const getTimeColorClass = (minutes: number): TimeColorClass => {
    if (minutes < TIME_THRESHOLDS.GREEN_MINUTES) {
      return "text-green-600";
    }
    if (minutes < TIME_THRESHOLDS.YELLOW_MINUTES) {
      return "text-yellow-600";
    }
    return "text-red-600";
  };

  const elapsedMinutes = calculateElapsedMinutes();
  const timeColorClass = getTimeColorClass(elapsedMinutes);
  const primaryDish = order.dishes[0];

  return (
    <Card className='flex flex-col'>
      <CardHeader>
        <CardTitle className='flex justify-between items-center'>
          <span>
            {primaryDish
              ? `${primaryDish.name} × ${primaryDish.qty}`
              : "メニュー未設定"}
          </span>
          <span className={`font-bold ${timeColorClass}`}>
            {elapsedMinutes}分
          </span>
        </CardTitle>
      </CardHeader>
      <CardContent>
        <div>H</div>
      </CardContent>
      <CardFooter className='mt-auto'>
        <Button className='w-full'>調理完了</Button>
      </CardFooter>
    </Card>
  );
}
