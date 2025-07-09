import React, { useState } from "react";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import {
  Archive,
  ArchiveRestore,
  Ellipsis,
  Pencil,
  Trash2,
  Users2,
} from "lucide-react";
import { useWorkspace } from "@/hooks/use-workspace";
import { useWorkspaceCard } from "@/hooks/use-workspace-card";
import { WorkspaceCard } from "@/types/workspace-card";

interface WorkspaceActionProps {
  setSelectedCard?: (item: any) => void;
  card: WorkspaceCard;
  isArchived: boolean;
  workspaceId?: number;
}

export default function WorkspaceAction({
  setSelectedCard = () => {},
  card,
  isArchived = false,
  workspaceId = 0,
}: WorkspaceActionProps) {
  const [openPopover, setOpenPopover] = useState(false);
  const [openDialog, setOpenDialog] = useState({
    member: false,
  });

  const {
    archivedWorkspaceCard,
    unarchivedWorkspaceCard,
    deleteWorkspaceCard,
  } = useWorkspaceCard();

  const { mutateAsync: archiveCard } = archivedWorkspaceCard(workspaceId);
  const { mutateAsync: unarchiveCard } = unarchivedWorkspaceCard(workspaceId);
  const { mutateAsync: deleteCard } = deleteWorkspaceCard(workspaceId);

  const handleArchiveWorkspace = async (id: number) => {
    try {
      await archiveCard(id);
    } catch (error) {
      console.error(error);
    }
    setOpenPopover(false);
  };

  const handleUnarchiveWorkspace = async (id: number) => {
    try {
      await unarchiveCard(id);
    } catch (error) {
      console.error(error);
    }
    setOpenPopover(false);
  };

  const handleDeleteWorkspace = async (id: number) => {
    try {
      await deleteCard(id);
    } catch (error) {
      console.error(error);
    }
    setOpenPopover(false);
  };

  return (
    <>
      {/* {openDialog.member && (
        <AssignMember
          open={openDialog.member}
          onOpenChange={() =>
            setOpenDialog((prevVal) => ({ ...prevVal, member: false }))
          }
          workspace={workspace}
        />
      )} */}

      <Popover onOpenChange={setOpenPopover} open={openPopover}>
        <PopoverTrigger className="absolute top-2 right-2 z-[10] size-8 flex items-center justify-center cursor-pointer">
          <Ellipsis className="size-4 text-zinc-500" />
        </PopoverTrigger>
        <PopoverContent className="w-fit flex flex-col p-1">
          {isArchived ? (
            <>
              <div
                className="flex items-center gap-2 hover:bg-zinc-50 py-2 px-3 rounded cursor-pointer"
                onClick={() => handleUnarchiveWorkspace(card.id)}
              >
                <ArchiveRestore className="size-4" />
                <p className="text-sm">Restore</p>
              </div>
              <div
                className="flex items-center gap-2 hover:bg-zinc-50 py-2 px-3 rounded cursor-pointer"
                onClick={() => handleDeleteWorkspace(card.id)}
              >
                <Trash2 className="size-4" />
                <p className="text-sm">Delete</p>
              </div>
            </>
          ) : (
            <>
              <div
                className="flex items-center gap-2 hover:bg-zinc-50 py-2 px-3 rounded cursor-pointer"
                onClick={() => setSelectedCard(card)}
              >
                <Pencil className="size-4" />
                <p className="text-sm">Update</p>
              </div>
              <div
                className="flex items-center gap-2 hover:bg-zinc-50 py-2 px-3 rounded cursor-pointer"
                onClick={() => {
                  setOpenPopover(false);
                  setOpenDialog((prevVal) => ({ ...prevVal, member: true }));
                }}
              >
                <Users2 className="size-4" />
                <p className="text-sm">Assign Member</p>
              </div>
              <div
                className="flex items-center gap-2 hover:bg-zinc-50 py-2 px-3 rounded cursor-pointer"
                onClick={() => handleArchiveWorkspace(card.id)}
              >
                <Archive className="size-4" />
                <p className="text-sm">Archive</p>
              </div>
            </>
          )}
        </PopoverContent>
      </Popover>
    </>
  );
}
