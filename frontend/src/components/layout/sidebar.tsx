import { ChevronsUpDown, PanelLeftClose, Search } from "lucide-react";
import React, { useEffect, useState } from "react";
import { Avatar, AvatarFallback, AvatarImage } from "../ui/avatar";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { sidebarItem, sidebarItemManager } from "@/const/sidebar-item";
import Link from "next/link";
import { cn } from "@/lib/utils";
import { paths } from "@/const/routes";
import { useRouter } from "next/router";
import { useAuth, UserDetail } from "@/hooks/auth/use-auth";
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
} from "../ui/sidebar";
import FormUser from "@/modules/dashboard/employees/components/form";
import Cookies from "js-cookie";

export default function AppSidebar() {
  const router = useRouter();
  const { user, userSignOut } = useAuth();
  const [userDetail, setUserDetail] = useState<UserDetail | null>();
  const [openDialog, setOpenDialog] = useState(false);

  useEffect(() => {
    setUserDetail(user);
  }, []);

  return (
    <>
      <FormUser
        open={openDialog}
        onOpenChange={() => setOpenDialog(false)}
        user={userDetail!}
        afterUpdate={(val) => {
          Cookies.set("user", JSON.stringify(val));
          setUserDetail(val);
        }}
      />

      <Sidebar>
        <SidebarHeader>
          <SidebarMenu>
            <SidebarMenuItem className="flex flex-col gap-4 px-1 pt-2">
              <div className="flex items-center justify-between">
                <p className="text-xl font-medium tracking-wider">Tralalelo</p>
              </div>
            </SidebarMenuItem>
          </SidebarMenu>
        </SidebarHeader>
        <SidebarContent>
          <SidebarGroup>
            <SidebarGroupLabel>Menu</SidebarGroupLabel>
            <SidebarGroupContent>
              <SidebarMenu>
                {sidebarItem.map((item, index) => (
                  <SidebarMenuItem key={index}>
                    <SidebarMenuButton asChild>
                      <Link
                        href={item.href}
                        className={cn(
                          "flex items-center gap-3 min-h-10 rounded-lg bg-white border duration-200",
                          router.asPath === item.href ||
                            (router.asPath.includes(item.href) &&
                              item.href !== paths.dashboard.root)
                            ? "shadow border-zinc-200 opacity-100"
                            : "border-transparent hover:bg-zinc-300/50 opacity-30"
                        )}
                      >
                        {item.icon}

                        <p className="font-medium leading-none mt-0.5">
                          {item.title}
                        </p>
                      </Link>
                    </SidebarMenuButton>
                  </SidebarMenuItem>
                ))}
              </SidebarMenu>
            </SidebarGroupContent>
          </SidebarGroup>

          {userDetail?.role === "manager" && (
            <SidebarGroup>
              <SidebarGroupLabel>Settings</SidebarGroupLabel>
              <SidebarGroupContent>
                <SidebarMenu>
                  {sidebarItemManager.map((item, index) => (
                    <SidebarMenuItem key={index}>
                      <SidebarMenuButton asChild>
                        <Link
                          href={item.href}
                          className={cn(
                            "flex items-center gap-3 min-h-10 rounded-lg bg-white border duration-200",
                            router.asPath === item.href ||
                              (router.asPath.includes(item.href) &&
                                item.href !== paths.dashboard.root)
                              ? "shadow border-zinc-200 opacity-100"
                              : "border-transparent hover:bg-zinc-300/50 opacity-30"
                          )}
                        >
                          {item.icon}

                          <p className="font-medium leading-none mt-0.5">
                            {item.title}
                          </p>
                        </Link>
                      </SidebarMenuButton>
                    </SidebarMenuItem>
                  ))}
                </SidebarMenu>
              </SidebarGroupContent>
            </SidebarGroup>
          )}
        </SidebarContent>

        <SidebarFooter>
          <SidebarMenu>
            <SidebarMenuItem>
              <DropdownMenu>
                <DropdownMenuTrigger className="mt-auto p-2 bg-white rounded-lg shadow border border-zinc-200 cursor-pointer w-full">
                  <div className="flex w-full items-center gap-3 focus-within:outline-none">
                    <Avatar className="size-10 border border-zinc-300">
                      <AvatarImage src="https://github.com/shadcn.png" />
                      <AvatarFallback>A</AvatarFallback>
                    </Avatar>

                    <div className="flex flex-col items-start gap-1.5">
                      <p className="text-xs text-zinc-400 font-light capitalize leading-none">
                        {userDetail?.role}
                      </p>
                      <p className="font-medium font-quicksand text-left capitalize leading-none">
                        {userDetail?.name}
                      </p>
                    </div>

                    <div className="flex items-center justify-center ml-auto">
                      <ChevronsUpDown className="size-4 opacity-40 duration-200" />
                    </div>
                  </div>
                </DropdownMenuTrigger>

                <DropdownMenuContent
                  className="mb-2 w-[calc(300px_-_32px)]"
                  side="right"
                  align="center"
                >
                  <DropdownMenuLabel>Options</DropdownMenuLabel>
                  <DropdownMenuSeparator />
                  <DropdownMenuItem onClick={() => setOpenDialog(true)}>
                    Edit Profile
                  </DropdownMenuItem>
                  <DropdownMenuItem onClick={userSignOut}>
                    Logout
                  </DropdownMenuItem>
                </DropdownMenuContent>
              </DropdownMenu>
            </SidebarMenuItem>
          </SidebarMenu>
        </SidebarFooter>
      </Sidebar>
    </>
  );
}
