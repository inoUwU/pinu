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

  if (error) return <div>failed to load</div>;
  if (isLoading) return <div>loading...</div>;
  console.log(data);

  return <div>hello {data[0].name}!</div>;
}
