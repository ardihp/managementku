"use client";

import React, { useContext, useEffect, useState } from "react";
import LayoutPage from "@/components/layout";
import { useWorkspace } from "@/hooks/use-workspace";
import { useParams } from "next/navigation";
import Link from "next/link";
import { paths } from "@/const/routes";
import { Skeleton } from "@/components/ui/skeleton";
import { useWorkspaceCard } from "@/hooks/use-workspace-card";
import { WorkspaceCard, WorkspaceCardByStatus } from "@/types/workspace-card";
import ColorfulBadge from "@/components/layout/colorful-badge";
import { DragDropContext, Draggable, Droppable } from "@hello-pangea/dnd";
import { cn } from "@/lib/utils";
import { useAuth } from "@/hooks/auth/use-auth";
import FormWorkspaceCard from "./components/form";
import { Archive, History, Plus } from "lucide-react";
import ColumnStatus from "./components/column-status";
import { WorkspaceStatus } from "@/types/workspace";
import { CardOrderItem } from "@/types/hooks/workspace-card";
import FormWorkspaceStatus from "../components/form-status";
import { Button } from "@/components/ui/button";
import ArchivedItems from "../components/archived-items";
import HistoryLog from "./components/history-log";
import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import { SocketContext } from "@/context/socket";

export default function DetailWorkspaceView() {
  const params = useParams();
  const { user } = useAuth();
  const { getDetailWorkspace } = useWorkspace();
  const {
    getListWorkspaceCard,
    updateWorkspaceCardStatusOrder,
    updateWorkspaceCardOrder,
  } = useWorkspaceCard();
  const [selectedCard, setSelectedCard] = useState<WorkspaceCard | null>(null);
  const [selectedStatus, setSelectedStatus] = useState<WorkspaceStatus | null>(
    null
  );
  const [listWorkspaceCardByStatus, setListWorkspaceCardByStatus] = useState<
    WorkspaceCardByStatus[]
  >([]);
  const [openDialog, setOpenDialog] = useState({
    card: false,
    status: false,
    archived: false,
    history: false,
  });
  const { socket } = useContext(SocketContext);

  const { data: detailWorkspace, isLoading: loadingWorkspace } =
    getDetailWorkspace(Number(params?.id || 0));
  const {
    data: listWorkspaceCard,
    isLoading: loadingWorkspaceCard,
    isFetching: fetchWorkspaceCard,
  } = getListWorkspaceCard(Number(params?.id || 0));
  const { mutate: updateStatusOrder } = updateWorkspaceCardStatusOrder(
    Number(params?.id || 0)
  );
  const { mutate: updateCardOrder } = updateWorkspaceCardOrder(
    Number(params?.id || 0)
  );

  const isLoading = loadingWorkspace || loadingWorkspaceCard || !params?.id;

  const isAdmin =
    detailWorkspace?.members.find((member) => member.id === user?.id)?.role ===
    "admin";

  useEffect(() => {
    if (!loadingWorkspaceCard && !fetchWorkspaceCard) {
      setListWorkspaceCardByStatus(listWorkspaceCard || []);
    }
  }, [loadingWorkspaceCard, fetchWorkspaceCard]);

  const handleDragEnd = (item: any) => {
    if (!item.destination) return;

    if (item.type === "STATUS") {
      const draggableItem: WorkspaceStatus = JSON.parse(item.draggableId);

      const sourceWorkspace: WorkspaceStatus[] = Array.from(
        listWorkspaceCardByStatus.map((item) => {
          return { ...item, cards: undefined };
        })
      );
      sourceWorkspace.splice(item.source.index, 1);
      sourceWorkspace.splice(item.destination.index, 0, draggableItem);

      setListWorkspaceCardByStatus((prevVal) => [
        ...sourceWorkspace.map((item) => {
          return {
            ...item,
            cards: prevVal.find((status) => status.id === item.id)?.cards!,
          };
        }),
      ]);

      if (
        item.source.droppableId !== item.destination.droppableId ||
        item.source.index !== item.destination.index
      ) {
        updateStatusOrder(
          {
            data: { order_status: sourceWorkspace.map((item) => item.id) },
          },
          {
            onError: () => {
              setListWorkspaceCardByStatus(listWorkspaceCardByStatus);
            },
          }
        );
      }
    } else {
      const statusWorkspace: CardOrderItem[] = [];
      const source: { droppableId: WorkspaceStatus; index: number } = {
        droppableId: JSON.parse(item.source?.droppableId),
        index: item.source?.index,
      };
      const destination: { droppableId: WorkspaceStatus; index: number } = {
        droppableId: item.destination?.droppableId
          ? JSON.parse(item.destination?.droppableId)
          : "",
        index: item.destination?.index,
      };

      const sourceWorkspace = Array.from(
        listWorkspaceCardByStatus.find(
          (item) => item.id === source.droppableId.id
        )!.cards
      );

      const draggableItem = sourceWorkspace.find(
        (workspace) => workspace.id.toString() === item.draggableId
      )!;

      sourceWorkspace.splice(source.index, 1);

      if (destination.droppableId.name === source.droppableId.name) {
        sourceWorkspace.splice(destination.index, 0, draggableItem);
      }

      statusWorkspace.push({
        status_id: source.droppableId.id,
        card_id: sourceWorkspace.map((item) => item.id),
      });

      setListWorkspaceCardByStatus((prevVal) => [
        ...prevVal.map((status) => {
          if (status.id === source.droppableId.id) {
            return { ...status, cards: sourceWorkspace };
          } else {
            return { ...status };
          }
        }),
      ]);

      if (source.droppableId.name !== destination.droppableId.name) {
        const destinationWorkspace = Array.from(
          listWorkspaceCardByStatus.find(
            (item) => item.id === destination.droppableId.id
          )!.cards
        );

        destinationWorkspace.splice(destination.index, 0, {
          ...draggableItem,
          status: destination.droppableId,
        });

        statusWorkspace.push({
          status_id: destination.droppableId.id,
          card_id: destinationWorkspace.map((item) => item.id),
        });

        setListWorkspaceCardByStatus((prevVal) => [
          ...prevVal.map((status) => {
            if (status.id === destination.droppableId.id) {
              return { ...status, cards: destinationWorkspace };
            } else {
              return { ...status };
            }
          }),
        ]);
      }

      if (
        source.droppableId.id !== destination.droppableId.id ||
        source.index !== destination.index
      ) {
        updateCardOrder(
          {
            data: {
              card_id: Number(item.draggableId),
              order_status: statusWorkspace,
            },
          },
          {
            onError: () =>
              setListWorkspaceCardByStatus(listWorkspaceCardByStatus),
          }
        );
      }
    }

    socket?.send(
      JSON.stringify({
        type: "detail-card",
        workspace_id: detailWorkspace?.id,
      })
    );
  };

  return (
    <>
      {openDialog.status && (
        <FormWorkspaceStatus
          open={openDialog.status}
          onOpenChange={() => {
            setOpenDialog((prevVal) => ({ ...prevVal, status: false }));
            setSelectedStatus(null);
          }}
          status={selectedStatus}
          type="workspace-detail"
          workspaceId={Number(params?.id || 0)}
        />
      )}

      {openDialog.archived && (
        <ArchivedItems
          open={openDialog.archived}
          onOpenChange={() =>
            setOpenDialog((prevVal) => ({ ...prevVal, archived: false }))
          }
          type="workspace-detail"
          workspaceId={Number(params?.id || 0)}
        />
      )}

      {openDialog.history && (
        <HistoryLog
          open={openDialog.history}
          onOpenChange={() =>
            setOpenDialog((prevVal) => ({ ...prevVal, history: false }))
          }
          workspaceId={Number(params?.id || 0)}
        />
      )}

      <FormWorkspaceCard
        workspaceId={Number(params?.id || 0)}
        selectedItem={selectedCard}
        resetSelectedItem={() => setSelectedCard(null)}
      >
        <></>
      </FormWorkspaceCard>

      <div className="flex flex-col h-full">
        <section className="flex items-center justify-between border-b border-zinc-200 p-[28px]">
          <div className="h-9 flex items-center gap-1.5">
            <Link href={paths.dashboard.workspace}>
              <p className="text-2xl text-foreground/40 font-light my-auto">
                Workspace{" "}
              </p>
            </Link>

            <div className="text-2xl text-foreground my-auto flex items-center gap-1">
              /{" "}
              {isLoading ? (
                <Skeleton className="w-60 h-7" />
              ) : (
                <span>{detailWorkspace?.name}</span>
              )}
            </div>
          </div>

          <div className="flex h-full items-center gap-4">
            <div className="*:data-[slot=avatar]:ring-zinc-200 *:data-[slot=avatar]:bg-zinc-100 flex -space-x-2 *:data-[slot=avatar]:ring-2 mt-auto ml-auto">
              {detailWorkspace?.members.map((member, key) => (
                <Avatar key={key} title={member.username} className="size-8">
                  <AvatarFallback className="uppercase text-sm pt-0.5">
                    {member.username?.split(" ")?.map((text) => text?.[0])}
                  </AvatarFallback>
                </Avatar>
              ))}
            </div>

            <div className="w-[1px] bg-zinc-300 h-5 ml-4" />

            <Button
              className="cursor-pointer"
              onClick={() =>
                setOpenDialog((prevVal) => ({ ...prevVal, archived: true }))
              }
            >
              <Archive className="size-4" />
              <p className="leading-none">Archived Items</p>
            </Button>

            <Button
              className="cursor-pointer"
              onClick={() =>
                setOpenDialog((prevVal) => ({ ...prevVal, history: true }))
              }
            >
              <History className="size-4" />
              <p className="leading-none">History Logs</p>
            </Button>
          </div>
        </section>

        <section className="flex flex-col h-full gap-6">
          {isLoading ? (
            <section className="flex p-[28px] gap-6">
              {[...Array(3)].map((item, key) => (
                <div key={key} className="flex flex-col gap-4">
                  <Skeleton className="w-full min-w-[300px] max-w-[300px] h-[45px] rounded-xl" />

                  {[...Array(key + 1)].map((item, key2) => (
                    <Skeleton
                      key={key2}
                      className="w-full min-w-[300px] max-w-[300px] h-[177px] rounded-xl"
                    />
                  ))}
                </div>
              ))}
            </section>
          ) : (
            <DragDropContext onDragEnd={handleDragEnd}>
              <Droppable
                droppableId="column"
                direction="horizontal"
                type="STATUS"
              >
                {(
                  { droppableProps, innerRef, placeholder },
                  { isDraggingOver }
                ) => {
                  return (
                    <div
                      className="flex overflow-auto p-[28px] h-full -ml-3"
                      {...droppableProps}
                      ref={innerRef}
                    >
                      {listWorkspaceCardByStatus.map(
                        (status: WorkspaceCardByStatus, key: number) => (
                          <ColumnStatus
                            key={status.id}
                            status={status}
                            index={key}
                            isAdmin={isAdmin}
                            workspaceId={Number(params?.id || 0)}
                            setSelectedCard={setSelectedCard}
                            setSelectedStatus={(val) => {
                              setSelectedStatus(val);
                              setOpenDialog((prevVal) => ({
                                ...prevVal,
                                status: true,
                              }));
                            }}
                          />
                        )
                      )}

                      <div className={cn(isDraggingOver ? "flex" : "hidden")}>
                        {placeholder}
                      </div>

                      <div
                        className="flex items-center gap-2 rounded-xl border opacity-100 py-2 px-4 relative h-[45px] min-w-[300px] cursor-pointer hover:bg-white transition-colors"
                        onClick={() =>
                          setOpenDialog((prevVal) => ({
                            ...prevVal,
                            status: true,
                          }))
                        }
                      >
                        <Plus className="size-5" />
                        <p className="font-medium">Add another list</p>
                      </div>
                    </div>
                  );
                }}
              </Droppable>
            </DragDropContext>
          )}
        </section>
      </div>
    </>
  );
}
