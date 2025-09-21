"use client";
import { QRCodeSVG } from "qrcode.react";

const QR_CODE_URL = "https://nextjs.org/";
const QR_IMAGE_SRC = "/favicon.ico";
const QR_IMAGE_SIZE = 24;
const QR_CONTAINER_SIZE = 450;

export default function QRPage() {
  return (
    <div className='h-dvh w-dvw flex flex-col items-center justify-center'>
      <div
        className={`aspect-square max-w-[${QR_CONTAINER_SIZE}px] lg:w-[${QR_CONTAINER_SIZE}px] bg-neutral-200 rounded-lg flex items-center justify-center`}
      >
        <QRCodeSVG
          value={QR_CODE_URL}
          className='w-full h-full p-6'
          fgColor='#000000'
          imageSettings={{
            src: QR_IMAGE_SRC,
            height: QR_IMAGE_SIZE,
            width: QR_IMAGE_SIZE,
            excavate: true,
          }}
        />
      </div>
    </div>
  );
}
