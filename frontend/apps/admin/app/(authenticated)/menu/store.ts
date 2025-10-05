"use client";

import { create } from "zustand";
import { devtools } from "zustand/middleware";
import { dummyCategories, dummyMenuItems } from "./dummyData";
import type { Category, MenuItem, MenuOption } from "./types";

/**
 * メニュー管理ストアの状態型定義
 */
interface MenuStore {
  // 状態
  categories: Category[];
  menuItems: MenuItem[];
  isLoading: boolean;
  selectedMenuId: string | null;

  // カテゴリー操作
  addCategory: (name: string) => void;
  deleteCategory: (categoryId: string) => void;
  updateCategory: (categoryId: string, updates: Partial<Category>) => void;
  reorderCategories: (categoryIds: string[]) => void;

  // メニューアイテム操作
  addMenuItem: (
    categoryId: string,
    menuData: Omit<MenuItem, "id" | "displayOrder">,
  ) => void;
  deleteMenuItem: (menuId: string) => void;
  updateMenuItem: (menuId: string, updates: Partial<MenuItem>) => void;
  moveMenuItemToCategory: (menuId: string, newCategoryId: string) => void;
  reorderMenuItemsInCategory: (categoryId: string, menuIds: string[]) => void;

  // メニューオプション操作
  addMenuOption: (menuId: string, option: Omit<MenuOption, "id">) => void;
  updateMenuOption: (
    menuId: string,
    optionId: string,
    updates: Partial<MenuOption>,
  ) => void;
  deleteMenuOption: (menuId: string, optionId: string) => void;

  // UI状態操作
  selectMenu: (menuId: string | null) => void;
  setLoading: (loading: boolean) => void;

  // 初期化
  initialize: () => void;
}

/**
 * メニュー管理のZustandストア
 * devtoolsで開発を支援
 */
export const useMenuStore = create<MenuStore>()(
  devtools(
    (set, _get) => ({
      // 初期状態
      categories: [],
      menuItems: [],
      isLoading: false,
      selectedMenuId: null,

      // カテゴリー操作
      addCategory: (name: string) => {
        set(state => {
          const newCategory: Category = {
            id: `cat-${Date.now()}`,
            name,
            displayOrder: state.categories.length + 1,
          };
          return {
            categories: [...state.categories, newCategory],
          };
        });
      },

      deleteCategory: (categoryId: string) => {
        set(state => ({
          categories: state.categories.filter(cat => cat.id !== categoryId),
          menuItems: state.menuItems.filter(
            item => item.categoryId !== categoryId,
          ),
        }));
      },

      updateCategory: (categoryId: string, updates: Partial<Category>) => {
        set(state => ({
          categories: state.categories.map(cat =>
            cat.id === categoryId ? { ...cat, ...updates } : cat,
          ),
        }));
      },

      reorderCategories: (categoryIds: string[]) => {
        set(state => {
          const reorderedCategories = categoryIds
            .map((id, index) => {
              const category = state.categories.find(cat => cat.id === id);
              return category ? { ...category, displayOrder: index + 1 } : null;
            })
            .filter(Boolean) as Category[];

          return { categories: reorderedCategories };
        });
      },

      // メニューアイテム操作
      addMenuItem: (
        categoryId: string,
        menuData: Omit<MenuItem, "id" | "displayOrder">,
      ) => {
        set(state => {
          const categoryItems = state.menuItems.filter(
            item => item.categoryId === categoryId,
          );
          const newMenuItem: MenuItem = {
            ...menuData,
            id: `menu-${Date.now()}`,
            categoryId,
            displayOrder: categoryItems.length + 1,
          };
          return {
            menuItems: [...state.menuItems, newMenuItem],
          };
        });
      },

      deleteMenuItem: (menuId: string) => {
        set(state => ({
          menuItems: state.menuItems.filter(item => item.id !== menuId),
        }));
      },

      updateMenuItem: (menuId: string, updates: Partial<MenuItem>) => {
        set(state => ({
          menuItems: state.menuItems.map(item =>
            item.id === menuId ? { ...item, ...updates } : item,
          ),
        }));
      },

      moveMenuItemToCategory: (menuId: string, newCategoryId: string) => {
        set(state => {
          const newCategoryItems = state.menuItems.filter(
            item => item.categoryId === newCategoryId,
          );

          return {
            menuItems: state.menuItems.map(item =>
              item.id === menuId
                ? {
                    ...item,
                    categoryId: newCategoryId,
                    displayOrder: newCategoryItems.length + 1,
                  }
                : item,
            ),
          };
        });
      },

      reorderMenuItemsInCategory: (categoryId: string, menuIds: string[]) => {
        set(state => ({
          menuItems: state.menuItems.map(item => {
            if (item.categoryId === categoryId) {
              const index = menuIds.indexOf(item.id);
              return index !== -1 ? { ...item, displayOrder: index + 1 } : item;
            }
            return item;
          }),
        }));
      },

      // メニューオプション操作
      addMenuOption: (menuId: string, option: Omit<MenuOption, "id">) => {
        set(state => ({
          menuItems: state.menuItems.map(item => {
            if (item.id === menuId) {
              const newOption: MenuOption = {
                ...option,
                id: `opt-${Date.now()}`,
              };
              return {
                ...item,
                options: [...item.options, newOption],
              };
            }
            return item;
          }),
        }));
      },

      updateMenuOption: (
        menuId: string,
        optionId: string,
        updates: Partial<MenuOption>,
      ) => {
        set(state => ({
          menuItems: state.menuItems.map(item => {
            if (item.id === menuId) {
              return {
                ...item,
                options: item.options.map(opt =>
                  opt.id === optionId ? { ...opt, ...updates } : opt,
                ),
              };
            }
            return item;
          }),
        }));
      },

      deleteMenuOption: (menuId: string, optionId: string) => {
        set(state => ({
          menuItems: state.menuItems.map(item => {
            if (item.id === menuId) {
              return {
                ...item,
                options: item.options.filter(opt => opt.id !== optionId),
              };
            }
            return item;
          }),
        }));
      },

      // UI状態操作
      selectMenu: (menuId: string | null) => {
        set({ selectedMenuId: menuId });
      },

      setLoading: (loading: boolean) => {
        set({ isLoading: loading });
      },

      // 初期化
      initialize: () => {
        set({
          categories: [...dummyCategories],
          menuItems: [...dummyMenuItems],
          isLoading: false,
          selectedMenuId: null,
        });
      },
    }),
    {
      name: "menu-store",
    },
  ),
);

/**
 * カテゴリー別のメニューアイテムを取得するセレクター
 */
export const useMenuItemsByCategory = (categoryId: string) => {
  return useMenuStore(state =>
    state.menuItems
      .filter(item => item.categoryId === categoryId)
      .sort((a, b) => a.displayOrder - b.displayOrder),
  );
};

/**
 * 選択されたメニューアイテムを取得するセレクター
 */
export const useSelectedMenuItem = () => {
  return useMenuStore(state => {
    if (!state.selectedMenuId) return null;
    return (
      state.menuItems.find(item => item.id === state.selectedMenuId) || null
    );
  });
};
