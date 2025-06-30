"use client";

import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { ColumnDef } from "@tanstack/react-table";
import { MoreVertical } from "lucide-react";
import { useState } from "react";
import DialogConfirmation from "./dialog-confirmation";
import { colorsType } from "@/const";
import FormWorkspaceType from "./form";

export interface WorkspaceType {
  id: number;
  name: string;
  color: colorsType;
}

export const columns: ColumnDef<WorkspaceType>[] = [
  {
    accessorKey: "name",
    header: "Name",
  },
  {
    accessorKey: "color",
    header: "Color",
  },
  {
    id: "actions",
    cell: ({ row }) => {
      const [openDialog, setOpenDialog] = useState({
        confirmation: false,
        edit: false,
      });

      return (
        <>
          <DialogConfirmation
            open={openDialog.confirmation}
            onOpenChange={() =>
              setOpenDialog((prevVal) => ({ ...prevVal, confirmation: false }))
            }
            type={row.original}
          />

          <FormWorkspaceType
            open={openDialog.edit}
            onOpenChange={() =>
              setOpenDialog((prevVal) => ({ ...prevVal, edit: false }))
            }
            type={row.original}
          />

          <div className="flex w-full items-center justify-end">
            <DropdownMenu>
              <DropdownMenuTrigger asChild>
                <Button variant="ghost" className="h-8 w-8 p-0">
                  <span className="sr-only">Open menu</span>
                  <MoreVertical className="h-4 w-4" />
                </Button>
              </DropdownMenuTrigger>
              <DropdownMenuContent className="min-w-40" align="end">
                <DropdownMenuItem
                  onClick={() =>
                    setOpenDialog((prevVal) => ({ ...prevVal, edit: true }))
                  }
                >
                  Edit
                </DropdownMenuItem>
                <DropdownMenuSeparator />
                <DropdownMenuItem
                  className="text-red-500 hover:text-red-500!"
                  onClick={() =>
                    setOpenDialog((prevVal) => ({
                      ...prevVal,
                      confirmation: true,
                    }))
                  }
                >
                  Delete
                </DropdownMenuItem>
              </DropdownMenuContent>
            </DropdownMenu>
          </div>
        </>
      );
    },
  },
];
