"use client";

import { Card } from "@workspace/ui/components/card";
import { cn } from "@workspace/ui/lib/utils";
import type {
  Table,
  TableStatus,
} from "@/app/admin/(authenticated)/bill/types/table";

type Props = {
  table: Table;
  isSelected: boolean;
  onSelect: (table: Table) => void;
};

/** ステータスに応じたバッジのスタイルとラベル */
const STATUS_CONFIG: Record<TableStatus, { label: string; className: string }> =
  {
    occupied: {
      label: "利用中",
      className: "bg-blue-100 text-blue-800",
    },
    billing: {
      label: "会計中",
      className: "bg-amber-100 text-amber-800",
    },
    available: {
      label: "空席",
      className: "bg-gray-100 text-gray-600",
    },
  };

export default function TableCard({ table, isSelected, onSelect }: Props) {
  const statusConfig = STATUS_CONFIG[table.status];

  return (
    <Card
      className={cn(
        "cursor-pointer p-4 transition-all hover:shadow-md",
        isSelected && "ring-2 ring-primary",
      )}
      onClick={() => onSelect(table)}
    >
      <div className='flex items-start justify-between'>
        <h3 className='text-lg font-semibold'>テーブル {table.table_id}</h3>
        <span
          className={cn(
            "inline-flex items-center rounded-full px-2 py-0.5 text-xs font-medium",
            statusConfig.className,
          )}
        >
          {statusConfig.label}
        </span>
      </div>
    </Card>
  );
}
