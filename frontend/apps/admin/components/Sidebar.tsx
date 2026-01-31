import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@workspace/ui/components/dropdown-menu";
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarSeparator,
} from "@workspace/ui/components/sidebar";
import {
  ChevronUp,
  CookingPot,
  Home,
  QrCode,
  User,
  User2,
  Wrench,
} from "lucide-react";
import Link from "next/link";

const items = [
  {
    title: "Home",
    url: "/",
    icon: <Home aria-hidden='true' />,
  },
  {
    title: "User",
    url: "user",
    icon: <User aria-hidden='true' />,
  },
  {
    title: "Menu",
    url: "menu",
    icon: <CookingPot aria-hidden='true' />,
  },
  {
    title: "Settings",
    url: "settings",
    icon: <Wrench aria-hidden='true' />,
  },
  {
    title: "Qr Code",
    url: "qrcode",
    icon: <QrCode aria-hidden='true' />,
  },
];

const AppSidebar = async () => {
  return (
    <Sidebar collapsible='icon'>
      <SidebarHeader>
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton asChild>
              <Link href='/admin'>
                <span className='text-2xl font-bold'>Pinu</span>
              </Link>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarHeader>
      <SidebarSeparator className='mx-0' />
      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupLabel>Application</SidebarGroupLabel>
          <SidebarGroupContent>
            <SidebarMenu>
              {items.map(item => (
                <SidebarMenuItem key={item.title}>
                  <SidebarMenuButton asChild>
                    <Link href={item.url}>
                      {item.icon}
                      <span>{item.title}</span>
                    </Link>
                  </SidebarMenuButton>
                </SidebarMenuItem>
              ))}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>
      <SidebarFooter>
        <SidebarMenu>
          <SidebarMenuItem>
            <DropdownMenu>
              <DropdownMenuTrigger asChild>
                <SidebarMenuButton>
                  <User2 aria-hidden='true' />
                  John Doe
                  <ChevronUp className='ml-auto' aria-hidden='true' />
                </SidebarMenuButton>
              </DropdownMenuTrigger>
              <DropdownMenuContent className='align-end'>
                <DropdownMenuItem>Account</DropdownMenuItem>
                <DropdownMenuItem>Sign Out</DropdownMenuItem>
              </DropdownMenuContent>
            </DropdownMenu>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarFooter>
    </Sidebar>
  );
};

export default AppSidebar;
