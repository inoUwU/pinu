import type { Metadata } from 'next';
import { Geist, Geist_Mono } from 'next/font/google';
import { cookies } from 'next/headers';
import AppSidebar from '@/components/AdminSidebar';
import Navbar from '@/components/Navbar';
import { SidebarProvider } from '@/components/ui/sidebar';

const geistSans = Geist({
  variable: '--font-geist-sans',
  subsets: ['latin'],
});

const geistMono = Geist_Mono({
  variable: '--font-geist-mono',
  subsets: ['latin'],
});

export const metadata: Metadata = {
  title: 'Pinu - Admin Dashboard',
  description: 'Admin dashboard for Pinu',
};

export default async function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const cookieStore = await cookies();
  const defaultOpen = cookieStore.get('sidebar_state')?.value === 'true';
  return (
    <html lang='ja'>
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased flex`}
      >
        <SidebarProvider defaultOpen={defaultOpen}>
          <AppSidebar />
          <main className='w-full'>
            <Navbar />
            <div className='px-4'>{children}</div>
          </main>
        </SidebarProvider>
      </body>
    </html>
  );
}
