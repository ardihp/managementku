import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import { cn } from "@/lib/utils";
import { WorkspaceCard } from "@/types/workspace-card";
import React, { useContext, useState } from "react";
import WorkspaceAction from "./action";
import { CircleCheck, Link2, MessageSquare, Text } from "lucide-react";
import DetailWorkspaceCard from "./detail-card";
import { SocketContext } from "@/context/socket";

interface TaskCardProps {
  workspaceId?: number;
  isArchived: boolean;
  card: WorkspaceCard;
  isAdmin?: boolean;
  setSelectedCard?: (val: WorkspaceCard) => void;
  customClassName?: string;
  isDragging?: boolean;
  ref?: (element?: HTMLElement | null) => void;
}

export default function TaskCard({
  card,
  isAdmin = false,
  workspaceId = 0,
  customClassName,
  setSelectedCard = () => {},
  isArchived = false,
  isDragging,
  ref,
  ...rest
}: TaskCardProps) {
  const [openDrawer, setOpenDrawer] = useState(false);
  const { socket } = useContext(SocketContext);

  return (
    <>
      <DetailWorkspaceCard
        open={openDrawer}
        setSelectedCard={setSelectedCard}
        onOpenChange={() => setOpenDrawer(false)}
        card={card}
        workspaceId={workspaceId}
        isAdmin={isAdmin}
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
          isAdmin={isAdmin}
          afterAction={() =>
            socket?.send(
              JSON.stringify({
                type: "mark-progress",
                workspace_id: workspaceId,
                card_id: card.id,
              })
            )
          }
        />

        <div
          className="w-full border rounded-xl flex flex-col p-4 gap-3 duration-200 bg-white group-hover:bg-zinc-50/50"
          onClick={() => setOpenDrawer(true)}
        >
          <div className="flex items-start gap-3">
            {card.is_completed && (
              <div title="This card mark completed">
                <CircleCheck className="size-6 opacity-70 text-green-700 fill-green-100" />
              </div>
            )}

            <p
              className="font-medium capitalize line-clamp-2 min-h-6"
              title={card.title}
            >
              {card.title}
            </p>
          </div>

          {(card.description ||
            card.members.length >= 1 ||
            card.comment_count >= 1) && (
            <div className="flex items-center gap-4 min-h-6">
              {card.description && (
                <div className="flex items-center gap-3">
                  <div title="This card has a description.">
                    <Text className="size-5 opacity-70" />
                  </div>

                  {card.description.includes("<a") && (
                    <div title="This card has a link.">
                      <Link2 className="size-5 opacity-70" />
                    </div>
                  )}
                </div>
              )}

              {card.comment_count >= 1 && (
                <div
                  className="flex items-center gap-2"
                  title="This card has a comment."
                >
                  <MessageSquare className="size-5 opacity-70" />
                  <p className="text-sm">{card.comment_count}</p>
                </div>
              )}

              {card.members.length >= 1 && (
                <div className="*:data-[slot=avatar]:ring-zinc-200 *:data-[slot=avatar]:bg-zinc-100 flex -space-x-2 *:data-[slot=avatar]:ring-2 mt-auto ml-auto">
                  {card.members.map((member, key) => (
                    <Avatar
                      key={key}
                      title={member.username}
                      className="size-6"
                    >
                      <AvatarFallback className="uppercase text-xs">
                        {member.username?.split(" ")?.map((text) => text?.[0])}
                      </AvatarFallback>
                    </Avatar>
                  ))}
                </div>
              )}
            </div>
          )}
        </div>
      </div>
    </>
  );
}
