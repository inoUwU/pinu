"use client";
import { Skeleton } from "@workspace/ui/components/skeleton";
import { QRCodeSVG } from "qrcode.react";
import { useEffect, useState } from "react";

/*
  このページはテスト用にスマートフォンからアクセスするためのQRコードを表示するページです。
*/

const QR_IMAGE_SRC = "/favicon.ico";
const QR_IMAGE_SIZE = 24;
const QR_CONTAINER_SIZE = 450;
const CLIENT_ENTRY_PATH = "/category";
const DEFAULT_FRONTEND_PORT = "3000";

export default function QRPage() {
  const [qrCodeUrl, setQrCodeUrl] = useState("");
  const [loading, setLoading] = useState(true);

  // サーバーのローカルIPアドレスを取得
  // 空の dependency arrayを渡して、マウント時に一度だけ実行
  useEffect(() => {
    const fetchLocalIP = async () => {
      try {
        const response = await fetch("/api/local-ip");
        if (!response.ok) {
          throw new Error("IPアドレスの取得に失敗しました");
        }

        const data = await response.json();
        const ip = data.ip;
        const protocol = window.location.protocol;
        const port = window.location.port || DEFAULT_FRONTEND_PORT;

        // QRコードに含めるURL
        const url = `${protocol}//${ip}:${port}${CLIENT_ENTRY_PATH}`;
        console.debug("QRコードのURL:", url);

        setQrCodeUrl(url);
      } catch (err) {
        console.error(err);
      } finally {
        setLoading(false);
      }
    };

    fetchLocalIP();
  }, []);

  return (
    <div className='h-dvh w-dvw flex flex-col items-center justify-center'>
      <div
        className='flex aspect-square items-center justify-center rounded-lg bg-neutral-200'
        style={{
          maxWidth: QR_CONTAINER_SIZE,
          width: "100%",
        }}
      >
        {loading ? (
          <Skeleton
            className='w-full h-full p-6'
            style={{
              height: 128,
              minHeight: 128,
            }}
          />
        ) : (
          <QRCodeSVG
            value={qrCodeUrl}
            className='w-full h-full p-6'
            fgColor='#000000'
            imageSettings={{
              src: QR_IMAGE_SRC,
              height: QR_IMAGE_SIZE,
              width: QR_IMAGE_SIZE,
              excavate: true,
            }}
          />
        )}
      </div>
    </div>
  );
}
