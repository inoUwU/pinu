import type { TableOrder } from "../types/order";
import OrderCard from "./OrderCard";

export default function OrderList({ orders }: { orders: TableOrder[] }) {
  // 1. テーブルごとにグループ化
  const grouped = orders.reduce<Record<number, TableOrder[]>>((acc, order) => {
    const key = order.tableNumber;
    if (!acc[key]) acc[key] = [];
    acc[key].push(order);
    return acc;
  }, {});

  // 2. 並び替え（テーブル番号順）
  const tableNumbers = Object.keys(grouped)
    .map(Number)
    .sort((a, b) => a - b);

  // テーブルを横並びに配置
  return (
    <div className='flex  gap-4 '>
      {tableNumbers.map(tableNum => {
        const tableOrders = grouped[tableNum] ?? [];

        // Talbe毎の列を作成する
        // Table毎の注文を取得し、経過時間でソート
        return (
          <div key={tableNum} className=' w-auto'>
            {tableOrders.map(order => (
              <OrderCard key={order.id} order={order} />
            ))}
          </div>
        );
      })}
    </div>
  );
}
