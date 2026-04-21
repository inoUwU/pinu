import { SWRConfig } from "swr";
import { fetcher } from "../api/fetcher";

export const SWRProvider = ({ children }: { children: React.ReactNode }) => {
  return (
    <SWRConfig
      value={{
        fetcher,
        revalidateOnFocus: true,
        shouldRetryOnError: false,
        refreshInterval: 3000,
      }}
    >
      {children}
    </SWRConfig>
  );
};
