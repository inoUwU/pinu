import { useState } from "react";
import { login } from "@/app/admin/(auth)/api/login";
import { persistAuthTokens } from "@/lib/auth/token-storage";

export function useLogin() {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<Error | null>(null);

  const doLogin = async (userId: string, password: string) => {
    setLoading(true);
    setError(null);
    try {
      const tokens = await login({ userId, password });
      persistAuthTokens(tokens);
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
