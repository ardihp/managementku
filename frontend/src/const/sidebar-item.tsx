import {
  Archive,
  FolderOpen,
  Gauge,
  MessageSquareDot,
  StretchHorizontal,
} from "lucide-react";
import { paths } from "./routes";
import { cn } from "@/lib/utils";
import { ReactNode } from "react";

interface SidebarItem {
  title: string;
  href: string;
  icon: ReactNode;
}

export const sidebarItem: SidebarItem[] = [
  {
    title: "Dashboard",
    href: paths.dashboard.root,
    icon: <Gauge className={cn("size-6")} />,
  },
  {
    title: "Workspaces",
    href: paths.dashboard.workspace,
    icon: <FolderOpen className={cn("size-6")} />,
  },
  {
    title: "Notifications",
    href: paths.dashboard.notification,
    icon: <MessageSquareDot className={cn("size-6")} />,
  },
];

export const sidebarItemManager: SidebarItem[] = [
  {
    title: "Workspace Types",
    href: paths.dashboard.type,
    icon: <StretchHorizontal className={cn("size-6")} />,
  },
];
