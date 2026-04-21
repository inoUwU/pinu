import useSWR from "swr";
import { getOrdersByTableSession } from "@/lib/api/orders/orders";

/**
 * テーブルセッションに紐づく注文を取得するフック
 * tableSessionId が null の場合はフェッチしない
 */
export const useTableOrders = (tableSessionId: string | null) => {
  return useSWR(
    tableSessionId ? ["orders", tableSessionId] : null,
    () => {
      if (!tableSessionId) return [];
      return getOrdersByTableSession(tableSessionId);
    },
    {
      revalidateOnFocus: true,
      dedupingInterval: 2000,
    },
  );
};
