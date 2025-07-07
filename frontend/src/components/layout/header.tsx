import React from "react";
import { SidebarTrigger } from "../ui/sidebar";

export default function Header() {
  return (
    <div className="p-4 border-b border-zinc-200">
      <SidebarTrigger />
    </div>
  );
}
