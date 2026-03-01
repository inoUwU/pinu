import type {
  CheckoutTableResponse,
  GetTablesResponse,
  Table,
} from "@/app/(authenticated)/bill/types/table";
import api from "../api";

const TABLE_ENDPOINTS = {
  tables: "tables",
  checkout: (id: string) => `tables/${id}/checkout`,
} as const;

/**
 * 全テーブルを取得する
 */
export const getAllTables = async (): Promise<Table[]> => {
  const response: GetTablesResponse = await api
    .get(TABLE_ENDPOINTS.tables)
    .json();
  return response.tables;
};

/**
 * テーブルの会計処理を行う
 */
export const checkoutTable = async (
  tableId: string,
): Promise<CheckoutTableResponse> => {
  return api.post(TABLE_ENDPOINTS.checkout(tableId)).json();
};
