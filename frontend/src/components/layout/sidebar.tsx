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
import { Input } from "../ui/input";
import { sidebarItem, sidebarItem2 } from "@/const/sidebar-item";
import Link from "next/link";
import { cn } from "@/lib/utils";
import { paths } from "@/const/routes";
import { useRouter } from "next/router";
import { useAuth, UserDetail } from "@/hooks/auth/use-auth";

export default function Sidebar() {
  const router = useRouter();
  const { user, userSignOut } = useAuth();
  const [userDetail, setUserDetail] = useState<UserDetail | null>();

  useEffect(() => {
    setUserDetail(user);
  }, []);

  return (
    <div className="p-4 flex flex-col gap-5 h-full border-r border-zinc-200">
      <div className="flex flex-col gap-4">
        <div className="flex items-center justify-between">
          <p className="text-xl font-medium tracking-wider">Tralalelo</p>

          <PanelLeftClose className="size-5 opacity-40 cursor-pointer" />
        </div>

        <div className="flex items-center px-2 justify-center w-full border border-zinc-300 rounded-lg bg-secondary/30">
          <Search className="size-5 opacity-50 w-full max-w-5 pb-0.5" />

          <Input
            type="text"
            placeholder="Search..."
            wrapperClassname="border-none focus-within:ring-0 shadow-none py-0! px-0! w-full"
            className="bg-transparent! shadow-none"
          />
        </div>
      </div>

      <hr className="border-zinc-200" />

      <div className="flex flex-col gap-1">
        <p className="mb-3 opacity-40 text-sm font-medium px-3">Menu</p>

        {sidebarItem.map((item, index) => (
          <Link
            href={item.href}
            key={index}
            className={cn(
              "flex items-center gap-4 p-3 rounded-lg bg-white border duration-200",
              router.asPath === item.href ||
                (router.asPath.includes(item.href) &&
                  item.href !== paths.dashboard.root)
                ? "shadow border-zinc-200 opacity-100"
                : "border-transparent hover:bg-zinc-300/50 opacity-30"
            )}
          >
            {item.icon}

            <p className="font-medium">{item.title}</p>
          </Link>
        ))}
      </div>

      {userDetail?.role === "manager" && (
        <>
          <hr className="border-zinc-200" />

          <div className="flex flex-col gap-1">
            {sidebarItem2.map((item, index) => (
              <Link
                href={item.href}
                key={index}
                className={cn(
                  "flex items-center gap-4 p-3 rounded-lg bg-white border duration-200",
                  router.asPath === item.href ||
                    (router.asPath.includes(item.href) &&
                      item.href !== paths.dashboard.root)
                    ? "shadow border-zinc-200 opacity-100"
                    : "border-transparent hover:bg-zinc-300/50 opacity-30"
                )}
              >
                {item.icon}

                <p className="font-medium">{item.title}</p>
              </Link>
            ))}
          </div>
        </>
      )}

      <DropdownMenu>
        <DropdownMenuTrigger className="mt-auto p-2 bg-white rounded-lg shadow border border-zinc-200 cursor-pointer">
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
        <DropdownMenuContent className="mb-2 w-[calc(300px_-_32px)]">
          <DropdownMenuLabel>My Account</DropdownMenuLabel>
          <DropdownMenuSeparator />
          <DropdownMenuItem>Profile</DropdownMenuItem>
          <DropdownMenuItem onClick={userSignOut}>Logout</DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
    </div>
  );
}
