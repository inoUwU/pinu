"use client";

import PageHeader from "@/app/client/components/PageHeader";

// メニューバーの高さを定数で管理
const MENU_BAR_HEIGHT_PX = 64; // 例: 64px。実際の高さに合わせて調整してください
const CHECKOUT_BUTTON_HEIGHT_PX = 64; // 会計ボタン領域の高さ

// 注文履歴部分のコンポーネント
function OrderHistorySection() {
  // TODO: 実際の注文履歴データをここに表示
  return (
    <div className='h-full w-full flex items-center justify-center overflow-auto'>
      <p>ここに会計や履歴のコンテンツが表示されます。</p>
    </div>
  );
}

// 会計ボタン部分のコンポーネント
function CheckoutButtonSection() {
  // TODO: 会計処理の実装
  return (
    <div className='flex items-center justify-center h-full'>
      <button
        type='button'
        className='px-8 py-3 bg-blue-600 text-white rounded shadow hover:bg-blue-700 transition'
      >
        会計する
      </button>
    </div>
  );
}

export default function HistoryPage() {
  return (
    <div className='flex flex-col h-full'>
      <div className='basis-1/12 flex-shrink-0 flex align-middle justify-center'>
        <PageHeader title='会計・履歴' />
      </div>
      <div className='basis-10/12 grow overflow-auto bg-gray-50'>
        <OrderHistorySection />
      </div>
      <div className='basis-1/12 flex-shrink-0 flex align-middle justify-center bg-gray-50'>
        <CheckoutButtonSection />
      </div>
    </div>
  );
}
