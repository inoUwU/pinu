"use client";

import CategoryCard from "@/app/(client)/category/components/CategoryCard";
import PageHeader from "@/app/(client)/components/PageHeader";

const CategorySection = () => {
  // TODO: Fetch categories from API

  return (
    <div className='h-full w-full p-4 flex flex-col gap-2 justify-start overflow-auto'>
      <CategoryCard
        key={1}
        category={{
          name: "飲み物",
          imageUrl: "https://placehold.jp/150x150.png",
          menuId: 1,
        }}
        showImage
      />
      <CategoryCard
        key={2}
        category={{
          name: "食べ物",
          imageUrl: "https://placehold.jp/150x150.png",
          menuId: 2,
        }}
        showImage
      />
    </div>
  );
};

export default function CategoryPage() {
  return (
    <div className='flex flex-col h-full'>
      <div className='w-full'>
        <PageHeader title='カテゴリー' />
      </div>
      <main
        className='basis-11/12 grow overflow-auto bg-gray-50 pt-2'
        id='main-content'
      >
        <CategorySection />
      </main>
    </div>
  );
}
