import { z } from "zod";

export const LoginFormSchema = z.object({
  userId: z.string({ error: "User ID is required" }),
  password: z.string({ error: "Password is required" }),
});
