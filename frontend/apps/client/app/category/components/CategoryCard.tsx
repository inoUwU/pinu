import { Card, CardContent, CardTitle } from "@workspace/ui/components/card";
import { useRouter } from "nextjs-toploader/app";

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

  return (
    <Card onClick={handleClick}>
      <CardContent>
        <div className='flex gap-2'>
          {showImage && category.imageUrl && (
            // biome-ignore lint/performance/noImgElement: Using placeholder image
            <img
              src={category.imageUrl}
              alt={category.name}
              className='w-24 h-24 object-cover rounded-md mb-2'
            />
          )}
          <CardTitle className='text-2xl break-all'>{category.name}</CardTitle>
        </div>
      </CardContent>
    </Card>
  );
}
