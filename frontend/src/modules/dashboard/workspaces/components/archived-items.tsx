import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { useWorkspace } from "@/hooks/use-workspace";
import { useWorkspaceStatus } from "@/hooks/use-workspace-status";
import React from "react";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Workspace, WorkspaceStatus } from "@/types/workspace";
import Card from "./card";
import ColorfulBadge from "@/components/layout/colorful-badge";
import ActionStatus from "./action-status";

interface ArchivedItemsProps {
  open: boolean;
  onOpenChange: () => void;
}

export default function ArchivedItems({
  open,
  onOpenChange,
}: ArchivedItemsProps) {
  const { getListArchivedWorkspace } = useWorkspace();
  const { getListArchivedWorkspaceStatus } = useWorkspaceStatus();

  const { data: archivedWorkspaces, isLoading: loadingWorkspace } =
    getListArchivedWorkspace();
  const { data: archivedStatus, isLoading: loadingStatus } =
    getListArchivedWorkspaceStatus();

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Archived Items</DialogTitle>
          <DialogDescription>
            List of cards or status that have been archived
          </DialogDescription>

          <Tabs defaultValue="workspace" className="w-auto mt-2 gap-4">
            <TabsList>
              <TabsTrigger value="workspace">Workspace</TabsTrigger>
              <TabsTrigger value="status">Status</TabsTrigger>
            </TabsList>
            <TabsContent
              value="workspace"
              className="flex flex-col gap-3 max-h-[315px]"
            >
              {archivedWorkspaces?.length >= 1 ? (
                archivedWorkspaces.map((workspace: Workspace, key: number) => (
                  <Card
                    key={key}
                    workspace={workspace}
                    customClassName=""
                    isArchived
                  />
                ))
              ) : (
                <p>No workspace archived</p>
              )}
            </TabsContent>
            <TabsContent
              value="status"
              className="flex flex-col gap-2 max-h-[315px]"
            >
              {archivedStatus?.length >= 1 ? (
                archivedStatus.map((status: WorkspaceStatus, key: number) => (
                  <div
                    key={key}
                    className="flex items-center justify-between gap-2 rounded-xl bg-white border opacity-100 py-2 px-2 relative"
                  >
                    <div className="flex items-center gap-5">
                      <ColorfulBadge content={"-"} color={status.color} />

                      <p className="font-medium">{status.name}</p>
                    </div>

                    <ActionStatus
                      status={status}
                      isArchived={status.is_archived}
                    />
                  </div>
                ))
              ) : (
                <p>No status archived</p>
              )}
            </TabsContent>
          </Tabs>
        </DialogHeader>
      </DialogContent>
    </Dialog>
  );
}
