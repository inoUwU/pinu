"use client";

import { useDroppable } from "@dnd-kit/core";
import {
  SortableContext,
  verticalListSortingStrategy,
} from "@dnd-kit/sortable";
import { Button } from "@workspace/ui/components/button";
import { Card, CardContent, CardHeader } from "@workspace/ui/components/card";
import { Plus, Trash2 } from "lucide-react";
import type { Category, MenuItem } from "../types";
import { MenuCard } from "./MenuCard";

interface CategoryColumnProps {
  category: Category;
  menuItems: MenuItem[];
  onDeleteCategory: (categoryId: string) => void;
  onEditMenu: (menuItem: MenuItem) => void;
  onAddMenu: (categoryId: string) => void;
}

/**
 * カテゴリー列コンポーネント
 * Microsoft Planner風のカラム表示で、ヘッダーにカテゴリー名と削除ボタン
 */
export function CategoryColumn({
  category,
  menuItems,
  onDeleteCategory,
  onEditMenu,
  onAddMenu,
}: CategoryColumnProps) {
  const { setNodeRef } = useDroppable({
    id: category.id,
  });

  const handleDeleteClick = () => {
    if (menuItems.length > 0) {
      const confirmDelete = confirm(
        `「${category.name}」カテゴリーを削除しますか？\nこのカテゴリーの${menuItems.length}個のメニューも削除されます。`,
      );
      if (!confirmDelete) return;
    }
    onDeleteCategory(category.id);
  };

  const handleAddMenuClick = () => {
    onAddMenu(category.id);
  };

  return (
    <div className='w-80 shrink-0'>
      <Card className='h-full'>
        <CardHeader className='pb-3'>
          <div className='flex items-center justify-between'>
            <div>
              <h3 className='font-semibold text-lg'>{category.name}</h3>
              <p className='text-sm text-muted-foreground'>
                {menuItems.length}個のメニュー
              </p>
            </div>
            <div className='flex gap-1'>
              <Button
                variant='ghost'
                size='sm'
                className='h-8 w-8 p-0'
                onClick={handleAddMenuClick}
                title='メニューを追加'
              >
                <Plus className='h-4 w-4' />
              </Button>
              <Button
                variant='ghost'
                size='sm'
                className='h-8 w-8 p-0 text-red-600 hover:text-red-700 hover:bg-red-50'
                onClick={handleDeleteClick}
                title='カテゴリーを削除'
              >
                <Trash2 className='h-4 w-4' />
              </Button>
            </div>
          </div>
        </CardHeader>

        <CardContent className='pt-0'>
          <div ref={setNodeRef} className='min-h-[200px] space-y-3'>
            <SortableContext
              items={menuItems.map(item => item.id)}
              strategy={verticalListSortingStrategy}
            >
              {menuItems.map(menuItem => (
                <MenuCard
                  key={menuItem.id}
                  menuItem={menuItem}
                  onEdit={onEditMenu}
                />
              ))}
            </SortableContext>

            {menuItems.length === 0 && (
              <div className='text-center py-8 text-muted-foreground'>
                <Plus className='h-8 w-8 mx-auto mb-2 opacity-50' />
                <p className='text-sm'>メニューがありません</p>
                <Button
                  variant='ghost'
                  size='sm'
                  className='mt-2'
                  onClick={handleAddMenuClick}
                >
                  メニューを追加
                </Button>
              </div>
            )}
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
