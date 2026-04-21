"use client";

import { useSortable } from "@dnd-kit/sortable";
import { CSS } from "@dnd-kit/utilities";
import { Button } from "@workspace/ui/components/button";
import {
  Card,
  CardContent,
  CardHeader,
  CardTitle,
} from "@workspace/ui/components/card";
import { AlertCircle, DollarSign, Pencil } from "lucide-react";
import type { MenuItem } from "../types";

interface MenuCardProps {
  menuItem: MenuItem;
  onEdit?: (menuItem: MenuItem) => void;
}

/**
 * ドラッグ&ドロップ可能なメニューカード
 * Microsoft Planner風のカードデザイン
 */
export function MenuCard({ menuItem, onEdit }: MenuCardProps) {
  const {
    attributes,
    listeners,
    setNodeRef,
    transform,
    transition,
    isDragging,
  } = useSortable({ id: menuItem.id });

  const style = {
    transform: CSS.Transform.toString(transform),
    transition,
  };

  const handleEditClick = (e: React.MouseEvent) => {
    e.stopPropagation();
    onEdit?.(menuItem);
  };

  return (
    <Card
      ref={setNodeRef}
      style={style}
      className={`cursor-grab active:cursor-grabbing transition-shadow hover:shadow-md ${
        isDragging ? "opacity-50 shadow-lg" : ""
      } ${menuItem.isSoldOut ? "opacity-60" : ""}`}
      {...attributes}
      {...listeners}
    >
      <CardHeader className='pb-2'>
        <div className='flex items-start justify-between'>
          <CardTitle className='text-lg font-medium leading-tight pr-2'>
            {menuItem.name}
            {menuItem.isSoldOut && (
              <AlertCircle
                className='inline-block ml-1 h-4 w-4 text-red-500'
                aria-label='売り切れ'
                role='img'
              />
            )}
          </CardTitle>
          <Button
            variant='ghost'
            size='sm'
            className='h-6 w-6 p-0 shrink-0'
            onClick={handleEditClick}
            aria-label={`${menuItem.name}を編集`}
          >
            <Pencil className='h-3 w-3' aria-hidden='true' />
          </Button>
        </div>
      </CardHeader>

      <CardContent className='pt-0 space-y-2'>
        {menuItem.description && (
          <p className='text-sm text-muted-foreground line-clamp-2'>
            {menuItem.description}
          </p>
        )}

        <div className='flex items-center justify-between'>
          <div className='flex items-center text-sm font-medium'>
            <DollarSign className='h-4 w-4 mr-1' aria-hidden='true' />¥
            {menuItem.price.toLocaleString()}
          </div>

          {menuItem.options.length > 0 && (
            <div className='text-xs text-muted-foreground'>
              オプション {menuItem.options.length}個
            </div>
          )}
        </div>

        {menuItem.isSoldOut && (
          <div className='text-xs bg-red-100 text-red-800 px-2 py-1 rounded'>
            売り切れ
          </div>
        )}

        {menuItem.imageUrl && (
          <div className='mt-2'>
            {/* biome-ignore lint/performance/noImgElement: Using placeholder image */}
            <img
              src={menuItem.imageUrl}
              alt=''
              aria-hidden='true'
              className='w-full h-20 object-cover rounded-md'
              loading='lazy'
            />
          </div>
        )}
      </CardContent>
    </Card>
  );
}
