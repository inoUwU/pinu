import LoginPageHeader from "@/app/admin/(auth)/components/header";

export default async function AuthLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <>
      <div className='h-20'>
        <LoginPageHeader />
      </div>
      <main className='w-full'>
        <div className='px-4'>{children}</div>
      </main>
    </>
  );
}
