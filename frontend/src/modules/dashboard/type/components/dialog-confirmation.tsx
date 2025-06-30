import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { useWorkspaceType } from "@/hooks/use-workspace-type";
import { cn } from "@/lib/utils";
import { WorkspaceType } from "@/types/workspace";
import React from "react";

interface DialogConfirmationProps {
  open: boolean;
  onOpenChange: () => void;
  type: WorkspaceType;
}

export default function DialogConfirmation({
  open,
  onOpenChange,
  type,
}: DialogConfirmationProps) {
  const { deleteWorkspaceType } = useWorkspaceType();
  const { mutateAsync: deleteType, isPending } = deleteWorkspaceType();

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    try {
      await deleteType(type.id);
    } catch (error) {
      console.error(error);
    }

    onOpenChange();
  };

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>
            Confirm delete type <strong>"{type.name}"</strong>?
          </DialogTitle>
          <DialogDescription>
            Are you sure you want to permanently delete this user? This action
            cannot be undone.
          </DialogDescription>
        </DialogHeader>

        <form onSubmit={handleSubmit}>
          <DialogFooter>
            <DialogClose asChild>
              <Button
                type="button"
                variant="outline"
                className={cn(isPending && "opacity-70")}
                disabled={isPending}
              >
                Cancel
              </Button>
            </DialogClose>
            <Button
              type="submit"
              className={cn(
                "bg-red-500 hover:bg-red-900 w-[90px]",
                isPending && "opacity-70"
              )}
              disabled={isPending}
            >
              {isPending ? "Loading..." : "Confirm"}
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
}
