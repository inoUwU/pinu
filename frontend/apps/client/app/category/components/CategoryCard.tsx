import { Card, CardContent, CardTitle } from "@workspace/ui/components/card";
import { useRouter } from "nextjs-toploader/app";
import type { KeyboardEvent } from "react";

export type CategoryCardProps = {
  category: { name: string; menuId: number; imageUrl?: string };
  showImage?: boolean;
};

export default function CategoryCard({
  category,
  showImage = false,
}: CategoryCardProps) {
  const router = useRouter();

  const handleClick = () => {
    router.push(`/menu/${category.menuId}`);
  };

  const handleKeyDown = (e: KeyboardEvent<HTMLDivElement>) => {
    if (e.key === "Enter" || e.key === " ") {
      e.preventDefault();
      handleClick();
    }
  };

  return (
    <Card
      onClick={handleClick}
      onKeyDown={handleKeyDown}
      tabIndex={0}
      role='button'
      aria-label={`${category.name}カテゴリーを表示`}
      className='cursor-pointer focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2'
    >
      <CardContent>
        <div className='flex gap-2'>
          {showImage && category.imageUrl && (
            // biome-ignore lint/performance/noImgElement: Using placeholder image
            <img
              src={category.imageUrl}
              alt={`${category.name}のカテゴリー画像`}
              className='w-24 h-24 object-cover rounded-md mb-2'
            />
          )}
          <CardTitle className='text-2xl break-all'>{category.name}</CardTitle>
        </div>
      </CardContent>
    </Card>
  );
}
