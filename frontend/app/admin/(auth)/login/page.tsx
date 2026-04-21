"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { Button } from "@workspace/ui/components/button";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@workspace/ui/components/form";
import { Input } from "@workspace/ui/components/input";
import Link from "next/link";
import { useForm } from "react-hook-form";
import type { z } from "zod";
import { useLogin } from "@/app/admin/(auth)/hooks/useLogin";
import { LoginFormSchema } from "@/app/admin/(auth)/schema";

export default function LoginPage() {
  const { doLogin, loading, error } = useLogin();

  const form = useForm<z.infer<typeof LoginFormSchema>>({
    resolver: zodResolver(LoginFormSchema),
    defaultValues: {},
  });

  async function onSubmit(data: z.infer<typeof LoginFormSchema>) {
    const { userId, password } = data;
    await doLogin(userId, password);
  }

  return (
    <div>
      <Form {...form}>
        <form
          onSubmit={form.handleSubmit(onSubmit)}
          className='flex flex-col items-center justify-center min-h-[calc(100vh-10rem)]' // 10rem = h-40
          style={{ minHeight: "calc(100vh - 10rem)" }} // ヘッダー分の高さを引く
        >
          <h1 className='text-4xl font-bold mb-2'>Welcome back</h1>
          <div className='mb-4 w-screen flex flex-col items-center gap-4'>
            <FormField
              control={form.control}
              name='userId'
              render={({ field }) => (
                <FormItem className='w-10/12 lg:w-2/12'>
                  <FormLabel>User ID</FormLabel>
                  <FormControl>
                    <Input placeholder='userId' {...field} className='w-auto' />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name='password'
              render={({ field }) => (
                <FormItem className='w-10/12 lg:w-2/12'>
                  <FormLabel>Password</FormLabel>
                  <FormControl>
                    <Input placeholder='password' {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <div className='flex flex-col items-center gap-2  w-full'>
              <Button
                type='submit'
                className='w-10/12 lg:w-2/12'
                disabled={loading}
              >
                Login
              </Button>
              {error && (
                <p className='text-sm text-destructive'>{error.message}</p>
              )}
              {/* アカウント未登録の場合のリンク */}
              <Link
                href='/admin/register'
                className='text-blue-500 hover:underline'
              >
                アカウントをお持ちでないですか？新規登録
              </Link>
            </div>
          </div>
        </form>
      </Form>
    </div>
  );
}
