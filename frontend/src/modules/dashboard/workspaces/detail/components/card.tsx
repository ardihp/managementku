import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import { cn } from "@/lib/utils";
import { WorkspaceCard, WorkspaceCardMember } from "@/types/workspace-card";
import { Draggable } from "@hello-pangea/dnd";
import React, { useState } from "react";
import WorkspaceAction from "./action";
import { Link2, Text } from "lucide-react";
import DetailWorkspaceCard from "./detail-card";

interface TaskCardProps {
  workspaceId?: number;
  isArchived: boolean;
  card: WorkspaceCard;
  setSelectedCard?: (val: WorkspaceCard) => void;
  customClassName?: string;
  isDragging?: boolean;
  ref?: (element?: HTMLElement | null) => void;
}

export default function TaskCard({
  card,
  workspaceId = 0,
  customClassName,
  setSelectedCard = () => {},
  isArchived = false,
  isDragging,
  ref,
  ...rest
}: TaskCardProps) {
  const [openDrawer, setOpenDrawer] = useState(false);

  return (
    <>
      <DetailWorkspaceCard
        open={openDrawer}
        setSelectedCard={setSelectedCard}
        onOpenChange={() => setOpenDrawer(false)}
        card={card}
        workspaceId={workspaceId}
      />

      <div
        className={cn(
          "relative opacity-100 group cursor-pointer",
          isDragging && "opacity-85",
          customClassName
        )}
        ref={ref}
        {...rest}
      >
        <WorkspaceAction
          setSelectedCard={setSelectedCard}
          card={card}
          isArchived={isArchived}
          workspaceId={workspaceId}
        />

        <div
          className="w-full border rounded-xl bg-white group-hover:bg-zinc-50/50 flex flex-col p-4 gap-3 duration-200"
          onClick={() => (isArchived ? {} : setOpenDrawer(true))}
        >
          <p
            className="font-medium capitalize leading-none line-clamp-2"
            title={card.title}
          >
            {card.title}
          </p>

          {card.description && (
            <div className="flex items-center gap-3">
              <div title="This card has a description.">
                <Text className="size-4 opacity-70" />
              </div>

              {card.description.includes("<a") && (
                <div title="This card has a link.">
                  <Link2 className="size-4 opacity-70" />
                </div>
              )}
            </div>
          )}

          {card.members.length >= 1 && (
            <div className="*:data-[slot=avatar]:ring-zinc-200 *:data-[slot=avatar]:bg-zinc-100 flex -space-x-2 *:data-[slot=avatar]:ring-2 mt-auto">
              {card.members.map((member, key) => (
                <Avatar key={key} title={member.username} className="size-8">
                  <AvatarFallback className="uppercase text-sm">
                    {member.username?.split(" ")?.map((text) => text?.[0])}
                  </AvatarFallback>
                </Avatar>
              ))}
            </div>
          )}
        </div>
      </div>
    </>
  );
}
