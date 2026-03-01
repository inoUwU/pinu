import useSWR from "swr";
import type { Table } from "@/app/(authenticated)/bill/types/table";
import { getAllTables } from "@/lib/api/tables/tables";

const SWR_KEYS = {
  allTables: "tables",
} as const;

/** リフレッシュ間隔: 5秒（テーブル状態のリアルタイム性を確保） */
const REFRESH_INTERVAL = 5000;

/**
 * 会計画面用のテーブル一覧を取得するフック
 * occupied + billing のテーブルのみ返す
 */
export const useBillTables = () => {
  const result = useSWR(SWR_KEYS.allTables, () => getAllTables(), {
    refreshInterval: REFRESH_INTERVAL,
    revalidateOnFocus: true,
    dedupingInterval: 2000,
  });

  const billTables: Table[] =
    result.data?.filter(
      t => t.status === "occupied" || t.status === "billing",
    ) ?? [];

  return {
    ...result,
    data: billTables,
    allTables: result.data,
  };
};
