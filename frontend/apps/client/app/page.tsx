"use client";

import { worker } from "@workspace/mocks/browser";
import useSWR from "swr";

// if (process.env.NODE_ENV === "development") {
//   worker.start();
// }
//

worker.start();

export default function Page() {
  const fetcher = (...args: any) => fetch(...args).then(res => res.json());

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
      <div role='status' aria-live='polite' aria-busy='true' className='p-4'>
        <span>読み込み中...</span>
        <span className='sr-only'>
          データを読み込んでいます。しばらくお待ちください。
        </span>
      </div>
    );
  }

  console.log(data);

  return <div>hello {data[0].name}!</div>;
}
