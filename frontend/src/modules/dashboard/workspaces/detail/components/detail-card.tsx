import React, { useEffect, useState } from "react";
import {
  Drawer,
  DrawerContent,
  DrawerDescription,
  DrawerHeader,
  DrawerTitle,
} from "@/components/ui/drawer";
import { Button } from "@/components/ui/button";
import { WorkspaceCard, WorkspaceCardActivity } from "@/types/workspace-card";
import { ChevronRight, Loader2 } from "lucide-react";
import dayjs from "dayjs";
import WorkspaceAction from "./action";
import ColorfulBadge from "@/components/layout/colorful-badge";
import TextEditor from "@/components/ui/text-editor";
import { useWorkspaceCard } from "@/hooks/use-workspace-card";
import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import { cn } from "@/lib/utils";

interface DetailWorkspaceCardProps {
  open: boolean;
  onOpenChange: () => void;
  card: WorkspaceCard;
  workspaceId: number;
  setSelectedCard: (val: WorkspaceCard) => void;
}

export default function DetailWorkspaceCard({
  open,
  onOpenChange,
  card,
  workspaceId,
  setSelectedCard,
}: DetailWorkspaceCardProps) {
  const [message, setMessage] = useState("");
  const [refetch, setRefetch] = useState(0);
  const [contentLoaded, setContentLoaded] = useState(false);

  const { getListWorkspaceCardActivities, createWorkspaceCardActivity } =
    useWorkspaceCard();

  const { data: activities, isLoading: loadingActivity } =
    getListWorkspaceCardActivities(workspaceId, open ? card.id : 0);

  const { mutateAsync: create, isPending: loadingCreate } =
    createWorkspaceCardActivity(workspaceId, open ? card.id : 0);

  useEffect(() => {
    if (contentLoaded && open) {
      const descriptionheight =
        document.getElementById("description")?.clientHeight || 0;
      const detailHeaderheight =
        document.getElementById("detail-header")?.clientHeight || 0;
      const floatingFormheight =
        document.getElementById("floating-form")?.clientHeight || 0;

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
    setMessage("");
    setRefetch(refetch + 1);

    if (open) {
      setContentLoaded(false);
    }

    setTimeout(() => {
      setContentLoaded(true);
    }, 100);
  }, [open]);

  async function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();

    try {
      await create({ message });
    } catch (error) {
      console.log(error);
    }

    setMessage("");
    setRefetch(refetch + 1);
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
                ) : activities?.length >= 1 ? (
                  activities?.map(
                    (activity: WorkspaceCardActivity, key: number) => (
                      <div
                        key={key}
                        className="flex gap-4 items-start first:border-0 border-t py-4 px-[25px] bg-transparent hover:bg-zinc-300/50 duration-200"
                      >
                        <Avatar className="size-10">
                          <AvatarFallback className="capitalize">
                            {activity.user_created[0]}
                          </AvatarFallback>
                        </Avatar>

                        <div className="flex flex-col w-full">
                          <div className="flex items-center justify-between">
                            <p className="font-medium capitalize truncate">
                              {activity.user_created}
                            </p>
                            <p className="text-xs">
                              {dayjs(activity.created_at).format(
                                "DD/MM/YYYY, HH:mm:ss"
                              )}{" "}
                              WIB
                            </p>
                          </div>

                          <div
                            className="text-xs opacity-80 font-light description-detail"
                            dangerouslySetInnerHTML={{
                              __html: activity.message,
                            }}
                          />
                        </div>
                      </div>
                    )
                  )
                ) : (
                  <></>
                )}
              </div>
            </div>

            <form
              id="floating-form"
              onSubmit={handleSubmit}
              className="flex p-[14px] fixed bottom-0 border-t items-end gap-2 w-xl z-10 bg-white"
            >
              <div className="flex flex-col gap-2 w-full">
                <TextEditor
                  value={message}
                  onValueChange={(newVal) =>
                    setMessage(newVal === "<p></p>" ? "" : newVal)
                  }
                  editorClassName="max-h-[80px] overflow-auto"
                  refetch={refetch}
                  disabled={loadingCreate}
                />
              </div>

              <Button className="h-[34px]" disabled={!message || loadingCreate}>
                <p className="leading-none">Send</p>
              </Button>
            </form>
          </div>
        )}
      </DrawerContent>
    </Drawer>
  );
}
