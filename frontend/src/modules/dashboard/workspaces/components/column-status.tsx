import ColorfulBadge from "@/components/layout/colorful-badge";
import {
  Workspace,
  WorkspaceByStatus,
  WorkspaceStatus,
} from "@/types/workspace";
import { Draggable, Droppable } from "@hello-pangea/dnd";
import { Plus } from "lucide-react";
import React from "react";
import { cn } from "@/lib/utils";
import WorkspaceCard from "./card";
import FormWorkspace from "./form";
import ActionStatus from "./action-status";

interface ColumnStatusProps {
  status: WorkspaceByStatus;
  index: number;
  isAdmin: boolean;
  setSelectedWorkspace: (item: Workspace | null) => void;
  setSelectedStatus: (item: WorkspaceStatus | null) => void;
}

export default function ColumnStatus({
  status,
  index,
  isAdmin,
  setSelectedWorkspace,
  setSelectedStatus,
}: ColumnStatusProps) {
  return (
    <Draggable
      draggableId={JSON.stringify({
        ...status,
        workspaces: undefined,
      })}
      index={index}
    >
      {({ dragHandleProps, draggableProps, innerRef }) => {
        return (
          <div
            ref={innerRef}
            className="flex flex-col w-full min-w-[300px] max-w-[300px] h-auto mx-3 relative -left-3"
            {...draggableProps}
          >
            <div
              className="flex items-center justify-between gap-2 rounded-xl bg-white border opacity-100 py-2 px-2 mb-2 relative cursor-grab"
              {...dragHandleProps}
            >
              <div className="flex items-center gap-5">
                <ColorfulBadge
                  content={status.workspaces.length.toString()}
                  color={status.color}
                />

                <p className="font-medium">{status.name}</p>
              </div>

              <ActionStatus
                onClickUpdate={() => setSelectedStatus(status)}
                status={status}
              />
            </div>

            <Droppable
              droppableId={JSON.stringify({
                ...status,
                workspaces: undefined,
              })}
              type="CARD"
            >
              {(
                { droppableProps, innerRef, placeholder },
                { isDraggingOver }
              ) => {
                return (
                  <div className="h-full" ref={innerRef} {...droppableProps}>
                    <div className="flex flex-col overflow-y-auto overflow-x-hidden max-h-[calc(100vh_-_308px)]">
                      {status.workspaces.map(
                        (workspace: Workspace, cardKey) => (
                          <WorkspaceCard
                            key={workspace.id}
                            index={cardKey}
                            workspace={workspace}
                            setSelectedWorkspace={setSelectedWorkspace}
                            isArchived={workspace.is_archived}
                            customClassName="my-2"
                          />
                        )
                      )}

                      <div className={cn(isDraggingOver ? "flex" : "hidden")}>
                        {placeholder}
                      </div>
                    </div>

                    {isAdmin && (
                      <FormWorkspace
                        resetSelectedWorkspace={() =>
                          setSelectedWorkspace(null)
                        }
                        defaultStatus={status.id.toString()}
                        buttonClassName="flex items-center gap-2 rounded-lg bg-white border opacity-100 py-2 px-4 cursor-pointer hover:bg-zinc-50 transition-colors w-full mt-2"
                      >
                        <Plus className="size-4" />
                        <p className="text-sm">Add a workspace</p>
                      </FormWorkspace>
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
