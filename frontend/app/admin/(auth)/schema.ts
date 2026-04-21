import { z } from "zod";

export const LoginFormSchema = z.object({
  userId: z.string({ required_error: "User ID は必須です" }),
  password: z.string({ required_error: "Password は必須です" }),
});

export const RegisterFormSchema = z
  .object({
    userId: z.string({ required_error: "User ID は必須です" }),
    password: z.string({ required_error: "Password は必須です" }),
    confirmPassword: z.string({
      required_error: "Confirm Password は必須です",
    }),
  })
  .refine(data => data.password === data.confirmPassword, {
    message: "Passwords が一致しません",
  });
