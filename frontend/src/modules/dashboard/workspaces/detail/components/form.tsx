import React, { useContext, useEffect, useState } from "react";
import { Textarea } from "@/components/ui/textarea";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import * as yup from "yup";
import { useWorkspace } from "@/hooks/use-workspace";
import { cn } from "@/lib/utils";
import { WorkspaceCard } from "@/types/workspace-card";
import { useWorkspaceCard } from "@/hooks/use-workspace-card";
import TextEditor from "@/components/ui/text-editor";
import { SocketContext } from "@/context/socket";

interface FormInputWorkspaceCard {
  title: string;
  description?: string;
}

const schema = yup.object().shape({
  title: yup.string().required("Workspace card title is required"),
});

interface FormWorkspaceProps {
  selectedItem?: WorkspaceCard | null;
  resetSelectedItem?: () => void;
  buttonClassName?: string;
  defaultStatus?: string;
  workspaceId: number;
  children: React.ReactNode;
}

export default function FormWorkspaceCard({
  selectedItem,
  resetSelectedItem = () => {},
  buttonClassName,
  defaultStatus,
  workspaceId,
  children,
}: FormWorkspaceProps) {
  const [openDialog, setOpenDialog] = useState(false);
  const { socket } = useContext(SocketContext);

  const { createWorkspaceCard, updateWorkspaceCard } = useWorkspaceCard();

  const { mutateAsync: create, isPending: loadingCreate } =
    createWorkspaceCard(workspaceId);
  const { mutateAsync: update, isPending: loadingUpdate } =
    updateWorkspaceCard(workspaceId);

  const isEdit = !!selectedItem;

  const {
    register,
    handleSubmit,
    formState: { errors },
    reset,
    clearErrors,
    setValue,
    watch,
  } = useForm<FormInputWorkspaceCard>({
    resolver: yupResolver(schema),
    defaultValues: {
      title: "",
      description: "",
    },
  });

  useEffect(() => {
    if (!openDialog) {
      reset({
        title: "",
        description: "",
      });
      resetSelectedItem();
      clearErrors();
    }
  }, [openDialog]);

  useEffect(() => {
    setOpenDialog(isEdit);

    if (isEdit) {
      reset({
        title: selectedItem.title || "",
        description: selectedItem.description || "",
      });
    }
  }, [selectedItem]);

  const onSubmit = async (value: FormInputWorkspaceCard) => {
    try {
      if (isEdit) {
        await update({
          card_id: selectedItem.id,
          data: { ...value, status_id: selectedItem.status.id },
        });
      } else {
        await create({
          data: {
            ...value,
            status_id: Number(defaultStatus || 0),
          },
        });
      }
    } catch (error) {
      console.error(error);
    }

    setOpenDialog(false);
    resetSelectedItem();
    reset({ title: "", description: "" });
    socket?.send(
      JSON.stringify({ type: "detail-card", workspace_id: workspaceId })
    );
  };

  return (
    <Dialog onOpenChange={setOpenDialog} open={openDialog}>
      <DialogTrigger className={buttonClassName}>{children}</DialogTrigger>
      <DialogContent className="min-w-xl">
        <DialogHeader>
          <DialogTitle>{isEdit ? "Update" : "New"} Card</DialogTitle>
          <DialogDescription>
            {isEdit ? "Update" : "Create"} Card Form
          </DialogDescription>
        </DialogHeader>

        <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col gap-2">
          <div className="flex flex-col gap-2">
            <label htmlFor="title" className="text-sm">
              Title
            </label>
            <Input
              id="title"
              type="text"
              placeholder="Input card title..."
              className="bg-transparent!"
              disabled={loadingCreate || loadingUpdate}
              {...register("title")}
            />
            <small className="text-red-500">{errors.title?.message}</small>
          </div>

          <div className="flex flex-col gap-2 w-auto">
            <label htmlFor="description" className="text-sm">
              Description
            </label>
            <TextEditor
              value={watch("description") || ""}
              placeholder="Input card decription..."
              onValueChange={(val) => setValue("description", val)}
              disabled={loadingCreate || loadingUpdate}
            />
            <small className="text-red-500">
              {errors.description?.message}
            </small>
          </div>

          <button
            className={cn(
              "border border-zinc-200 h-10 bg-zinc-700 rounded-lg shadow w-full cursor-pointer duration-200",
              (loadingCreate || loadingUpdate) &&
                "cursor-not-allowed opacity-60"
            )}
            disabled={loadingCreate || loadingUpdate}
          >
            <p className="text-white text-sm">{isEdit ? "Save" : "Create"}</p>
          </button>
        </form>
      </DialogContent>
    </Dialog>
  );
}
