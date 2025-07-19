"use client";

import { Moon, Sun } from "lucide-react";
import Link from "next/link";
import { useTheme } from "next-themes";
import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Input } from "@/components/ui/input";

export default function LoginPage() {
  const { theme, setTheme } = useTheme();
  return (
    <>
      <main>
        <header className='flex justify-between p-4'>
          {/*RIGHT*/}
          <div className='flex items-center gap-4'>
            <span className='text-2xl font-bold'>Pinu</span>
          </div>
          {/*LEFT*/}
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button variant='outline' size='icon'>
                <Sun className='h-[1.2rem] w-[1.2rem] scale-100 rotate-0 transition-all dark:scale-0 dark:-rotate-90' />
                <Moon className='absolute h-[1.2rem] w-[1.2rem] scale-0 rotate-90 transition-all dark:scale-100 dark:rotate-0' />
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
        </header>
        <div className='flex flex-col items-center justify-center h-screen'>
          <h1>Welcome Back</h1>
          <Input placeholder='Email' className='mb-4 w-80' />
          <Input type='password' placeholder='Password' className='mb-4 w-80' />
          <Button className='w-80'>Login</Button>
          <Link
            href='/register'
            className='mt-4 text-sm text-blue-500 hover:underline'
          >
            Don't have an account? Register
          </Link>
        </div>
      </main>
    </>
  );
}
