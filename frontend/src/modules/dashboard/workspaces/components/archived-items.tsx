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
import { useWorkspaceCard } from "@/hooks/use-workspace-card";
import { useWorkspaceCardStatus } from "@/hooks/use-workspace-card-status";
import TaskCard from "../detail/components/card";
import { WorkspaceCard } from "@/types/workspace-card";

interface ArchivedItemsProps {
  open: boolean;
  onOpenChange: () => void;
  type: "workspace" | "workspace-detail";
  workspaceId?: number;
}

export default function ArchivedItems({
  open,
  onOpenChange,
  type,
  workspaceId = 0,
}: ArchivedItemsProps) {
  // type workspace
  const { getListArchivedWorkspace } = useWorkspace();
  const { getListArchivedWorkspaceStatus } = useWorkspaceStatus();

  const { data: archivedWorkspaces } = getListArchivedWorkspace();
  const { data: archivedStatus } = getListArchivedWorkspaceStatus();

  // type workspace-detail
  const { getListArchivedWorkspaceCard } = useWorkspaceCard();
  const { getListArchivedWorkspaceCardStatus } = useWorkspaceCardStatus();

  const { data: archivedCards } = getListArchivedWorkspaceCard(workspaceId);
  const { data: archivedCardStatus } =
    getListArchivedWorkspaceCardStatus(workspaceId);

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Archived Items</DialogTitle>
          <DialogDescription>
            List of cards or status that have been archived
          </DialogDescription>

          <Tabs
            defaultValue={type === "workspace" ? "workspace" : "card"}
            className="w-auto mt-2 gap-4"
          >
            <TabsList>
              {type === "workspace" && (
                <TabsTrigger value="workspace">Workspace</TabsTrigger>
              )}
              {type === "workspace-detail" && (
                <TabsTrigger value="card">Card</TabsTrigger>
              )}
              <TabsTrigger value="status">Status</TabsTrigger>
            </TabsList>
            <TabsContent
              value="card"
              className="flex flex-col gap-3 max-h-[315px] overflow-auto"
            >
              {archivedCards?.length >= 1 ? (
                archivedCards.map((card: WorkspaceCard, key: number) => (
                  <TaskCard
                    key={key}
                    card={card}
                    isArchived={card.is_archived}
                    workspaceId={workspaceId}
                  />
                ))
              ) : (
                <p>No workspace archived</p>
              )}
            </TabsContent>
            <TabsContent
              value="workspace"
              className="flex flex-col gap-3 max-h-[315px] overflow-auto"
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
              {(type === "workspace" ? archivedStatus : archivedCardStatus)
                ?.length >= 1 ? (
                (type === "workspace"
                  ? archivedStatus
                  : archivedCardStatus
                ).map((status: WorkspaceStatus, key: number) => (
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
                      type={type}
                      workspaceId={workspaceId}
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
