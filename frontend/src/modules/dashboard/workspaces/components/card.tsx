import { Workspace } from "@/types/workspace";
import React, { memo } from "react";
import WorkspaceAction from "./action";
import { paths } from "@/const/routes";
import Link from "next/link";
import ColorfulBadge from "@/components/layout/colorful-badge";
import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import { cn } from "@/lib/utils";
import { useAuth } from "@/hooks/auth/use-auth";

interface WorkspaceCardProps {
  isArchived: boolean;
  setSelectedWorkspace?: (workspace: Workspace) => void;
  customClassName?: string;
  workspace: Workspace;
  isDragging?: boolean;
  ref?: (element?: HTMLElement | null) => void;
}

const WorkspaceCard = ({
  workspace,
  setSelectedWorkspace,
  customClassName,
  isArchived = false,
  isDragging,
  ref,
  ...rest
}: WorkspaceCardProps) => {
  const { user } = useAuth();

  return (
    <div
      className={cn(
        "relative opacity-100 group",
        isDragging && "opacity-85",
        customClassName
      )}
      ref={ref}
      {...rest}
    >
      {user!.role === "manager" && (
        <WorkspaceAction
          setSelectedWorkspace={setSelectedWorkspace}
          workspace={workspace}
          isArchived={isArchived}
        />
      )}

      <Link href={`${paths.dashboard.workspace}/${workspace.id}`}>
        <div className="w-full border rounded-xl bg-white group-hover:bg-zinc-50/50 flex flex-col p-4 gap-2 min-h-[152px] duration-200">
          <div className="flex items-center gap-2 overflow-auto">
            <ColorfulBadge content={"🛠️"} color={workspace.status.color} />
            <ColorfulBadge
              content={workspace.type.name}
              color={workspace.type.color}
            />

            {workspace.tags.map((tag, key) => (
              <ColorfulBadge key={key} content={tag} color="gray" />
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

          <div className="*:data-[slot=avatar]:ring-zinc-200 *:data-[slot=avatar]:bg-zinc-100 flex -space-x-2 *:data-[slot=avatar]:ring-2 mt-auto ml-auto">
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
};

export default memo(WorkspaceCard);
