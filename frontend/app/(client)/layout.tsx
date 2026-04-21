import type { Metadata } from "next";
import TabBar from "@/app/(client)/components/TabBar";

export const metadata: Metadata = {
  title: "Pinu",
  description: "toy ordering system for restaurants",
};

export default function ClientLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <div className='flex min-h-screen w-screen flex-col'>
      <div className='flex-1'>{children}</div>
      <div className='shrink-0'>
        <TabBar />
      </div>
    </div>
  );
}
