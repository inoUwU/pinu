import type { Metadata } from "next";
import { Geist, Geist_Mono } from "next/font/google";
import TabBar from "@/app/components/TabBar";

import "@workspace/ui/globals.css";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "Pinu",
  description: "toy ordering system for restaurants",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ja" suppressHydrationWarning>
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased`}
      >
        <div className="flex flex-col h-screen w-screen">
          <div className="h-15/16">{children}</div>
          <div className="h-1/16 flex w-lvw">
            <TabBar />
          </div>
        </div>
      </body>
    </html>
  );
}
