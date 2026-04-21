"use client";

import type { ColumnDef } from "@tanstack/react-table";
import { Button } from "@workspace/ui/components/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuTrigger,
} from "@workspace/ui/components/dropdown-menu";
import { MoreHorizontal } from "lucide-react";
import type { UserFromData } from "@/app/admin/(authenticated)/user/types/table";

interface ColumnActions {
  onEdit?: (data: UserFromData) => void;
  onDelete?: (id: string) => void;
}

export const createColumns = (): ColumnDef<UserFromData>[] => {
  const columns: ColumnDef<UserFromData>[] = [
    {
      accessorKey: "name_4603829743",
      header: "Name",
    },
    {
      accessorKey: "name_0878515932",
      header: "Age",
    },
    {
      accessorKey: "name_0706064476",
      header: "Gender",
    },
    {
      accessorKey: "name_6646786819",
      header: "Email",
    },
  ];

  columns.push({
    id: "actions",
    cell: ({ row, table }) => {
      const record = row.original;
      const { onEdit, onDelete } = table.options.meta as ColumnActions;

      return (
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button variant='ghost' className='h-8 w-8 p-0'>
              <span className='sr-only'>Open menu</span>
              <MoreHorizontal className='h-4 w-4' />
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent align='end'>
            <DropdownMenuLabel>Actions</DropdownMenuLabel>

            {onEdit && (
              <DropdownMenuItem onClick={() => onEdit(record)}>
                Edit
              </DropdownMenuItem>
            )}

            {onDelete && (
              <DropdownMenuItem onClick={() => onDelete(record.id)}>
                Delete
              </DropdownMenuItem>
            )}
          </DropdownMenuContent>
        </DropdownMenu>
      );
    },
  });

  return columns;
};
