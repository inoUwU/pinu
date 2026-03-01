import type {
  GetMenusResponse,
  Menu,
} from "@/app/(authenticated)/bill/types/menu";
import api from "../api";

const MENU_ENDPOINTS = {
  menus: "menus",
} as const;

/**
 * 全メニューを取得する
 */
export const getAllMenus = async (): Promise<Menu[]> => {
  const response: GetMenusResponse = await api.get(MENU_ENDPOINTS.menus).json();
  return response.menus;
};
