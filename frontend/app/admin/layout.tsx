import type { Metadata } from "next";

export const metadata: Metadata = {
  title: "Pinu - Admin Dashboard",
  description: "Admin dashboard for Pinu",
};

export default function AdminLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return children;
}
