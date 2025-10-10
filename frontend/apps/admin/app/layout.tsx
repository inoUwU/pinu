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
