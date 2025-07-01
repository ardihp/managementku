import DialogConfirmation from "@/components/layout/dialog-confirmation";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { useWorkspaceStatus } from "@/hooks/use-workspace-status";
import { WorkspaceStatus } from "@/types/workspace";
import { Archive, Ellipsis, Pencil } from "lucide-react";
import React, { useState } from "react";

interface ActionStatusProps {
  onClickUpdate: () => void;
  status: WorkspaceStatus;
}

export default function ActionStatus({
  onClickUpdate,
  status,
}: ActionStatusProps) {
  const [openPopover, setOpenPopover] = useState(false);
  const [openDialog, setOpenDialog] = useState(false);

  const { archiveWorkspaceStatus } = useWorkspaceStatus();
  const { mutateAsync: archive, isPending } = archiveWorkspaceStatus();

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    try {
      await archive(status.id);
    } catch (error) {
      console.error(error);
    }

    setOpenDialog(false);
  };

  return (
    <>
      <DialogConfirmation
        title={"Confirm archive status"}
        description="Are you sure you want to archive this status? This action
            cannot be undone."
        open={openDialog}
        onOpenChange={() => setOpenDialog(false)}
        isPending={isPending}
        onSubmit={handleSubmit}
      />

      <Popover open={openPopover} onOpenChange={setOpenPopover}>
        <PopoverTrigger className="size-8 flex items-center justify-center cursor-pointer">
          <Ellipsis className="size-4 text-zinc-500" />
        </PopoverTrigger>
        <PopoverContent className="w-fit flex flex-col p-1">
          <div
            className="flex items-center gap-2 hover:bg-zinc-50 py-2 px-3 rounded cursor-pointer"
            onClick={onClickUpdate}
          >
            <Pencil className="size-4" />
            <p className="text-sm">Update</p>
          </div>
          <div
            className="flex items-center gap-2 hover:bg-zinc-50 py-2 px-3 rounded cursor-pointer"
            onClick={() => setOpenDialog(true)}
          >
            <Archive className="size-4" />
            <p className="text-sm">Archive</p>
          </div>
        </PopoverContent>
      </Popover>
    </>
  );
}
