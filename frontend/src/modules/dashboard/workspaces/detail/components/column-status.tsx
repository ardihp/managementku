import ColorfulBadge from "@/components/layout/colorful-badge";
import { cn } from "@/lib/utils";
import { WorkspaceCard, WorkspaceCardByStatus } from "@/types/workspace-card";
import { Draggable, Droppable } from "@hello-pangea/dnd";
import React from "react";
import FormWorkspaceCard from "./form";
import { GripVertical, Plus } from "lucide-react";
import TaskCard from "./card";

interface ColumnStatusProps {
  status: WorkspaceCardByStatus;
  index: number;
  isAdmin: boolean;
  workspaceId: number;
  setSelectedItem: (item: null) => void;
}

export default function ColumnStatus({
  status,
  index,
  isAdmin,
  workspaceId,
  setSelectedItem,
}: ColumnStatusProps) {
  return (
    <Draggable
      draggableId={JSON.stringify({
        ...status,
        cards: undefined,
      })}
      index={index}
    >
      {({ dragHandleProps, draggableProps, innerRef }) => {
        return (
          <div
            ref={innerRef}
            className="flex flex-col gap-4 w-full min-w-[300px] max-w-[300px] mx-3"
            {...draggableProps}
          >
            <div
              className="flex items-center justify-between gap-2 rounded-xl bg-white border opacity-100 py-2 px-4 relative cursor-grab"
              {...dragHandleProps}
            >
              <GripVertical className="absolute left-2 size-5 text-zinc-400" />

              <p className="ml-5 font-medium">{status.name}</p>

              <ColorfulBadge
                content={status.cards.length.toString()}
                color={status.color}
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
                        <TaskCard
                          key={card.id}
                          index={cardKey}
                          card={card}
                          customClassName="my-2"
                        />
                      ))}

                      <div className={cn(isDraggingOver ? "flex" : "hidden")}>
                        {placeholder}
                      </div>
                    </div>

                    {isAdmin && (
                      <FormWorkspaceCard
                        resetSelectedItem={() => setSelectedItem(null)}
                        defaultStatus={status.id.toString()}
                        workspaceId={workspaceId}
                        buttonClassName="flex items-center gap-2 rounded-lg bg-white border opacity-100 py-2 px-4 cursor-pointer hover:bg-zinc-50 transition-colors w-full mt-2"
                      >
                        <Plus className="size-4" />
                        <p className="text-sm">Add New</p>
                      </FormWorkspaceCard>
                    )}
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
