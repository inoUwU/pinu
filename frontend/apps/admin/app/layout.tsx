import { Toaster } from "@workspace/ui/components/sonner";
import type { Metadata } from "next";
import NextTopLoader from "nextjs-toploader";
import { ThemeProvider } from "@/components/providers/ThemeProvider";

import "@workspace/ui/globals.css";

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
    <html lang='ja'>
      <body className='antialiased font-sans'>
        <a
          href='#main-content'
          className='sr-only focus:not-sr-only focus:absolute focus:top-4 focus:left-4 focus:z-50 focus:bg-primary focus:text-primary-foreground focus:px-4 focus:py-2 focus:rounded'
        >
          メインコンテンツへスキップ
        </a>
        <NextTopLoader crawl={false} />
        <Toaster />
        <ThemeProvider
          attribute='class'
          defaultTheme='light'
          enableSystem
          disableTransitionOnChange
        >
          {children}
        </ThemeProvider>
      </body>
    </html>
  );
}
