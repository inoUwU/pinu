import Link from "next/link";
import { Menubar, MenubarMenu, MenubarTrigger } from "@/components/ui/menubar";

export default function TabBar() {
  return (
    <Menubar className='w-11/12 h-10'>
      <MenubarMenu>
        <MenubarTrigger className='flex-1 text-center min-w-[5rem]'>
          <Link href='/client/category' className='block w-full h-full'>
            カテゴリー
          </Link>
        </MenubarTrigger>
        <MenubarTrigger className='flex-1 text-center min-w-[5rem]'>
          <Link href='/client/order' className='block w-full h-full'>
            注文
          </Link>
        </MenubarTrigger>
        <MenubarTrigger className='flex-1 text-center min-w-[5rem]'>
          <Link href='/client/history' className='block w-full h-full'>
            会計・履歴
          </Link>
        </MenubarTrigger>
      </MenubarMenu>
    </Menubar>
  );
}
