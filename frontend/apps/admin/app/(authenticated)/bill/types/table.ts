/** テーブルステータス */
export type TableStatus = "available" | "occupied" | "billing";

/** テーブル */
export type Table = {
  table_id: string;
  status: TableStatus;
  current_table_session_id: string | null;
  last_updated: string;
};

/** GET /api/tables レスポンス */
export type GetTablesResponse = {
  tables: Table[];
};

/** POST /api/tables/:id/checkout レスポンス */
export type CheckoutTableResponse = {
  table: Table;
  message: string;
};
