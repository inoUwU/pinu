"use client";

import { worker } from "@workspace/mocks/browser";
import { useEffect } from "react";
import useSWR from "swr";

type User = {
  id: number;
  name: string;
};

export default function Page() {
  useEffect(() => {
    if (process.env.NODE_ENV === "development") {
      void worker.start();
    }
  }, []);

  const fetcher = async (input: RequestInfo | URL): Promise<User[]> => {
    const response = await fetch(input);
    return response.json();
  };

  const { data, error, isLoading } = useSWR("/api/users", fetcher);

  if (error) {
    return (
      <div role='alert' aria-live='assertive' className='p-4'>
        <p className='text-destructive'>
          データの読み込みに失敗しました。ページを再読み込みしてください。
        </p>
      </div>
    );
  }

  if (isLoading) {
    return (
      <output aria-live='polite' aria-busy='true' className='p-4'>
        <span>読み込み中...</span>
        <span className='sr-only'>
          データを読み込んでいます。しばらくお待ちください。
        </span>
      </output>
    );
  }

  const firstUser = data?.[0];

  if (!firstUser) {
    return <div className='p-4'>ユーザーデータがありません。</div>;
  }

  return <div>hello {firstUser.name}!</div>;
}
