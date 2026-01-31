import {
  ToggleGroup,
  ToggleGroupItem,
} from "@workspace/ui/components/toggle-group";
import { JapaneseYen, ShoppingCart, Utensils } from "lucide-react";
import Link from "next/link";

export default function TabBar() {
  return (
    <nav aria-label='メインナビゲーション' className='w-lvw h-full'>
      <ToggleGroup type='single' size='lg' className='w-lvw h-full p-0 m-0'>
        <ToggleGroupItem value='category'>
          <Utensils className='w-6 h-6' aria-hidden='true' />
          <Link href='/category'>メニュー</Link>
        </ToggleGroupItem>
        <ToggleGroupItem value='order'>
          <ShoppingCart className='w-6 h-6' aria-hidden='true' />
          <Link href='/order'>注文</Link>
        </ToggleGroupItem>
        <ToggleGroupItem value='history'>
          <JapaneseYen className='w-6 h-6' aria-hidden='true' />
          <Link href='/history'>会計・履歴</Link>
        </ToggleGroupItem>
      </ToggleGroup>
    </nav>
  );
}
