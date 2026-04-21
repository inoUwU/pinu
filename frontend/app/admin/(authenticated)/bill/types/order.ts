/** 注文アイテムオプション */
export type OrderItemOption = {
  menu_option_id: string;
};

/** 注文アイテム */
export type OrderItem = {
  order_item_id: string;
  menu_id: string;
  quantity: number;
  price_at_order: number;
  status: "pending" | "preparing" | "served" | "cancelled";
  created_at: string;
  options: OrderItemOption[];
};

/** 注文グループ */
export type OrderGroup = {
  orders_id: string;
  table_session_id: string;
  status: "open" | "closed" | "cancelled";
  created_at: string;
  items: OrderItem[];
};

/** GET /api/orders レスポンス */
export type GetOrdersResponse = {
  order_groups: OrderGroup[];
};
