import React, { useEffect, useState } from "react";
import { Textarea } from "@/components/ui/textarea";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
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
import { Workspace, WorkspaceStatus, WorkspaceType } from "@/types/workspace";
import { Plus, X } from "lucide-react";
import { cn, handleConffeti } from "@/lib/utils";
import { WorkspaceCard } from "@/types/workspace-card";
import { useWorkspaceCard } from "@/hooks/use-workspace-card";

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
  defaultStatus: string;
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

  const { getListWorkspaceStatus, updateWorkspace } = useWorkspace();
  const { createWorkspaceCard } = useWorkspaceCard();
  const { data: listWorkspaceStatus, isLoading: loadingWorkspaceStatus } =
    getListWorkspaceStatus("false");

  const { mutate: create, isPending: loadingCreate } = createWorkspaceCard();
  const { mutate: update, isPending: loadingUpdate } = updateWorkspace();
  const isEdit = !!selectedItem;

  const {
    register,
    handleSubmit,
    formState: { errors },
    setValue,
    watch,
    reset,
    clearErrors,
  } = useForm<FormInputWorkspaceCard>({
    resolver: yupResolver(schema),
  });

  useEffect(() => {
    if (!openDialog) {
      reset({
        title: "",
        description: "",
      });
      resetSelectedItem();
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

  const onSubmit = (value: FormInputWorkspaceCard) => {
    if (isEdit) {
    } else {
      create({
        id: workspaceId,
        data: {
          ...value,
          status_id: Number(defaultStatus || 0),
        },
      });
    }

    setOpenDialog(false);
    resetSelectedItem();
    reset({ title: "", description: "" });
  };

  return (
    <Dialog onOpenChange={setOpenDialog} open={openDialog}>
      <DialogTrigger className={buttonClassName}>{children}</DialogTrigger>
      <DialogContent>
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

          <div className="flex flex-col gap-2">
            <label htmlFor="description" className="text-sm">
              Description
            </label>
            <Textarea
              id="description"
              rows={4}
              placeholder="Input card description..."
              className="bg-transparent! resize-none h-[80px]"
              disabled={loadingCreate || loadingUpdate}
              {...register("description")}
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
