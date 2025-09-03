import { Card } from "@workspace/ui/components/card";
import type { Table } from "@/app/(authenticated)/bill/types/table";
import BillDialog from "./BillDialog";

export default function TableCard({ table }: { table: Table }) {
  return (
    <Card>
      <h2>{table.name}</h2>
      <p>価格: {table.price}円</p>
      <p>数量: {table.quantity}</p>
      <BillDialog table={table.id} />
    </Card>
  );
}
