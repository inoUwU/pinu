import { ThemeProvider } from "@/components/providers/ThemeProvider";
export default function QRLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <>
      <ThemeProvider
        attribute='class'
        defaultTheme='light'
        enableSystem
        disableTransitionOnChange
      >
        <div>{children}</div>
      </ThemeProvider>
    </>
  );
}
