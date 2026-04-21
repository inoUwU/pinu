"use client";

import { Button } from "@workspace/ui/components/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@workspace/ui/components/dropdown-menu";
import { SidebarTrigger } from "@workspace/ui/components/sidebar";
import { Coins, Moon, Sun, UtensilsCrossed } from "lucide-react";
import { useTheme } from "next-themes";
import { useRouter } from "nextjs-toploader/app";

const Navbar = () => {
  const { setTheme } = useTheme();
  const router = useRouter();

  const handleOperationClick = () => {
    router.push("/admin/operation");
  };

  const handleBillClick = () => {
    router.push("/admin/bill");
  };

  return (
    <nav className='p-4 flex items-center justify-between'>
      {/*LEFT*/}
      <SidebarTrigger />
      {/*RIGHT*/}
      <div className='flex items-center gap-4'>
        <Button variant='outline' onClick={handleBillClick}>
          <Coins className='h-4 w-4 mr-2' aria-hidden='true' />
          会計
        </Button>
        <Button variant='outline' onClick={handleOperationClick}>
          <UtensilsCrossed className='h-4 w-4 mr-2' aria-hidden='true' />
          オペレーション
        </Button>

        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button variant='outline' size='icon' aria-label='テーマ切り替え'>
              <Sun
                className='h-[1.2rem] w-[1.2rem] scale-100 rotate-0 transition-all dark:scale-0 dark:-rotate-90'
                aria-hidden='true'
              />
              <Moon
                className='absolute h-[1.2rem] w-[1.2rem] scale-0 rotate-90 transition-all dark:scale-100 dark:rotate-0'
                aria-hidden='true'
              />
              <span className='sr-only'>Toggle theme</span>
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent align='end'>
            <DropdownMenuItem onClick={() => setTheme("light")}>
              Light
            </DropdownMenuItem>
            <DropdownMenuItem onClick={() => setTheme("dark")}>
              Dark
            </DropdownMenuItem>
            <DropdownMenuItem onClick={() => setTheme("system")}>
              System
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      </div>
    </nav>
  );
};

export default Navbar;
