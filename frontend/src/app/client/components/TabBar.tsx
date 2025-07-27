import Link from "next/link";
import { ToggleGroup, ToggleGroupItem } from "@/components/ui/toggle-group";

export default function TabBar() {
  return (
    <ToggleGroup type='single' size='lg' className='w-lvw h-full p-0 m-0'>
      <ToggleGroupItem value='category'>
        <Link href='/client/category'>カテゴリー</Link>
      </ToggleGroupItem>
      <ToggleGroupItem value='order'>
        <Link href='/client/order'>注文</Link>
      </ToggleGroupItem>
      <ToggleGroupItem value='history'>
        <Link href='/client/history'>会計・履歴</Link>
      </ToggleGroupItem>
    </ToggleGroup>
  );
}
