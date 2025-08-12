"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import Link from "next/link";
import { useForm } from "react-hook-form";
import type { z } from "zod";
import { LoginFormSchema } from "@/app/login/schema";
import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";

export function LoginPage() {
  const form = useForm<z.infer<typeof LoginFormSchema>>({
    resolver: zodResolver(LoginFormSchema),
    defaultValues: {},
  });
  function onSubmit(data: z.infer<typeof LoginFormSchema>) {}

  return (
    <>
      <Form {...form}>
        <form
          onSubmit={form.handleSubmit(onSubmit)}
          className='flex flex-col items-center justify-center gap-5'
        >
          <h1 className='text-2xl font-bold'>Welcome back</h1>
          <FormField
            control={form.control}
            name='userId'
            render={({ field }) => (
              <FormItem className='w-10/12 lg:w-2/12'>
                <FormLabel>User ID</FormLabel>
                <FormControl>
                  <Input placeholder='userId' {...field} />
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
          <Button type='submit' className='w-10/12 lg:w-2/12'>
            Login
          </Button>
          <Link href='/register' className='text-blue-500 hover:underline'>
            Don't have an account? Sign up
          </Link>
        </form>
      </Form>
    </>
  );
}
export default LoginPage;
