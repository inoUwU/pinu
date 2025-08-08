"use client";

import useSWRSubscription from "swr/subscription";

function MyComponent() {
  const { data, error } = useSWRSubscription(
    "http://localhost:8000/api/sse",
    (key, { next }) => {
      const eventSource = new EventSource(key);

      eventSource.onmessage = event => {
        console.log("Received SSE message:", event.data);
        next(null, JSON.parse(event.data));
      };

      eventSource.onerror = error => {
        next(error);
      };

      return () => eventSource.close();
    },
  );

  if (error) return <div>エラーが発生しました: {error.message}</div>;
  if (!data) return <div>SSE接続中...</div>;

  return <div>{JSON.stringify(data)}</div>;
}

export default MyComponent;
