import React, { useEffect, useState } from "react";
import { useAuth, UserDetail } from "@/hooks/auth/use-auth";
import { useWorkspace } from "@/hooks/use-workspace";
import {
  Workspace,
  WorkspaceByStatus,
  WorkspaceStatus,
} from "@/types/workspace";
import { Skeleton } from "@/components/ui/skeleton";
import FormWorkspace from "./components/form";
import { Archive, Plus } from "lucide-react";
import WorkspaceCard from "./components/card";
import ColorfulBadge from "@/components/layout/colorful-badge";
import { DragDropContext, Droppable } from "@hello-pangea/dnd";
import { OrderStatusItem, UpdateWorkspaceOrder } from "@/types/hooks/workspace";
import confetti from "canvas-confetti";
import { cn, handleConffeti } from "@/lib/utils";

export default function WorkspacesView() {
  const [userDetail, setUserDetail] = useState<UserDetail | null>();
  const [selectedItem, setSelectedItem] = useState<Workspace | null>(null);
  const [listWorkspaceByStatus, setListWorkspaceByStatus] = useState<
    WorkspaceByStatus[]
  >([]);

  const { user } = useAuth();
  const { getListWorkspace, updateWorkspaceOrder } = useWorkspace();

  const {
    data: listWorkspace,
    isLoading: loadingWorkspace,
    isFetching: fetchWorkspace,
  } = getListWorkspace();
  const { mutate: updateOrder } = updateWorkspaceOrder();

  useEffect(() => {
    setUserDetail(user);
  }, []);

  useEffect(() => {
    if (!loadingWorkspace && !fetchWorkspace) {
      setListWorkspaceByStatus(listWorkspace);
    }
  }, [loadingWorkspace, fetchWorkspace]);

  const handleDragEnd = (item: any) => {
    if (!item.destination) return;

    const statusWorkspace: OrderStatusItem[] = [];
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
      listWorkspaceByStatus.find((item) => item.id === source.droppableId.id)!
        .workspaces
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
      workspace_id: sourceWorkspace.map((item) => item.id),
    });

    setListWorkspaceByStatus((prevVal) => [
      ...prevVal.map((status) => {
        if (status.id === source.droppableId.id) {
          return { ...status, workspaces: sourceWorkspace };
        } else {
          return { ...status };
        }
      }),
    ]);

    if (source.droppableId.name !== destination.droppableId.name) {
      const destinationWorkspace = Array.from(
        listWorkspaceByStatus.find(
          (item) => item.id === destination.droppableId.id
        )!.workspaces
      );

      destinationWorkspace.splice(destination.index, 0, {
        ...draggableItem,
        status: destination.droppableId,
      });

      statusWorkspace.push({
        status_id: destination.droppableId.id,
        workspace_id: destinationWorkspace.map((item) => item.id),
      });

      setListWorkspaceByStatus((prevVal) => [
        ...prevVal.map((status) => {
          if (status.id === destination.droppableId.id) {
            return { ...status, workspaces: destinationWorkspace };
          } else {
            return { ...status };
          }
        }),
      ]);
    }

    if (destination.droppableId.name === "Completed") handleConffeti();

    if (
      source.droppableId.id !== destination.droppableId.id ||
      source.index !== destination.index
    ) {
      updateOrder(
        { order_status: statusWorkspace },
        { onError: () => setListWorkspaceByStatus(listWorkspaceByStatus) }
      );
    }
  };

  return (
    <div className="flex flex-col h-full">
      <section className="flex items-center justify-between border-b border-zinc-200 p-8">
        <div className="h-9 flex flex-col">
          <p className="text-2xl text-foreground/40 font-light my-auto">
            Workspace{" "}
            <span className="text-foreground font-normal">/ All Workspace</span>
          </p>
        </div>

        <div className="flex">
          {userDetail?.role === "manager" && (
            <FormWorkspace
              selectedItem={selectedItem}
              resetSelectedItem={() => setSelectedItem(null)}
              buttonClassName="flex items-center gap-2 bg-primary py-2 px-3 rounded-md cursor-pointer h-9"
            >
              <Plus className="size-4 text-white" />
              <p className="text-white text-sm leading-none">Add New</p>
            </FormWorkspace>
          )}
        </div>
      </section>

      {loadingWorkspace ? (
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
      ) : listWorkspaceByStatus?.length >= 1 ? (
        <DragDropContext onDragEnd={handleDragEnd}>
          <section className="flex p-8 gap-6">
            {listWorkspaceByStatus?.map(
              (status: WorkspaceByStatus, key: number) => (
                <div
                  key={key}
                  className="flex flex-col w-full min-w-[300px] max-w-[300px] h-auto"
                >
                  <div className="flex items-center justify-between gap-2 rounded-xl bg-white border opacity-100 py-2 px-4 mb-2">
                    <p className="font-medium">{status.name}</p>

                    <ColorfulBadge
                      content={status.workspaces.length.toString()}
                      color={status.color}
                    />
                  </div>

                  <Droppable
                    droppableId={JSON.stringify({
                      ...status,
                      workspaces: undefined,
                    })}
                  >
                    {(
                      { droppableProps, innerRef, placeholder },
                      { isDraggingOver }
                    ) => {
                      return (
                        <div
                          className="h-full"
                          ref={innerRef}
                          {...droppableProps}
                        >
                          <div className="flex flex-col overflow-y-auto overflow-x-hidden max-h-[calc(100vh_-_308px)]">
                            {status.workspaces.map(
                              (workspace: Workspace, cardKey) => (
                                <WorkspaceCard
                                  key={workspace.id}
                                  index={cardKey}
                                  workspace={workspace}
                                  setSelectedItem={setSelectedItem}
                                  isArchived={workspace.is_archived}
                                  customClassName="my-2"
                                />
                              )
                            )}

                            <div
                              className={cn(isDraggingOver ? "flex" : "hidden")}
                            >
                              {placeholder}
                            </div>
                          </div>

                          {userDetail?.role === "manager" && (
                            <FormWorkspace
                              resetSelectedItem={() => setSelectedItem(null)}
                              defaultStatus={status.id.toString()}
                              buttonClassName="flex items-center gap-2 rounded-lg bg-white border opacity-100 py-2 px-4 cursor-pointer hover:bg-zinc-50 transition-colors w-full mt-2"
                            >
                              <Plus className="size-4" />
                              <p className="text-sm">Add New</p>
                            </FormWorkspace>
                          )}
                        </div>
                      );
                    }}
                  </Droppable>
                </div>
              )
            )}
          </section>
        </DragDropContext>
      ) : (
        <div className="flex flex-col h-full gap-2 items-center justify-center">
          <Archive className="size-10 stroke-[1.5px] text-zinc-700" />

          <p className="text-zinc-700">There's no workspace available</p>
        </div>
      )}
    </div>
  );
}
