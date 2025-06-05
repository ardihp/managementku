import { Workspace } from "@/types/workspace";
import React from "react";
import WorkspaceAction from "./action";
import { paths } from "@/const/routes";
import Link from "next/link";
import ColorfulBadge from "@/components/layout/colorful-badge";
import { Avatar, AvatarFallback } from "@/components/ui/avatar";

interface WorkspaceCardProps {
  setSelectedItem?: (workspace: Workspace) => void;
  workspace: Workspace;
  isArchived: boolean;
}

export default function WorkspaceCard({
  workspace,
  setSelectedItem,
  isArchived = false,
}: WorkspaceCardProps) {
  return (
    <div className="relative">
      <WorkspaceAction
        setSelectedItem={setSelectedItem}
        workspace={workspace}
        isArchived={isArchived}
      />

      <Link href={`${paths.dashboard.workspace}/${workspace.id}`}>
        <div className="w-full border rounded-xl bg-white flex flex-col p-4 gap-2 min-h-[177px]">
          <div className="flex items-center gap-2">
            <ColorfulBadge
              text={workspace.type.name}
              color={workspace.type.color}
            />

            {workspace.tags.map((tag, key) => (
              <ColorfulBadge key={key} text={tag} color="gray" />
            ))}
          </div>

          <p
            className="font-medium text-lg capitalize leading-[25px] line-clamp-2"
            title={workspace.name}
          >
            {workspace.name}
          </p>

          <p
            className="font-quicksand text-xs font-medium text-zinc-400 capitalize leading-[18px] line-clamp-2"
            title={workspace.description}
          >
            {workspace.description}
          </p>

          <div className="*:data-[slot=avatar]:ring-zinc-200 *:data-[slot=avatar]:bg-zinc-100 flex -space-x-2 *:data-[slot=avatar]:ring-2 mt-3">
            {workspace.members.map((member, key) => (
              <Avatar key={key} title={member.username} className="size-8">
                <AvatarFallback className="uppercase text-sm">
                  {member.username?.split(" ")?.map((text) => text?.[0])}
                </AvatarFallback>
              </Avatar>
            ))}
          </div>
        </div>
      </Link>
    </div>
  );
}
