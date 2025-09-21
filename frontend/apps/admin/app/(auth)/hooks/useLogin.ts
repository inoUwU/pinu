import { useState } from "react";
import { login } from "@/app/(auth)/api/login";

export function useLogin() {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<Error | null>(null);

  const doLogin = async (userId: string, password: string) => {
    setLoading(true);
    setError(null);
    try {
      const tokens = await login({ userId, password });
      localStorage.setItem("access_token", tokens.access_token);
      localStorage.setItem("refresh_token", tokens.refresh_token);
      return tokens;
    } catch (err) {
      setError(err as Error);
      throw err;
    } finally {
      setLoading(false);
    }
  };

  return { doLogin, loading, error };
}
