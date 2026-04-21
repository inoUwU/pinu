"use client";

import { Button } from "@workspace/ui/components/button";
import { Separator } from "@workspace/ui/components/separator";
import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetFooter,
  SheetHeader,
  SheetTitle,
} from "@workspace/ui/components/sheet";
import { Loader2 } from "lucide-react";
import { useMemo, useState } from "react";
import { toast } from "sonner";
import type { OrderGroup } from "@/app/admin/(authenticated)/bill/types/order";
import type { Table } from "@/app/admin/(authenticated)/bill/types/table";
import { useMenuMap } from "@/hooks/useMenuMap";
import { useTableOrders } from "@/hooks/useOrders";
import { checkoutTable } from "@/lib/api/tables/tables";

type Props = {
  table: Table | null;
  open: boolean;
  onOpenChange: (open: boolean) => void;
  /** 会計完了後のコールバック */
  onCheckoutComplete: () => void;
};

/** 注文グループ内の合計金額を計算 */
const calcGroupTotal = (group: OrderGroup): number => {
  return group.items.reduce((sum, item) => {
    if (item.status === "cancelled") return sum;
    return sum + item.price_at_order * item.quantity;
  }, 0);
};

/** 全注文グループの合計金額を計算 */
const calcTotalAmount = (groups: OrderGroup[]): number => {
  return groups.reduce((sum, group) => {
    if (group.status === "cancelled") return sum;
    return sum + calcGroupTotal(group);
  }, 0);
};

export default function BillSheet({
  table,
  open,
  onOpenChange,
  onCheckoutComplete,
}: Props) {
  const [isCheckingOut, setIsCheckingOut] = useState(false);

  const sessionId = table?.current_table_session_id ?? null;
  const { data: orderGroups, isLoading: isLoadingOrders } =
    useTableOrders(sessionId);
  const { menuMap, isLoading: isLoadingMenus } = useMenuMap();

  const activeGroups = useMemo(
    () => (orderGroups ?? []).filter(g => g.status !== "cancelled"),
    [orderGroups],
  );

  const totalAmount = useMemo(
    () => calcTotalAmount(activeGroups),
    [activeGroups],
  );

  /** メニュー名を解決する */
  const resolveMenuName = (menuId: string): string => {
    return menuMap.get(menuId)?.name ?? "不明なメニュー";
  };

  /** 会計処理 */
  const handleCheckout = async () => {
    if (!table) return;

    setIsCheckingOut(true);
    try {
      await checkoutTable(table.table_id);
      toast.success(`テーブル ${table.table_id} の会計が完了しました`);
      onOpenChange(false);
      onCheckoutComplete();
    } catch {
      toast.error("会計処理に失敗しました");
    } finally {
      setIsCheckingOut(false);
    }
  };

  const isLoading = isLoadingOrders || isLoadingMenus;

  return (
    <Sheet open={open} onOpenChange={onOpenChange}>
      <SheetContent side='right' className='flex flex-col sm:max-w-md'>
        <SheetHeader>
          <SheetTitle>テーブル {table?.table_id}</SheetTitle>
          <SheetDescription>注文内容の確認と会計処理</SheetDescription>
        </SheetHeader>

        {/* 注文一覧 */}
        <div className='flex-1 overflow-y-auto px-4'>
          {isLoading ? (
            <div className='flex items-center justify-center py-12'>
              <Loader2 className='size-6 animate-spin text-muted-foreground' />
            </div>
          ) : activeGroups.length === 0 ? (
            <p className='py-12 text-center text-sm text-muted-foreground'>
              注文がありません
            </p>
          ) : (
            <div className='space-y-4'>
              {activeGroups.map((group, groupIndex) => (
                <div key={group.orders_id}>
                  {groupIndex > 0 && <Separator className='mb-4' />}
                  <div className='mb-2 flex items-center justify-between'>
                    <span className='text-xs font-medium text-muted-foreground'>
                      注文 #{groupIndex + 1}
                    </span>
                    <span className='text-xs text-muted-foreground'>
                      {new Date(group.created_at).toLocaleTimeString("ja-JP")}
                    </span>
                  </div>
                  <ul className='space-y-2'>
                    {group.items
                      .filter(item => item.status !== "cancelled")
                      .map(item => (
                        <li
                          key={item.order_item_id}
                          className='flex items-center justify-between text-sm'
                        >
                          <span className='flex-1'>
                            {resolveMenuName(item.menu_id)}
                            {item.quantity > 1 && (
                              <span className='ml-1 text-muted-foreground'>
                                ×{item.quantity}
                              </span>
                            )}
                          </span>
                          <span className='font-medium tabular-nums'>
                            ¥
                            {(
                              item.price_at_order * item.quantity
                            ).toLocaleString()}
                          </span>
                        </li>
                      ))}
                  </ul>
                </div>
              ))}
            </div>
          )}
        </div>

        {/* フッター: 合計 + 会計ボタン */}
        <SheetFooter className='border-t pt-4'>
          <div className='flex items-center justify-between'>
            <span className='text-lg font-semibold'>合計</span>
            <span className='text-lg font-bold tabular-nums'>
              ¥{totalAmount.toLocaleString()}
            </span>
          </div>
          <Button
            size='lg'
            className='w-full'
            onClick={handleCheckout}
            disabled={isCheckingOut || isLoading || activeGroups.length === 0}
          >
            {isCheckingOut ? (
              <>
                <Loader2 className='mr-2 size-4 animate-spin' />
                処理中...
              </>
            ) : (
              "会計済みにする"
            )}
          </Button>
        </SheetFooter>
      </SheetContent>
    </Sheet>
  );
}
