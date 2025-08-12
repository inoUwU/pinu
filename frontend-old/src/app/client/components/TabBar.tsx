import { JapaneseYen, ShoppingCart, Utensils } from "lucide-react";
import Link from "next/link";
import { ToggleGroup, ToggleGroupItem } from "@/components/ui/toggle-group";

export default function TabBar() {
  return (
    <ToggleGroup type='single' size='lg' className='w-lvw h-full p-0 m-0'>
      <ToggleGroupItem value='category'>
        <Utensils className='w-6 h-6' />
        <Link href='/client/category'>メニュー</Link>
      </ToggleGroupItem>
      <ToggleGroupItem value='order'>
        <ShoppingCart className='w-6 h-6' />
        <Link href='/client/order'>注文</Link>
      </ToggleGroupItem>
      <ToggleGroupItem value='history'>
        <JapaneseYen className='w-6 h-6' />
        <Link href='/client/history'>会計・履歴</Link>
      </ToggleGroupItem>
    </ToggleGroup>
  );
}
