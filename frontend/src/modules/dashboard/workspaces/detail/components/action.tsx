import React, { useState } from "react";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import {
  Archive,
  ArchiveRestore,
  BadgeMinus,
  CircleCheck,
  CircleX,
  Ellipsis,
  Merge,
  Pencil,
  Trash2,
  Users2,
} from "lucide-react";
import { useWorkspaceCard } from "@/hooks/use-workspace-card";
import { WorkspaceCard } from "@/types/workspace-card";
import AssignMember from "./assign-member";
import { useAuth } from "@/hooks/auth/use-auth";

interface WorkspaceActionProps {
  card: WorkspaceCard;
  isArchived: boolean;
  workspaceId?: number;
  isAdmin?: boolean;
  afterAction?: () => void;
  setSelectedCard?: (item: any) => void;
}

export default function WorkspaceAction({
  card,
  isArchived = false,
  workspaceId = 0,
  isAdmin = false,
  afterAction = () => {},
  setSelectedCard = () => {},
}: WorkspaceActionProps) {
  const [openPopover, setOpenPopover] = useState(false);
  const [openDialog, setOpenDialog] = useState({
    member: false,
  });
  const { user } = useAuth();

  const {
    archivedWorkspaceCard,
    unarchivedWorkspaceCard,
    deleteWorkspaceCard,
    watchWorkspaceCard,
    unwatchWorkspaceCard,
    completeWorkspaceCard,
    incompleteWorkspaceCard,
  } = useWorkspaceCard();

  const { mutateAsync: archiveCard } = archivedWorkspaceCard(workspaceId);
  const { mutateAsync: unarchiveCard } = unarchivedWorkspaceCard(workspaceId);
  const { mutateAsync: deleteCard } = deleteWorkspaceCard(workspaceId);
  const { mutateAsync: watch } = watchWorkspaceCard(workspaceId, card.id);
  const { mutateAsync: unwatch } = unwatchWorkspaceCard(workspaceId, card.id);
  const { mutateAsync: complete } = completeWorkspaceCard(workspaceId);
  const { mutateAsync: incomplete } = incompleteWorkspaceCard(workspaceId);

  const handleArchiveWorkspace = async (id: number) => {
    try {
      await archiveCard(id);
    } catch (error) {
      console.error(error);
    }
    setOpenPopover(false);
    afterAction();
  };

  const handleUnarchiveWorkspace = async (id: number) => {
    try {
      await unarchiveCard(id);
    } catch (error) {
      console.error(error);
    }
    setOpenPopover(false);
    afterAction();
  };

  const handleDeleteWorkspace = async (id: number) => {
    try {
      await deleteCard(id);
    } catch (error) {
      console.error(error);
    }
    setOpenPopover(false);
    afterAction();
  };

  const handleWatchWorkspace = async () => {
    try {
      await watch({ user_id: user?.id || 0 });
    } catch (error) {
      console.error(error);
    }
    setOpenPopover(false);
  };

  const handleUnwatchWorkspace = async () => {
    try {
      await unwatch({ user_id: user?.id || 0 });
    } catch (error) {
      console.error(error);
    }
    setOpenPopover(false);
  };

  const handleCompletedWorkspace = async () => {
    try {
      await complete(card.id);
    } catch (error) {
      console.error(error);
    }
    setOpenPopover(false);
  };

  const handleIncompletedWorkspace = async () => {
    try {
      await incomplete(card.id);
    } catch (error) {
      console.error(error);
    }
    setOpenPopover(false);
  };

  return (
    <>
      {openDialog.member && (
        <AssignMember
          open={openDialog.member}
          onOpenChange={() =>
            setOpenDialog((prevVal) => ({ ...prevVal, member: false }))
          }
          workspaceId={workspaceId}
          card={card}
        />
      )}

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
              {card.is_completed ? (
                <div
                  className="flex items-center gap-2 hover:bg-zinc-50 py-2 px-3 rounded cursor-pointer"
                  onClick={() => handleIncompletedWorkspace()}
                >
                  <CircleX className="size-4" />
                  <p className="text-sm">Mark Incomplete</p>
                </div>
              ) : (
                <div
                  className="flex items-center gap-2 hover:bg-zinc-50 py-2 px-3 rounded cursor-pointer"
                  onClick={() => handleCompletedWorkspace()}
                >
                  <CircleCheck className="size-4" />
                  <p className="text-sm">Mark Complete</p>
                </div>
              )}
              <div
                className="flex items-center gap-2 hover:bg-zinc-50 py-2 px-3 rounded cursor-pointer"
                onClick={() => setSelectedCard(card)}
              >
                <Pencil className="size-4" />
                <p className="text-sm">Update</p>
              </div>
              {isAdmin ? (
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
              ) : card.members.some((member) => member.id === user?.id) ? (
                <div
                  className="flex items-center gap-2 hover:bg-zinc-50 py-2 px-3 rounded cursor-pointer"
                  onClick={() => {
                    setOpenPopover(false);
                    handleUnwatchWorkspace();
                  }}
                >
                  <BadgeMinus className="size-4" />
                  <p className="text-sm">Unwatch this Card</p>
                </div>
              ) : (
                <div
                  className="flex items-center gap-2 hover:bg-zinc-50 py-2 px-3 rounded cursor-pointer"
                  onClick={() => {
                    setOpenPopover(false);
                    handleWatchWorkspace();
                  }}
                >
                  <Merge className="size-4 -rotate-90" />
                  <p className="text-sm">Watch this Card</p>
                </div>
              )}
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
