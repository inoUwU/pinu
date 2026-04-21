"use client";

import { useState } from "react";
import useSWRSubscription from "swr/subscription";
import OrderList from "./components/OrderList";
import type { TableOrder } from "./types/order";

const MINUTE_IN_MS = 60000;
const SSE_ENDPOINT = process.env.NEXT_PUBLIC_API_BASE_URL
  ? `${process.env.NEXT_PUBLIC_API_BASE_URL.replace(/\/$/, "")}/sse`
  : null;

export default function OperationPage() {
  const hasSseEndpoint = Boolean(SSE_ENDPOINT);
  const { data, error } = useSWRSubscription(SSE_ENDPOINT, (key, { next }) => {
    console.log("SSE接続を開始:", key);
    const eventSource = new EventSource(key);

    eventSource.onopen = () => {
      console.log("SSE接続が確立されました");
    };

    eventSource.onmessage = event => {
      console.log("Received SSE message:", event.data);
      try {
        const parsedData = JSON.parse(event.data);
        next(null, parsedData);
      } catch (parseError) {
        console.error("JSON解析エラー:", parseError);
        next(parseError);
      }
    };

    eventSource.onerror = error => {
      console.error("SSEエラー:", error);
      next(error);
    };

    return () => {
      console.log("SSE接続をクリーンアップ");
      eventSource.close();
    };
  });

  const [orders] = useState([
    {
      id: 1,
      tableNumber: 2,
      elapsedMinutes: new Date(Date.now() - 3 * MINUTE_IN_MS),
      dishes: [{ name: "餃子", qty: 3 }],
      status: "pending",
      completedAt: undefined,
    } satisfies TableOrder,
    {
      id: 1,
      tableNumber: 1,
      elapsedMinutes: new Date(Date.now() - 3 * MINUTE_IN_MS),
      dishes: [{ name: "餃子", qty: 3 }],
      status: "pending",
      completedAt: undefined,
    } satisfies TableOrder,
    {
      id: 1,
      tableNumber: 1,
      elapsedMinutes: new Date(Date.now() - 5 * MINUTE_IN_MS),
      dishes: [{ name: "チャーハン", qty: 3 }],
      status: "pending",
      completedAt: undefined,
    } satisfies TableOrder,
  ]);

  return (
    <div>
      <div className='mb-4'>
        <h2>厨房オーダー管理</h2>
        <div>
          <strong>接続状態:</strong>
          {!hasSseEndpoint ? (
            <span style={{ color: "red" }}>
              NEXT_PUBLIC_API_BASE_URL が設定されていません
            </span>
          ) : error ? (
            <span style={{ color: "red" }}>エラー: {error.message}</span>
          ) : data ? (
            <span style={{ color: "green" }}>
              接続中 - 最新データ: {JSON.stringify(data)}
            </span>
          ) : (
            <span style={{ color: "orange" }}>SSE接続中...</span>
          )}
        </div>
      </div>
      <main className='flex-1 overflow-y-auto p-4'>
        <OrderList orders={orders} />
      </main>
    </div>
  );
}
