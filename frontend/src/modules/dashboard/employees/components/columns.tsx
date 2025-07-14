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
import DialogConfirmation from "../../../../components/layout/dialog-confirmation";
import { useWorkspaceType } from "@/hooks/use-workspace-type";
import { Member } from "@/types/workspace";
import { useAuth } from "@/hooks/auth/use-auth";
import FormUser from "./form";
import { useUser } from "@/hooks/use-user";

export const columns: ColumnDef<Member>[] = [
  {
    accessorKey: "name",
    header: "Fullname",
  },
  {
    accessorKey: "username",
    header: "Username",
  },
  {
    accessorKey: "email",
    header: "Email",
  },
  {
    accessorKey: "role",
    header: "Role",
    cell: ({ row }) => {
      return <p className="capitalize">{row.original.role}</p>;
    },
  },
  {
    id: "actions",
    cell: ({ row }) => {
      const [openDialog, setOpenDialog] = useState({
        confirmation: false,
        edit: false,
      });
      const { user } = useAuth();

      const { deleteUser } = useUser();
      const { mutateAsync: userDelete, isPending } = deleteUser();

      const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        try {
          await userDelete(row.original.id);
        } catch (error) {
          console.error(error);
        }

        setOpenDialog((prevVal) => ({ ...prevVal, confirmation: false }));
      };

      return (
        <>
          <DialogConfirmation
            title="Confirm delete type"
            description="Are you sure you want to permanently delete this type? This action
            cannot be undone."
            open={openDialog.confirmation}
            onOpenChange={() =>
              setOpenDialog((prevVal) => ({ ...prevVal, confirmation: false }))
            }
            onSubmit={handleSubmit}
            isPending={isPending}
          />

          <FormUser
            open={openDialog.edit}
            onOpenChange={() =>
              setOpenDialog((prevVal) => ({ ...prevVal, edit: false }))
            }
            user={row.original}
          />

          <div className="flex w-full items-center justify-end">
            {row.original.username !== user?.username &&
              user?.role === "manager" && (
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
              )}
          </div>
        </>
      );
    },
  },
];
