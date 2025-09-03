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
import { LoginFormSchema } from "@/app/(auth)/schema";

export default function LoginPage() {
  const form = useForm<z.infer<typeof LoginFormSchema>>({
    resolver: zodResolver(LoginFormSchema),
    defaultValues: {},
  });
  function onSubmit(data: z.infer<typeof LoginFormSchema>) {
    console.log(data);
  }

  return (
    <div>
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
    </div>
  );
}
