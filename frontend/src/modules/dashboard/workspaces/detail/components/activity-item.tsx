import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import { useWorkspaceCard } from "@/hooks/use-workspace-card";
import { WorkspaceCardActivity } from "@/types/workspace-card";
import dayjs from "dayjs";
import React, { useContext, useState } from "react";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { Button } from "@/components/ui/button";
import { useAuth } from "@/hooks/auth/use-auth";
import { SocketContext } from "@/context/socket";

interface ActivityItemProps {
  activity: WorkspaceCardActivity;
  workspaceId: number;
  cardId: number;
  onSelectEdit: (activity_id: number) => void;
  afterDelete: () => void;
}

export default function ActivityItem({
  activity,
  workspaceId,
  cardId,
  onSelectEdit = () => {},
  afterDelete = () => {},
}: ActivityItemProps) {
  const [openPopoverDelete, setOpenPopoverDelete] = useState(false);
  const { user } = useAuth();
  const { socket } = useContext(SocketContext);

  const { deleteWorkspaceCardActivity } = useWorkspaceCard();

  const { mutateAsync: deleteActivity } = deleteWorkspaceCardActivity(
    workspaceId,
    cardId
  );

  async function handleDelete(activity_id: number) {
    try {
      await deleteActivity(activity_id);
    } catch (error) {
      console.error(error);
    }
    setOpenPopoverDelete(false);
    afterDelete();
    socket?.send(
      JSON.stringify({
        card_id: cardId,
        workspace_id: workspaceId,
        type: "card-activity",
      })
    );
  }

  return (
    <div className="flex gap-4 items-start first:border-0 border-t py-4 px-[25px] bg-transparent hover:bg-zinc-300/10 duration-200">
      <Avatar className="size-10">
        <AvatarFallback className="capitalize">
          {activity.user_created[0]}
        </AvatarFallback>
      </Avatar>

      <div className="flex flex-col w-full gap-2">
        <div className="flex items-center justify-between">
          <p className="font-medium capitalize truncate">
            {activity.user_created}
          </p>
          <p className="text-xs">
            {dayjs(activity.created_at).format("DD/MM/YYYY, HH:mm:ss")} WIB
          </p>
        </div>

        <div
          className="text-[12px] opacity-80 description-detail p-2 bg-zinc-50 shadow-xs border rounded-lg"
          dangerouslySetInnerHTML={{
            __html: activity.message,
          }}
        />

        {((activity.type === "comment" &&
          user?.username === activity.user_created) ||
          activity.created_at !== activity.updated_at) && (
          <div className="flex items-center gap-3">
            {activity.type === "comment" &&
              user?.username === activity.user_created && (
                <>
                  <div
                    className="cursor-pointer"
                    onClick={() => onSelectEdit(activity.id)}
                  >
                    <p className="text-xs">Edit</p>
                  </div>

                  <Popover
                    open={openPopoverDelete}
                    onOpenChange={setOpenPopoverDelete}
                  >
                    <PopoverTrigger className="cursor-pointer">
                      <p className="text-xs">Delete</p>
                    </PopoverTrigger>
                    <PopoverContent
                      className="flex flex-col gap-3"
                      align="start"
                    >
                      <p className="text-sm">
                        Deleting a comment is forever. There is no undo.
                      </p>

                      <Button
                        size="sm"
                        variant="destructive"
                        className="w-fit ml-auto cursor-pointer"
                        onClick={() => handleDelete(activity.id)}
                      >
                        <p className="text-sm">Delete</p>
                      </Button>
                    </PopoverContent>
                  </Popover>
                </>
              )}

            {activity.created_at !== activity.updated_at && (
              <p className="text-xs ml-auto">Edited</p>
            )}
          </div>
        )}
      </div>
    </div>
  );
}
