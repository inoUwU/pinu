import { SidebarProvider } from "@workspace/ui/components/sidebar";
import { cookies } from "next/headers";
import Navbar from "@/components/Navbar";
import AppSidebar from "@/components/Sidebar";

export default async function AuthenticatedLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const cookieStore = await cookies();
  const defaultOpen = cookieStore.get("sidebar_state")?.value === "true";

  return (
    <SidebarProvider defaultOpen={defaultOpen}>
      <AppSidebar />
      <main className='w-full'>
        <Navbar />
        <div className='px-4'>{children}</div>
      </main>
    </SidebarProvider>
  );
}
