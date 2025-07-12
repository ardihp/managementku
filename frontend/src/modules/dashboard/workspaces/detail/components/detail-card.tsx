import React, { useContext, useEffect, useState } from "react";
import {
  Drawer,
  DrawerContent,
  DrawerDescription,
  DrawerHeader,
  DrawerTitle,
} from "@/components/ui/drawer";
import { Button } from "@/components/ui/button";
import { WorkspaceCard, WorkspaceCardActivity } from "@/types/workspace-card";
import { ChevronRight, CircleX, Loader2 } from "lucide-react";
import dayjs from "dayjs";
import WorkspaceAction from "./action";
import ColorfulBadge from "@/components/layout/colorful-badge";
import TextEditor from "@/components/ui/text-editor";
import { useWorkspaceCard } from "@/hooks/use-workspace-card";
import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import ActivityItem from "./activity-item";
import { SocketContext } from "@/context/socket";
import { useQueryClient } from "@tanstack/react-query";

interface DetailWorkspaceCardProps {
  open: boolean;
  card: WorkspaceCard;
  workspaceId: number;
  isAdmin: boolean;
  setSelectedCard: (val: WorkspaceCard) => void;
  onOpenChange: () => void;
}

export default function DetailWorkspaceCard({
  open,
  card,
  workspaceId,
  isAdmin,
  setSelectedCard,
  onOpenChange,
}: DetailWorkspaceCardProps) {
  const [message, setMessage] = useState("");
  const [refetch, setRefetch] = useState(0);
  const [contentLoaded, setContentLoaded] = useState(false);
  const [editComment, setEditCommend] = useState(false);
  const [selectedComment, setSelectedComment] = useState(0);
  const { socket } = useContext(SocketContext);

  const {
    getListWorkspaceCardActivities,
    createWorkspaceCardActivity,
    updateWorkspaceCardActivity,
  } = useWorkspaceCard();

  const { data: activities, isLoading: loadingActivity } =
    getListWorkspaceCardActivities(workspaceId, open ? card.id : 0);
  const { mutateAsync: create, isPending: loadingCreate } =
    createWorkspaceCardActivity(workspaceId, open ? card.id : 0);
  const { mutateAsync: update, isPending: loadingUpdate } =
    updateWorkspaceCardActivity(workspaceId, open ? card.id : 0);

  useEffect(() => {
    if (contentLoaded && open) {
      const descriptionheight =
        document.getElementById("description")?.offsetHeight || 0;
      const detailHeaderheight =
        document.getElementById("detail-header")?.offsetHeight || 0;
      const floatingFormheight =
        document.getElementById("floating-form")?.offsetHeight || 0;

      const wrapperElement = document.getElementById("wrapper-detail");
      if (wrapperElement) {
        wrapperElement.style.maxHeight = `calc(100% - ${
          detailHeaderheight + floatingFormheight
        }px)`;
      }

      const activeElement = document.getElementById("activities");
      if (activeElement) {
        activeElement.style.maxHeight = `calc(100% - ${descriptionheight}px)`;
      }
    }
  }, [contentLoaded, open, message]);

  useEffect(() => {
    setEditCommend(false);
    setSelectedComment(0);
    setContentLoaded(false);

    if (open) {
      setMessage("");
      setRefetch(refetch + 1);
    }

    setTimeout(() => {
      setContentLoaded(true);
    }, 100);
  }, [open]);

  async function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();

    try {
      if (editComment) {
        await update({ activity_id: selectedComment, data: { message } });
      } else {
        await create({ message });
      }
    } catch (error) {
      console.log(error);
    }

    setMessage("");
    setRefetch(refetch + 1);
    setEditCommend(false);
    setSelectedComment(0);
    socket?.send(
      JSON.stringify({
        text: "",
        type: "card-activity",
        workspace_id: workspaceId,
        card_id: card.id,
      })
    );
  }

  return (
    <Drawer
      direction="right"
      open={open}
      onOpenChange={onOpenChange}
      autoFocus={false}
    >
      <DrawerContent className="min-w-xl max-w-xl h-full">
        <DrawerHeader
          id="detail-header"
          className="flex flex-col gap-3! py-[17.5px] px-[25px] border-b"
        >
          <div className="flex items-center justify-between">
            <ColorfulBadge
              content={card.status.name}
              color={card.status.color}
            />

            <div className="relative h-full w-10 -top-2 -right-2">
              <WorkspaceAction
                setSelectedCard={() => {
                  setSelectedCard(card);
                  onOpenChange();
                }}
                card={card}
                isArchived={card.is_archived}
                workspaceId={workspaceId}
                afterAction={() => {
                  onOpenChange();
                  socket?.send(
                    JSON.stringify({
                      text: "",
                      type: "card-activity",
                      workspace_id: workspaceId,
                    })
                  );
                }}
                isAdmin={isAdmin}
              />
            </div>
          </div>

          <div className="flex flex-col gap-3">
            <DrawerTitle className="text-2xl leading-none capitalize font-quicksand">
              {card.title}
            </DrawerTitle>

            <p className="text-sm text-zinc-500 leading-none">
              Created By{" "}
              <span className="text-black font-medium capitalize">
                {card.user_created}
              </span>
              , {dayjs(card.created_at).format("DD/MM/YYYY HH:mm:ss")} WIB
            </p>
          </div>

          <div className="flex flex-col gap-3 mt-3">
            <p className="text-sm font-medium leading-none">
              Who watch this card
            </p>

            {card.members.length >= 1 ? (
              <div className="*:data-[slot=avatar]:ring-zinc-200 *:data-[slot=avatar]:bg-zinc-100 flex -space-x-2 *:data-[slot=avatar]:ring-2">
                {card.members.map((member, key) => (
                  <Avatar key={key} title={member.username} className="size-6">
                    <AvatarFallback className="uppercase text-xs">
                      {member.username?.split(" ")?.map((text) => text?.[0])}
                    </AvatarFallback>
                  </Avatar>
                ))}
              </div>
            ) : (
              <p>-</p>
            )}
          </div>

          <DrawerDescription className="hidden"></DrawerDescription>
        </DrawerHeader>

        {contentLoaded && (
          <div id="wrapper-detail" className="overflow-auto h-full">
            <div id="description" className="flex flex-col min-h-[74px]">
              <div className="py-[10.5px] px-[25px] border-b">
                <p className="font-semibold text-lg">Description</p>
              </div>

              {card.description ? (
                <div className="max-h-[200px] py-[10.5px] px-[25px] overflow-auto">
                  <div
                    className="whitespace-break-spaces break-words description-detail"
                    dangerouslySetInnerHTML={{ __html: card.description }}
                  />
                </div>
              ) : (
                <p className="py-[10.5px] px-[25px]">-</p>
              )}
            </div>

            <div id="activities" className="flex flex-col h-full">
              <div className="px-[25px] py-[10.5px] border-y">
                <p className="font-semibold text-lg leading-none">Activities</p>
              </div>

              <div className="flex flex-col h-full overflow-auto">
                {loadingActivity ? (
                  <div className="flex flex-col h-full items-center justify-center">
                    <Loader2 className="animate-spin mx-auto mb-2 size-6" />
                  </div>
                ) : (
                  activities?.map(
                    (activity: WorkspaceCardActivity, key: number) => (
                      <ActivityItem
                        key={key}
                        activity={activity}
                        cardId={card.id}
                        workspaceId={workspaceId}
                        onSelectEdit={(selected_id) => {
                          setMessage(activity.message);
                          setRefetch(refetch + 1);
                          setEditCommend(true);
                          setSelectedComment(selected_id);
                        }}
                        afterDelete={() => {
                          setMessage("");
                          setRefetch(refetch + 1);
                          setEditCommend(false);
                          setSelectedComment(0);
                        }}
                      />
                    )
                  )
                )}
              </div>
            </div>

            <form
              id="floating-form"
              onSubmit={handleSubmit}
              className="flex p-[14px] min-h-[105px] fixed bottom-0 border-t gap-2 w-xl z-10 bg-white"
            >
              <div className="flex flex-col gap-2 w-full">
                <TextEditor
                  value={message}
                  onValueChange={(newVal) => {
                    setMessage(newVal === "<p></p>" ? "" : newVal);
                  }}
                  editorClassName="max-h-[80px] overflow-auto"
                  refetch={refetch}
                  disabled={loadingCreate || loadingUpdate}
                />
              </div>

              <div className="flex flex-col">
                {editComment && (
                  <div
                    className="flex items-center justify-center cursor-pointer h-9 w-fit mx-auto"
                    onClick={() => {
                      setEditCommend(false);
                      setSelectedComment(0);
                      setMessage("");
                      setRefetch(refetch + 1);
                    }}
                  >
                    <CircleX className="size-6" />
                  </div>
                )}

                <Button
                  className="h-[34px] mt-auto"
                  disabled={!message || loadingCreate || loadingUpdate}
                >
                  <p className="leading-none">
                    {editComment ? "Save" : "Send"}
                  </p>
                </Button>
              </div>
            </form>
          </div>
        )}
      </DrawerContent>
    </Drawer>
  );
}
