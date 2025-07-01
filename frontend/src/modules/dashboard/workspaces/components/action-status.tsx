import DialogConfirmation from "@/components/layout/dialog-confirmation";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { useWorkspaceStatus } from "@/hooks/use-workspace-status";
import { WorkspaceStatus } from "@/types/workspace";
import {
  Archive,
  ArchiveRestore,
  Ellipsis,
  Pencil,
  Trash2,
} from "lucide-react";
import React, { useState } from "react";

interface ActionStatusProps {
  onClickUpdate?: () => void;
  status: WorkspaceStatus;
  isArchived: boolean;
}

export default function ActionStatus({
  onClickUpdate = () => {},
  status,
  isArchived = false,
}: ActionStatusProps) {
  const [openPopover, setOpenPopover] = useState(false);
  const [openDialog, setOpenDialog] = useState(false);

  const {
    archiveWorkspaceStatus,
    unarchiveWorkspaceStatus,
    deleteWorkspaceStatus,
  } = useWorkspaceStatus();
  const { mutateAsync: archive } = archiveWorkspaceStatus();
  const { mutateAsync: unarchive } = unarchiveWorkspaceStatus();
  const { mutateAsync: deleteStatus, isPending: loadingDelete } =
    deleteWorkspaceStatus();

  const handleArchive = async () => {
    try {
      await archive(status.id);
    } catch (error) {
      console.error(error);
    }

    setOpenPopover(false);
  };

  const handleUnarchive = async () => {
    try {
      await unarchive(status.id);
    } catch (error) {
      console.error(error);
    }

    setOpenPopover(false);
  };

  const handleDelete = async () => {
    try {
      await deleteStatus(status.id);
    } catch (error) {
      console.error(error);
    }

    setOpenPopover(false);
  };

  return (
    <Popover open={openPopover} onOpenChange={setOpenPopover}>
      <PopoverTrigger className="size-8 flex items-center justify-center cursor-pointer">
        <Ellipsis className="size-4 text-zinc-500" />
      </PopoverTrigger>
      <PopoverContent className="w-fit flex flex-col p-1">
        {isArchived ? (
          <>
            <div
              className="flex items-center gap-2 hover:bg-zinc-50 py-2 px-3 rounded cursor-pointer"
              onClick={handleUnarchive}
            >
              <ArchiveRestore className="size-4" />
              <p className="text-sm">Restore</p>
            </div>

            <div
              className="flex items-center gap-2 hover:bg-zinc-50 py-2 px-3 rounded cursor-pointer"
              onClick={handleDelete}
            >
              <Trash2 className="size-4" />
              <p className="text-sm">Delete</p>
            </div>
          </>
        ) : (
          <>
            <div
              className="flex items-center gap-2 hover:bg-zinc-50 py-2 px-3 rounded cursor-pointer"
              onClick={onClickUpdate}
            >
              <Pencil className="size-4" />
              <p className="text-sm">Update</p>
            </div>
            <div
              className="flex items-center gap-2 hover:bg-zinc-50 py-2 px-3 rounded cursor-pointer"
              onClick={handleArchive}
            >
              <Archive className="size-4" />
              <p className="text-sm">Archive</p>
            </div>
          </>
        )}
      </PopoverContent>
    </Popover>
  );
}
