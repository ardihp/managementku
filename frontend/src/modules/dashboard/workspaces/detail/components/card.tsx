import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import { cn } from "@/lib/utils";
import { WorkspaceCard, WorkspaceCardMember } from "@/types/workspace-card";
import { Draggable } from "@hello-pangea/dnd";
import React from "react";

interface TaskCardProps {
  card: WorkspaceCard;
  index: number;
  customClassName?: string;
}

export default function TaskCard({
  card,
  index,
  customClassName,
}: TaskCardProps) {
  return (
    <Draggable draggableId={card.id.toString()} index={index}>
      {({ dragHandleProps, draggableProps, innerRef }, { isDragging }) => {
        return (
          <div
            className={cn(
              "relative opacity-100 group",
              isDragging && "opacity-85",
              customClassName
            )}
            ref={innerRef}
            {...draggableProps}
            {...dragHandleProps}
          >
            <div className="w-full border rounded-xl bg-white group-hover:bg-zinc-50/50 flex flex-col p-4 gap-2 duration-200">
              <p
                className="font-medium capitalize leading-none line-clamp-2"
                title={card.title}
              >
                {card.title}
              </p>

              {card.description && (
                <p
                  className="font-quicksand text-xs font-medium text-zinc-400 capitalize leading-[18px] line-clamp-2"
                  title={card.description}
                >
                  {card.description}
                </p>
              )}

              {card.members.length >= 1 && (
                <div className="*:data-[slot=avatar]:ring-zinc-200 *:data-[slot=avatar]:bg-zinc-100 flex -space-x-2 *:data-[slot=avatar]:ring-2 mt-auto">
                  {card.members.map((member, key) => (
                    <Avatar
                      key={key}
                      title={member.username}
                      className="size-8"
                    >
                      <AvatarFallback className="uppercase text-sm">
                        {member.username?.split(" ")?.map((text) => text?.[0])}
                      </AvatarFallback>
                    </Avatar>
                  ))}
                </div>
              )}
            </div>
          </div>
        );
      }}
    </Draggable>
  );
}
