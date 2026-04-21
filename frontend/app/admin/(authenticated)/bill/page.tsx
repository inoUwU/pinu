"use client";

import { Loader2, Receipt } from "lucide-react";
import { useCallback, useState } from "react";
import { useBillTables } from "@/hooks/useTables";
import BillSheet from "./components/BillSheet";
import TableCard from "./components/TableCard";
import type { Table } from "./types/table";

export default function BillPage() {
  const { data: tables, isLoading, mutate } = useBillTables();
  const [selectedTable, setSelectedTable] = useState<Table | null>(null);
  const [isSheetOpen, setIsSheetOpen] = useState(false);

  const handleSelectTable = useCallback((table: Table) => {
    setSelectedTable(table);
    setIsSheetOpen(true);
  }, []);

  const handleCheckoutComplete = useCallback(() => {
    setSelectedTable(null);
    mutate();
  }, [mutate]);

  return (
    <div className='space-y-6 p-6'>
      <div>
        <h1 className='text-2xl font-bold'>会計</h1>
        <p className='text-sm text-muted-foreground'>
          テーブルを選択して会計処理を行います
        </p>
      </div>

      {isLoading ? (
        <div className='flex items-center justify-center py-24'>
          <Loader2 className='size-8 animate-spin text-muted-foreground' />
        </div>
      ) : tables.length === 0 ? (
        <div className='flex flex-col items-center justify-center gap-4 py-24 text-muted-foreground'>
          <Receipt className='size-12' />
          <p className='text-lg'>会計待ちのテーブルはありません</p>
        </div>
      ) : (
        <div className='grid grid-cols-2 gap-4 md:grid-cols-3 lg:grid-cols-4'>
          {tables.map(table => (
            <TableCard
              key={table.table_id}
              table={table}
              isSelected={selectedTable?.table_id === table.table_id}
              onSelect={handleSelectTable}
            />
          ))}
        </div>
      )}

      <BillSheet
        table={selectedTable}
        open={isSheetOpen}
        onOpenChange={setIsSheetOpen}
        onCheckoutComplete={handleCheckoutComplete}
      />
    </div>
  );
}
