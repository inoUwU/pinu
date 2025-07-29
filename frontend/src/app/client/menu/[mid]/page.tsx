"use client";
import PageHeader from "@/app/client/components/PageHeader";

interface MenuPageProps {
  params: {
    mid: string;
  };
}

export default function MenuPage({ params }: MenuPageProps) {
  const menuId = params.mid;

  return (
    <div className='flex flex-col h-full'>
      <div className='basis-1/12 flex-shrink-0 flex align-middle justify-center'>
        <PageHeader title='メニュー' showBackButton={true} />
      </div>
      <main className='basis-11/12 grow overflow-auto bg-gray-50 pt-2'>
        {/* メニューのコンテンツをここに追加 */}
        <div className='text-center'>
          <p className='text-xl mb-4'>
            メニューID: {menuId} の内容がここに表示されます。
          </p>
          {/* 実際のメニューデータを取得して表示するロジックをここに追加 */}
        </div>
      </main>
    </div>
  );
}
