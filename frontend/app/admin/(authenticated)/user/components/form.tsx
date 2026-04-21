"use client";
import { zodResolver } from "@hookform/resolvers/zod";
import { Button } from "@workspace/ui/components/button";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@workspace/ui/components/form";
import { Input } from "@workspace/ui/components/input";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@workspace/ui/components/select";
import { useForm } from "react-hook-form";
import * as z from "zod";
import type { UserFromData } from "@/app/admin/(authenticated)/user/types/table";

const formSchema = z.object({
  name_4603829743: z.string(),
  name_0878515932: z.number(),
  name_0706064476: z.string().optional(),
  name_6646786819: z.string(),
});

// const formSchema = z.object({
//   id: z.string().optional(),
//   name_4603829743: z
//     .string()
//     .min(2, { message: "Name must be at least 2 characters" }),
//   name_0878515932: z.coerce
//     .number()
//     .min(0, { message: "Age must be a positive number" }),
//   name_0706064476: z.string().optional(),
//   name_6646786819: z.string().email({ message: "Invalid email address" }),
// });

interface UserFormProps {
  onSubmit: (data: UserFromData) => void;
  initialData?: UserFromData | null;
}

export default function UserForm({ onSubmit, initialData }: UserFormProps) {
  const defaultValues = initialData
    ? {
        name_4603829743: initialData.name,
        name_0878515932: 0,
        name_0706064476: initialData.isAdmin ? "admin" : "staff",
        name_6646786819: "",
      }
    : undefined;

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues,
  });

  //   function onSubmit(values: z.infer<typeof formSchema>) {
  //     try {
  //       console.log(values);
  //       toast(
  //         <pre className='mt-2 w-[340px] rounded-md bg-slate-950 p-4'>
  //           <code className='text-white'>{JSON.stringify(values, null, 2)}</code>
  //         </pre>,
  //       );
  //     } catch (error) {
  //       console.error("Form submission error", error);
  //       toast.error("Failed to submit the form. Please try again.");
  //     }
  //   }

  return (
    <Form {...form}>
      <form className='space-y-8 max-w-3xl mx-auto py-10'>
        <div className='grid grid-cols-12 gap-4'>
          <div className='col-span-6'>
            <FormField
              control={form.control}
              name='name_4603829743'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Name</FormLabel>
                  <FormControl>
                    <Input placeholder='shadcn' type='text' {...field} />
                  </FormControl>
                  <FormDescription>
                    This is your public display name.
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />
          </div>

          <div className='col-span-6'>
            <FormField
              control={form.control}
              name='name_0878515932'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Age</FormLabel>
                  <FormControl>
                    <Input placeholder='shadcn' type='number' {...field} />
                  </FormControl>
                  <FormDescription>
                    This is your public display name.
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />
          </div>
        </div>

        <div className='grid grid-cols-12 gap-4'>
          <div className='col-span-6'>
            <FormField
              control={form.control}
              name='name_0706064476'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Gender</FormLabel>
                  <Select
                    onValueChange={field.onChange}
                    defaultValue={field.value}
                  >
                    <FormControl>
                      <SelectTrigger>
                        <SelectValue placeholder='Select a verified email to display' />
                      </SelectTrigger>
                    </FormControl>
                    <SelectContent>
                      <SelectItem value='m@example.com'>
                        m@example.com
                      </SelectItem>
                      <SelectItem value='m@google.com'>m@google.com</SelectItem>
                      <SelectItem value='m@support.com'>
                        m@support.com
                      </SelectItem>
                    </SelectContent>
                  </Select>
                  <FormDescription>
                    You can manage email addresses in your email settings.
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />
          </div>

          <div className='col-span-6'>
            <FormField
              control={form.control}
              name='name_6646786819'
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Email</FormLabel>
                  <FormControl>
                    <Input placeholder='shadcn' type='email' {...field} />
                  </FormControl>
                  <FormDescription>
                    This is your public display name.
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />
          </div>
        </div>
        <Button type='submit'>Submit</Button>
      </form>
    </Form>
  );
}
