"use client";
import PageHeader from "@/app/components/PageHeader";

/**
 * 指定されたカテゴリー内のメニュー一覧を表示するコンポーネント
 */
// カテゴリーメニューリストセクションのprops型定義
type CategoryMenuListSectionProps = {
  id: number;
};

/**
 * 指定されたカテゴリー内のメニュー一覧を表示するコンポーネント
 */
const CategoryMenuListSection = ({ id }: CategoryMenuListSectionProps) => {
  // カテゴリーIDに基づいてメニュー一覧を取得するロジックをここに追加
  return (
    <div className='text-center'>
      <p className='text-xl mb-4'>
        カテゴリー内のメニュー一覧がここに表示されます。
      </p>
      {/* 実際のメニュー一覧を取得して表示するロジックをここに追加 */}
    </div>
  );
};

export default function CategoryDetailsPage({
  params,
}: {
  params: { categoryid: number };
}) {
  const categoryId: number = params.categoryid;

  return (
    <div className='flex flex-col h-full'>
      <div className='basis-1/12 flex-shrink-0 flex align-middle justify-center'>
        <PageHeader title='メニュー' showBackButton={true} />
      </div>
      <main className='basis-11/12 grow overflow-auto bg-gray-50 pt-2'>
        {/* メニューのコンテンツをここに追加 */}
        <div className='text-center'>
          <CategoryMenuListSection id={categoryId} />
        </div>
      </main>
    </div>
  );
}
