"use client";

import { Button } from "@workspace/ui/components/button";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@workspace/ui/components/dialog";
import { useRouter } from "next/navigation";
import PageHeader from "@/app/(client)/components/PageHeader";
import { DataTable } from "@/app/(client)/history/datatable";
import { columns, payments } from "@/app/(client)/history/SampleData";
import { useCheckOutStore } from "@/store/sotre";

// 注文履歴部分のコンポーネント
const OrderHistorySection = () => {
  return (
    <div className='h-full w-full flex flex-col items-center justify-start overflow-auto'>
      <p className='mb-2'>ここに会計や履歴のコンテンツが表示されます。</p>
      <DataTable columns={columns} data={payments} />
    </div>
  );
};

// 会計ボタン部分のコンポーネント
const CheckoutButtonSection = () => {
  const router = useRouter();
  const checkOutState = useCheckOutStore();

  const handleCheckout = () => {
    checkOutState.setIsCheckOuted();
    router.push("/checkout");
  };

  return (
    <div className='flex items-center justify-center h-full w-full'>
      <Dialog>
        <DialogTrigger asChild>
          <Button variant='default' className='w-3/4'>
            会計する
          </Button>
        </DialogTrigger>
        <DialogContent className='sm:max-w-[350px] [&>button]:hidden'>
          <DialogHeader>
            <DialogTitle>会計確認</DialogTitle>
            <DialogDescription>
              注文内容を確認し、会計を行います。よろしいですか？
            </DialogDescription>
          </DialogHeader>
          <DialogFooter>
            <div className='flex justify-between w-full'>
              <Button
                variant='destructive'
                className='w-35'
                onClick={handleCheckout}
              >
                会計
              </Button>
              <DialogClose asChild>
                <Button variant='outline' className='w-35'>
                  キャンセル
                </Button>
              </DialogClose>
            </div>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </div>
  );
};

export default function HistoryPage() {
  return (
    <div className='flex flex-col h-full'>
      <div className='basis-1/12 flex-shrink-0 flex align-middle justify-center'>
        <PageHeader title='会計・履歴' />
      </div>
      <div className='basis-10/12 grow overflow-auto bg-gray-50 pt-2'>
        <OrderHistorySection />
      </div>
      <div className='basis-1/12 flex-shrink-0 flex align-middle justify-center bg-gray-50'>
        <CheckoutButtonSection />
      </div>
    </div>
  );
}
