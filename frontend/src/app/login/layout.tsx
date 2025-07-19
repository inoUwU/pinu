import type { Metadata } from "next";
import { Geist, Geist_Mono } from "next/font/google";
import LoginPageHeader from "@/app/login/components/header";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "Pinu - Admin Dashboard",
  description: "Admin dashboard for Pinu",
};

export default async function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang='ja' className=''>
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased`}
      >
        <div className='h-20'>
          <LoginPageHeader />
        </div>
        <main className='w-full'>
          <div className='px-4'>{children}</div>
        </main>
      </body>
    </html>
  );
}
