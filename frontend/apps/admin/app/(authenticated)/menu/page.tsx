"use client";

import {
  DndContext,
  type DragEndEvent,
  DragOverlay,
  type DragStartEvent,
  PointerSensor,
  useSensor,
  useSensors,
} from "@dnd-kit/core";
import {
  arrayMove,
  horizontalListSortingStrategy,
  SortableContext,
} from "@dnd-kit/sortable";
import { useEffect, useState } from "react";
import { CategoryColumn } from "./components/CategoryColumn";
import { HeaderToolbar } from "./components/HeaderToolbar";
import { MenuCard } from "./components/MenuCard";
import { useMenuStore } from "./store";
import type { MenuItem } from "./types";

/**
 * メニュー管理メインページ
 * Microsoft Planner風のカードベースUI
 */
export default function MenuPage() {
  const {
    categories,
    menuItems,
    addCategory,
    deleteCategory,
    moveMenuItemToCategory,
    reorderMenuItemsInCategory,
    selectMenu,
    initialize,
  } = useMenuStore();

  const [activeMenuItem, setActiveMenuItem] = useState<MenuItem | null>(null);
  const [searchValue, setSearchValue] = useState("");

  // センサー設定（ドラッグ感度）
  const sensors = useSensors(
    useSensor(PointerSensor, {
      activationConstraint: {
        distance: 8,
      },
    }),
  );

  // 初期化
  useEffect(() => {
    initialize();
  }, [initialize]);

  const handleDragStart = (event: DragStartEvent) => {
    const { active } = event;
    const menuItem = menuItems.find(item => item.id === active.id);

    if (menuItem) {
      setActiveMenuItem(menuItem);
    }
  };

  const handleDragEnd = (event: DragEndEvent) => {
    const { active, over } = event;
    setActiveMenuItem(null);

    if (!over) return;

    const activeId = active.id as string;
    const overId = over.id as string;

    // カテゴリー間の移動かチェック
    const draggedMenuItem = menuItems.find(item => item.id === activeId);
    const targetCategory = categories.find(cat => cat.id === overId);

    if (draggedMenuItem && targetCategory) {
      // カテゴリー間移動
      if (draggedMenuItem.categoryId !== targetCategory.id) {
        moveMenuItemToCategory(activeId, targetCategory.id);
        return;
      }
    }

    // 同じカテゴリー内での並び替え
    const sourceCategory = categories.find(cat => {
      const items = menuItems.filter(item => item.categoryId === cat.id);
      return items.some(item => item.id === activeId);
    });

    if (sourceCategory) {
      const categoryItems = menuItems
        .filter(item => item.categoryId === sourceCategory.id)
        .sort((a, b) => a.displayOrder - b.displayOrder);

      const oldIndex = categoryItems.findIndex(item => item.id === activeId);
      const newIndex = categoryItems.findIndex(item => item.id === overId);

      if (oldIndex !== -1 && newIndex !== -1 && oldIndex !== newIndex) {
        const reorderedItems = arrayMove(categoryItems, oldIndex, newIndex);
        reorderMenuItemsInCategory(
          sourceCategory.id,
          reorderedItems.map(item => item.id),
        );
      }
    }
  };

  const handleEditMenu = (menuItem: MenuItem) => {
    selectMenu(menuItem.id);
    // TODO: メニューモーダルを開く
  };

  const handleAddMenu = (categoryId: string) => {
    // TODO: 新しいメニュー追加モーダルを開く
    console.log("Add menu to category:", categoryId);
  };

  // 検索フィルタリング
  const filteredCategories = categories.map(category => {
    const categoryItems = menuItems
      .filter(item => item.categoryId === category.id)
      .filter(
        item =>
          searchValue === "" ||
          item.name.toLowerCase().includes(searchValue.toLowerCase()) ||
          item.description.toLowerCase().includes(searchValue.toLowerCase()),
      )
      .sort((a, b) => a.displayOrder - b.displayOrder);

    return {
      category,
      items: categoryItems,
    };
  });

  return (
    <div className='h-full flex flex-col bg-gray-50'>
      <HeaderToolbar
        onAddCategory={addCategory}
        onRefresh={initialize}
        searchValue={searchValue}
        onSearchChange={setSearchValue}
      />

      <div className='flex-1 overflow-hidden'>
        <DndContext
          sensors={sensors}
          onDragStart={handleDragStart}
          onDragEnd={handleDragEnd}
        >
          <div className='h-full overflow-x-auto overflow-y-hidden'>
            <div className='flex gap-6 p-6 h-full min-w-max'>
              <SortableContext
                items={categories.map(cat => cat.id)}
                strategy={horizontalListSortingStrategy}
              >
                {filteredCategories.map(({ category, items }) => (
                  <CategoryColumn
                    key={category.id}
                    category={category}
                    menuItems={items}
                    onDeleteCategory={deleteCategory}
                    onEditMenu={handleEditMenu}
                    onAddMenu={handleAddMenu}
                  />
                ))}
              </SortableContext>

              {categories.length === 0 && (
                <div className='flex items-center justify-center w-full h-full text-muted-foreground'>
                  <div className='text-center'>
                    <p className='text-lg mb-2'>カテゴリーがありません</p>
                    <p className='text-sm'>
                      「カテゴリー追加」ボタンから新しいカテゴリーを作成してください
                    </p>
                  </div>
                </div>
              )}
            </div>
          </div>

          <DragOverlay>
            {activeMenuItem && (
              <MenuCard menuItem={activeMenuItem} onEdit={() => {}} />
            )}
          </DragOverlay>
        </DndContext>
      </div>
    </div>
  );
}
