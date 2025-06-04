import React, { useState } from "react";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { Archive, Ellipsis, Pencil } from "lucide-react";
import { useWorkspace } from "@/hooks/use-workspace";
import { Workspace } from "@/types/workspace";

interface WorkspaceActionProps {
  setSelectedItem: (item: any) => void;
  workspace: Workspace;
}

export default function WorkspaceAction({
  setSelectedItem,
  workspace,
}: WorkspaceActionProps) {
  const [openPopover, setOpenPopover] = useState(false);

  const { archiveWorkspace } = useWorkspace();
  const { mutate: archive } = archiveWorkspace();

  const handleArchiveWorkspace = (id: number) => {
    archive(id, {
      onSuccess: () => setOpenPopover(false),
    });
  };

  return (
    <Popover onOpenChange={setOpenPopover} open={openPopover}>
      <PopoverTrigger className="absolute top-4 bg-white right-4 z-[10] size-8 flex items-center justify-center cursor-pointer">
        <Ellipsis className="size-4 text-zinc-500" />
      </PopoverTrigger>
      <PopoverContent className="w-fit flex flex-col p-1">
        <div
          className="flex items-center gap-2 hover:bg-zinc-50 py-2 px-3 rounded cursor-pointer"
          onClick={() => setSelectedItem(workspace)}
        >
          <Pencil className="size-4" />
          <p className="text-sm">Update</p>
        </div>
        <div
          className="flex items-center gap-2 hover:bg-zinc-50 py-2 px-3 rounded cursor-pointer"
          onClick={() => handleArchiveWorkspace(workspace.id)}
        >
          <Archive className="size-4" />
          <p className="text-sm">Archive</p>
        </div>
      </PopoverContent>
    </Popover>
  );
}
