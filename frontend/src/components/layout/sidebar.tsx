import { ChevronDown, PanelLeftClose, Plus, Search } from "lucide-react";
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

export default function Sidebar() {
  return (
    <div className="p-4 h-full w-full">
      <div className="p-4 flex flex-col bg-secondary rounded-xl border border-zinc-200 shadow h-full">
        <div className="flex flex-col gap-4">
          <div className="flex items-center justify-between">
            <p className="text-xl font-medium tracking-wider">Tralalelo</p>

            <PanelLeftClose className="size-5 opacity-40 cursor-pointer" />
          </div>

          <div className="flex items-center gap-3">
            <div className="flex items-center gap-2 justify-center w-full py-2 border border-zinc-300 shadow-sm shadow-zinc-200 rounded-lg cursor-pointer active:shadow-none duration-200">
              <Plus className="size-4" />
              <p className="text-sm font-normal">Create Project</p>
            </div>
            <div className="flex items-center gap-2 justify-center w-full max-w-10 h-full border border-zinc-300 shadow-sm shadow-zinc-200 rounded-lg cursor-pointer active:shadow-none duration-200">
              <Search className="size-4 opacity-40 stroke-3" />
            </div>
          </div>
        </div>

        <DropdownMenu>
          <DropdownMenuTrigger className="mt-auto flex items-center gap-3 cursor-pointer focus-within:outline-none">
            <Avatar className="size-10 border border-zinc-300">
              <AvatarImage src="https://github.com/shadcn.png" />
              <AvatarFallback>A</AvatarFallback>
            </Avatar>

            <p className="font-medium font-quicksand">Puhdidi</p>

            <div className="flex items-center justify-center ml-auto">
              <ChevronDown className="size-4 stroke-3 opacity-40" />
            </div>
          </DropdownMenuTrigger>
          <DropdownMenuContent className="w-[242px] mb-4">
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
