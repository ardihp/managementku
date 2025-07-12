import React, { useContext, useEffect, useState } from "react";
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
import { DragDropContext, Droppable } from "@hello-pangea/dnd";
import { OrderStatusItem, UpdateWorkspaceOrder } from "@/types/hooks/workspace";
import { cn, handleConffeti } from "@/lib/utils";
import ColumnStatus from "./components/column-status";
import { useWorkspaceStatus } from "@/hooks/use-workspace-status";
import FormWorkspaceStatus from "./components/form-status";
import { Button } from "@/components/ui/button";
import ArchivedItems from "./components/archived-items";
import { SocketContext } from "@/context/socket";

export default function WorkspacesView() {
  const [selectedWorkspace, setSelectedWorkspace] = useState<Workspace | null>(
    null
  );
  const [selectedStatus, setSelectedStatus] = useState<WorkspaceStatus | null>(
    null
  );
  const [loadingPage, setLoadingPage] = useState(false);
  const [openDialog, setOpenDialog] = useState({
    workspace: false,
    status: false,
    archived: false,
  });
  const [listWorkspaceByStatus, setListWorkspaceByStatus] = useState<
    WorkspaceByStatus[]
  >([]);
  const { user } = useAuth();
  const { socket } = useContext(SocketContext);

  const { getListWorkspace, updateWorkspaceOrder } = useWorkspace();
  const { updateWorkspaceStatusOrder } = useWorkspaceStatus();

  const {
    data: listWorkspace,
    isLoading: loadingWorkspace,
    isFetching: fetchWorkspace,
  } = getListWorkspace();

  const { mutate: updateOrder } = updateWorkspaceOrder();
  const { mutate: updateStatusOrder } = updateWorkspaceStatusOrder();

  const isLoading = loadingPage || loadingWorkspace;

  useEffect(() => {
    if (loadingWorkspace) {
      setLoadingPage(true);
    }

    if (!loadingWorkspace && !fetchWorkspace) {
      setListWorkspaceByStatus(listWorkspace);
      setLoadingPage(false);
    }
  }, [loadingWorkspace, fetchWorkspace]);

  const handleDragEnd = (item: any) => {
    if (!item.destination) return;

    if (item.type === "STATUS") {
      const draggableItem: WorkspaceStatus = JSON.parse(item.draggableId);

      const sourceWorkspace: WorkspaceStatus[] = Array.from(
        listWorkspaceByStatus.map((item) => {
          return { ...item, workspaces: undefined };
        })
      );
      sourceWorkspace.splice(item.source.index, 1);
      sourceWorkspace.splice(item.destination.index, 0, draggableItem);

      setListWorkspaceByStatus((prevVal) => [
        ...sourceWorkspace.map((item) => {
          return {
            ...item,
            workspaces: prevVal.find((status) => status.id === item.id)
              ?.workspaces!,
          };
        }),
      ]);

      if (
        item.source.droppableId !== item.destination.droppableId ||
        item.source.index !== item.destination.index
      ) {
        updateStatusOrder(
          { order_status: sourceWorkspace.map((item) => item.id) },
          {
            onError: () => {
              setListWorkspaceByStatus(listWorkspaceByStatus);
            },
          }
        );
      }
    } else {
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
          {
            workspace_id: Number(item?.draggableId),
            order_status: statusWorkspace,
          },
          { onError: () => setListWorkspaceByStatus(listWorkspaceByStatus) }
        );
      }
    }

    socket?.send(JSON.stringify({ type: "refresh-workspace" }));
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
          type="workspace"
        />
      )}

      {openDialog.archived && (
        <ArchivedItems
          open={openDialog.archived}
          onOpenChange={() =>
            setOpenDialog((prevVal) => ({ ...prevVal, archived: false }))
          }
          type="workspace"
        />
      )}

      <div className="flex flex-col h-full">
        <section className="flex items-center justify-between border-b border-zinc-200 p-[28px]">
          <div className="h-9 flex flex-col">
            <p className="text-2xl text-foreground/40 font-light my-auto">
              Workspace{" "}
              <span className="text-foreground font-normal">
                / All Workspace
              </span>
            </p>
          </div>

          {user?.role === "manager" && !isLoading && (
            <div className="flex items-center gap-4">
              <Button
                className="cursor-pointer"
                onClick={() =>
                  setOpenDialog((prevVal) => ({ ...prevVal, archived: true }))
                }
              >
                <Archive className="size-4" />
                <p className="leading-none">Archived Items</p>
              </Button>

              <FormWorkspace
                selectedWorkspace={selectedWorkspace}
                resetSelectedWorkspace={() => setSelectedWorkspace(null)}
                buttonClassName="flex items-center gap-2 bg-primary py-2 px-3 rounded-md cursor-pointer h-9"
              >
                <Plus className="size-4 text-white" />
                <p className="text-white text-sm leading-none">Add New</p>
              </FormWorkspace>
            </div>
          )}
        </section>

        {isLoading ? (
          <section className="flex p-[28px] gap-6">
            {[...Array(3)].map((item, key) => (
              <div key={key} className="flex flex-col gap-4">
                <Skeleton className="w-full min-w-[300px] max-w-[300px] h-[45px] rounded-xl" />

                {[...Array(key + 1)].map((item, key2) => (
                  <Skeleton
                    key={key2}
                    className="w-full min-w-[300px] max-w-[300px] h-[152px] rounded-xl"
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
                    className="flex p-[28px] overflow-auto relative h-full"
                    ref={innerRef}
                    {...droppableProps}
                  >
                    {listWorkspaceByStatus?.map(
                      (status: WorkspaceByStatus, key: number) => (
                        <ColumnStatus
                          key={status.id}
                          index={key}
                          status={status}
                          isAdmin={user?.role === "manager"}
                          setSelectedWorkspace={setSelectedWorkspace}
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

                    {user?.role === "manager" && (
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
                    )}
                  </div>
                );
              }}
            </Droppable>
          </DragDropContext>
        )}
      </div>
    </>
  );
}
