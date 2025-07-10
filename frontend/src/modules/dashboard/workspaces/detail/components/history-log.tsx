import DataTable from "@/components/layout/data-table";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { useWorkspaceCard } from "@/hooks/use-workspace-card";
import React from "react";
import { historyColumns } from "./history-columns";

interface HistoryLogProps {
  open: boolean;
  onOpenChange: () => void;
  workspaceId: number;
}

export default function HistoryLog({
  open,
  onOpenChange,
  workspaceId,
}: HistoryLogProps) {
  const { getListWorkspaceHistoryLog } = useWorkspaceCard();
  const { data: historyLogs, isLoading: loadingHistory } =
    getListWorkspaceHistoryLog(workspaceId);

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="min-w-fit">
        <DialogHeader>
          <DialogTitle>Workspace History Logs</DialogTitle>
          <DialogDescription></DialogDescription>
        </DialogHeader>

        <DataTable
          columns={historyColumns}
          data={historyLogs}
          isLoading={loadingHistory}
          disablePaginationSize
        />
      </DialogContent>
    </Dialog>
  );
}
