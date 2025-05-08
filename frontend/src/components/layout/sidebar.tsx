import { ChevronsUpDown, PanelLeftClose, Search } from "lucide-react";
import React from "react";
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

export default function Sidebar() {
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
            className="bg-transparent! border-none focus-visible:ring-0 shadow-none py-0! px-2!"
          />
        </div>
      </div>

      <hr className="border-zinc-200" />

      <div className="mt-auto p-2 bg-white rounded-lg shadow border border-zinc-200">
        <DropdownMenu>
          <DropdownMenuTrigger className="flex w-full items-center gap-3 cursor-pointer focus-within:outline-none">
            <Avatar className="size-10 border border-zinc-300">
              <AvatarImage src="https://github.com/shadcn.png" />
              <AvatarFallback>A</AvatarFallback>
            </Avatar>

            <div className="flex flex-col justify-center gap-0.5">
              <p className="text-xs text-zinc-400 font-light">Employee</p>
              <p className="font-medium font-quicksand">Puhdidi</p>
            </div>

            <div className="flex items-center justify-center ml-auto">
              <ChevronsUpDown className="size-4 opacity-40 duration-200" />
            </div>
          </DropdownMenuTrigger>
          <DropdownMenuContent className="mb-4">
            <DropdownMenuLabel>My Account</DropdownMenuLabel>
            <DropdownMenuSeparator />
            <DropdownMenuItem>Profile</DropdownMenuItem>
            <DropdownMenuItem>Logout</DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      </div>
    </div>
  );
}
