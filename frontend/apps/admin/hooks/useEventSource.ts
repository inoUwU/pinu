import { useEffect, useState } from "react";

/**
 * 指定URLに対してSSE接続し、受信メッセージを配列で返す
 * @param url SSEエンドポイントのURL
 */
export function useEventSource(url: string): string[] {
  const [messages, setMessages] = useState<string[]>([]);

  useEffect(() => {
    if (typeof window === "undefined" || !window.EventSource) {
      console.error("SSE非対応ブラウザ");
      return;
    }
    const eventSource = new EventSource(url);

    // メッセージ受信時の処理
    eventSource.onmessage = (e: MessageEvent) => {
      setMessages((prev) => [...prev, e.data]);
    };

    // エラー発生時はログ出力して接続をクローズ
    eventSource.onerror = (err) => {
      console.error("SSE接続エラー", err);
      eventSource.close();
    };

    // コンポーネントアンマウント時に接続をクローズ
    return () => {
      eventSource.close();
    };
  }, [url]);

  return messages;
}
