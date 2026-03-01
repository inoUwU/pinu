/** メニューアイテム */
export type Menu = {
  menu_id: string;
  name: string;
  price: number;
  category_id: string;
  image_url: string;
  description: string;
};

/** GET /api/menus レスポンス */
export type GetMenusResponse = {
  menus: Menu[];
  count: number;
};
