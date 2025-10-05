/**
 * メニュー管理システムの型定義
 */

export interface MenuOption {
  id: string;
  name: string;
  price: number;
  description?: string;
}

export interface MenuItem {
  id: string;
  name: string;
  description: string;
  price: number;
  imageUrl?: string;
  isSoldOut: boolean;
  categoryId: string;
  displayOrder: number;
  options: MenuOption[];
}

export interface Category {
  id: string;
  name: string;
  displayOrder: number;
  imageUrl?: string;
}
