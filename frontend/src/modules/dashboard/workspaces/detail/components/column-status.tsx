import ColorfulBadge from "@/components/layout/colorful-badge";
import { cn } from "@/lib/utils";
import { WorkspaceCard, WorkspaceCardByStatus } from "@/types/workspace-card";
import { Draggable, Droppable } from "@hello-pangea/dnd";
import React from "react";
import FormWorkspaceCard from "./form";
import { GripVertical, Plus } from "lucide-react";
import TaskCard from "./card";
import ActionStatus from "../../components/action-status";
import { WorkspaceStatus } from "@/types/workspace";

interface ColumnStatusProps {
  status: WorkspaceCardByStatus;
  index: number;
  isAdmin: boolean;
  workspaceId: number;
  setSelectedCard: (item: WorkspaceCard | null) => void;
  setSelectedStatus: (item: WorkspaceStatus | null) => void;
}

export default function ColumnStatus({
  status,
  index,
  isAdmin,
  workspaceId,
  setSelectedCard,
  setSelectedStatus,
}: ColumnStatusProps) {
  return (
    <Draggable
      draggableId={JSON.stringify({
        ...status,
        cards: undefined,
      })}
      index={index}
      isDragDisabled={!isAdmin}
    >
      {({ dragHandleProps, draggableProps, innerRef }) => {
        return (
          <div
            ref={innerRef}
            className="flex flex-col w-full min-w-[300px] max-w-[300px] mx-3"
            {...draggableProps}
          >
            <div
              className="flex items-center justify-between gap-2 mb-2 rounded-xl bg-white border opacity-100 p-2 relative cursor-grab"
              {...dragHandleProps}
            >
              <div className="flex items-center gap-5">
                <ColorfulBadge
                  content={status.cards.length.toString()}
                  color={status.color}
                />

                <p className="font-medium">{status.name}</p>
              </div>

              <ActionStatus
                onClickUpdate={() => setSelectedStatus(status)}
                status={status}
                isArchived={status.is_archived}
                type="workspace-detail"
                workspaceId={workspaceId}
              />
            </div>

            <Droppable
              droppableId={JSON.stringify({
                ...status,
                cards: undefined,
              })}
              type="CARD"
            >
              {(
                { droppableProps, innerRef, placeholder },
                { isDraggingOver }
              ) => {
                return (
                  <div className="h-full" ref={innerRef} {...droppableProps}>
                    <div className="flex flex-col overflow-y-auto overflow-x-hidden max-h-[calc(100vh_-_312px)]">
                      {status.cards.map((card: WorkspaceCard, cardKey) => (
                        <Draggable
                          key={card.id}
                          draggableId={card.id.toString()}
                          index={cardKey}
                        >
                          {(
                            { dragHandleProps, draggableProps, innerRef },
                            { isDragging }
                          ) => {
                            return (
                              <TaskCard
                                card={card}
                                workspaceId={workspaceId}
                                customClassName="my-2"
                                isArchived={card.is_archived}
                                isDragging={isDragging}
                                setSelectedCard={setSelectedCard}
                                ref={innerRef}
                                {...dragHandleProps}
                                {...draggableProps}
                              />
                            );
                          }}
                        </Draggable>
                      ))}

                      <div className={cn(isDraggingOver ? "flex" : "hidden")}>
                        {placeholder}
                      </div>
                    </div>

                    <FormWorkspaceCard
                      resetSelectedItem={() => setSelectedCard(null)}
                      defaultStatus={status.id.toString()}
                      workspaceId={workspaceId}
                      buttonClassName="flex items-center gap-2 rounded-lg bg-white border opacity-100 py-2 px-4 cursor-pointer hover:bg-zinc-50 transition-colors w-full mt-2"
                    >
                      <Plus className="size-4" />
                      <p className="text-sm">Add New</p>
                    </FormWorkspaceCard>
                  </div>
                );
              }}
            </Droppable>
          </div>
        );
      }}
    </Draggable>
  );
}
