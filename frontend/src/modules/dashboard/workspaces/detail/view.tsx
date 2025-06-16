"use client";

import React, { useEffect, useState } from "react";
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
import { Plus } from "lucide-react";
import ColumnStatus from "./components/column-status";
import { WorkspaceStatus } from "@/types/workspace";
import { CardOrderItem } from "@/types/hooks/workspace-card";

export default function DetailWorkspaceView() {
  const params = useParams();
  const { user } = useAuth();
  const { getDetailWorkspace } = useWorkspace();
  const {
    getListWorkspaceCard,
    updateWorkspaceCardStatusOrder,
    updateWorkspaceCardOrder,
  } = useWorkspaceCard();
  const [selectedItem, setSelectedItem] = useState<WorkspaceCard | null>(null);
  const [listWorkspaceCardByStatus, setListWorkspaceCardByStatus] = useState<
    WorkspaceCardByStatus[]
  >([]);

  const { data: detailWorkspace, isLoading: loadingWorkspace } =
    getDetailWorkspace(Number(params?.id || 0));
  const {
    data: listWorkspaceCard,
    isLoading: loadingWorkspaceCard,
    isFetching: fetchWorkspaceCard,
  } = getListWorkspaceCard(Number(params?.id || 0));
  const { mutate: updateStatusOrder } = updateWorkspaceCardStatusOrder();
  const { mutate: updateCardOrder } = updateWorkspaceCardOrder();

  const isMember =
    detailWorkspace?.members.find((member) => member.id === user?.id)?.role ===
    "Member";

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
            id: Number(params?.id || 0),
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
            id: Number(params?.id || 0),
            data: { order_status: statusWorkspace },
          },
          {
            onError: () =>
              setListWorkspaceCardByStatus(listWorkspaceCardByStatus),
          }
        );
      }
    }
  };

  return (
    <div className="flex flex-col h-full">
      <section className="flex items-center justify-between border-b border-zinc-200 p-8">
        <div className="h-9 flex items-center gap-1.5">
          <Link href={paths.dashboard.workspace}>
            <p className="text-2xl text-foreground/40 font-light my-auto">
              Workspace{" "}
            </p>
          </Link>

          <div className="text-2xl text-foreground my-auto flex items-center gap-1">
            /{" "}
            {loadingWorkspace ? (
              <Skeleton className="w-60 h-7" />
            ) : (
              <span>{detailWorkspace?.name}</span>
            )}
          </div>
        </div>
      </section>

      <section className="flex flex-col h-full gap-6">
        {loadingWorkspace || loadingWorkspaceCard ? (
          <section className="flex p-8 gap-6">
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
              {({ droppableProps, innerRef, placeholder }) => {
                return (
                  <div
                    className="flex overflow-auto p-8 h-full -ml-3"
                    {...droppableProps}
                    ref={innerRef}
                  >
                    {listWorkspaceCardByStatus.map(
                      (status: WorkspaceCardByStatus, key: number) => (
                        <ColumnStatus
                          key={status.id}
                          status={status}
                          index={key}
                          isMember={isMember}
                          workspaceId={Number(params?.id || 0)}
                          setSelectedItem={setSelectedItem}
                        />
                      )
                    )}

                    {placeholder}
                  </div>
                );
              }}
            </Droppable>
          </DragDropContext>
        )}
      </section>
    </div>
  );
}
