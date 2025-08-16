"use client";

import PageHeader from "@/app/components/PageHeader";
import { Button } from "@workspace/ui/components/button";

const CartSection = () => {
  return (
    <div className="h-full w-full flex flex-col items-center justify-start overflow-auto">
      <p className="mb-2">ここにカートの内容が表示されます。</p>
      {/* カートのコンテンツをここに追加 */}
    </div>
  );
};

const OrderButtonSection = () => {
  return (
    <div className="flex items-center justify-center h-full w-full">
      <Button className="w-3/4" variant="default">
        注文する
      </Button>
    </div>
  );
};

export default function OrderPage() {
  return (
    <div className="flex flex-col h-full">
      <div className="basis-1/12 flex-shrink-0 flex align-middle justify-center">
        <PageHeader title="注文" />
      </div>
      <div className="basis-10/12 grow overflow-auto bg-gray-50 pt-2">
        <CartSection />
      </div>
      <div className="basis-1/12 flex-shrink-0 flex align-middle justify-center bg-gray-50">
        <OrderButtonSection />
      </div>
    </div>
  );
}
