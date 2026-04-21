import useSWR from "swr";
import type { Menu } from "@/app/admin/(authenticated)/bill/types/menu";
import { getAllMenus } from "@/lib/api/menus/menus";

const SWR_KEYS = {
  allMenus: "menus",
} as const;

/**
 * 全メニューを Map<menu_id, Menu> で取得するフック
 * メニュー名の解決に使用する
 */
export const useMenuMap = () => {
  const result = useSWR(SWR_KEYS.allMenus, () => getAllMenus(), {
    revalidateOnFocus: false,
    dedupingInterval: 60000,
  });

  const menuMap = new Map<string, Menu>();
  if (result.data) {
    for (const menu of result.data) {
      menuMap.set(menu.menu_id, menu);
    }
  }

  return {
    ...result,
    menuMap,
  };
};
