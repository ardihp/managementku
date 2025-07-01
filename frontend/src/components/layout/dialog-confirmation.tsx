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
  title: string;
  description: string;
  open: boolean;
  onOpenChange: () => void;
  onSubmit: (e: React.FormEvent<HTMLFormElement>) => void;
  isPending: boolean;
}

export default function DialogConfirmation({
  title,
  description,
  open,
  onOpenChange,
  onSubmit,
  isPending,
}: DialogConfirmationProps) {
  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>{title}</DialogTitle>
          <DialogDescription>{description}</DialogDescription>
        </DialogHeader>

        <form onSubmit={onSubmit}>
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
