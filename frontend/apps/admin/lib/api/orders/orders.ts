import type {
  GetOrdersResponse,
  OrderGroup,
} from "@/app/(authenticated)/bill/types/order";
import api from "../api";

const ORDER_ENDPOINTS = {
  orders: "orders",
} as const;

/**
 * テーブルセッションに紐づく注文を取得する
 */
export const getOrdersByTableSession = async (
  tableSessionId: string,
): Promise<OrderGroup[]> => {
  const response: GetOrdersResponse = await api
    .get(ORDER_ENDPOINTS.orders, {
      searchParams: { table_session_id: tableSessionId },
    })
    .json();
  return response.order_groups;
};
