"use client";

import { useState } from "react";
import useSWRSubscription from "swr/subscription";
import OrderList from "./components/OrderList";
import type { TableOrder } from "./types/order";

const SSE_ENDPOINT = "http://localhost:8000/api/sse";

export default function OperationPage() {
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

  const [orders, setOrders] = useState([
    {
      id: 1,
      tableNumber: 2,
      elapsedMinutes: new Date(Date.now() - 3 * 60000), // 3分前
      dishes: [{ name: "餃子", qty: 3 }],
      status: "pending",
      completedAt: undefined,
    } satisfies TableOrder,
    {
      id: 1,
      tableNumber: 1,
      elapsedMinutes: new Date(Date.now() - 3 * 60000), // 3分前
      dishes: [{ name: "餃子", qty: 3 }],
      status: "pending",
      completedAt: undefined,
    } satisfies TableOrder,
    {
      id: 1,
      tableNumber: 1,
      elapsedMinutes: new Date(Date.now() - 5 * 60000), // 3分前
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
          {error ? (
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
