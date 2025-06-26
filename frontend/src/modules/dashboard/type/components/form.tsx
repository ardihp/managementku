import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { WorkspaceType } from "@/types/workspace";
import React, { ReactNode } from "react";

interface FormWorkspaceTypeProps {
  type?: WorkspaceType;
  children: ReactNode;
}

export default function FormWorkspaceType({
  type,
  children,
}: FormWorkspaceTypeProps) {
  const isEdit = !!type;

  return (
    <Dialog>
      <DialogTrigger asChild>{children}</DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>
            {isEdit ? "Update" : "Create"} Workspace Type
          </DialogTitle>
          <DialogDescription></DialogDescription>
        </DialogHeader>
      </DialogContent>
    </Dialog>
  );
}
