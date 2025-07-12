import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { colors } from "@/const";
import { SocketContext } from "@/context/socket";
import { useWorkspaceCardStatus } from "@/hooks/use-workspace-card-status";
import { useWorkspaceStatus } from "@/hooks/use-workspace-status";
import { cn } from "@/lib/utils";
import { WorkspaceStatus } from "@/types/workspace";
import { yupResolver } from "@hookform/resolvers/yup";
import React, { useContext, useEffect } from "react";
import { useForm } from "react-hook-form";
import * as yup from "yup";

interface FormWorkspaceStatusProps {
  open: boolean;
  onOpenChange: () => void;
  status?: WorkspaceStatus | null;
  type: "workspace" | "workspace-detail";
  workspaceId?: number;
}

interface FormInputWorkspaceType {
  name: string;
  color: string;
}

const schema = yup.object({
  name: yup.string().required("Name is required"),
  color: yup.string().required("Color is required"),
});

export default function FormWorkspaceStatus({
  open,
  onOpenChange,
  status,
  type,
  workspaceId = 0,
}: FormWorkspaceStatusProps) {
  const isEdit = !!status;
  const defaultValues = {
    name: status?.name || "",
    color: status?.color || "",
  };
  const { socket } = useContext(SocketContext);

  const {
    register,
    handleSubmit,
    watch,
    setValue,
    clearErrors,
    reset,
    formState: { errors },
  } = useForm<FormInputWorkspaceType>({
    resolver: yupResolver(schema),
    defaultValues,
  });

  // type workspace
  const { createWorkspaceStatus, updateWorkspaceStatus } = useWorkspaceStatus();

  const { mutateAsync: createStatus, isPending: loadingCreateStatus } =
    createWorkspaceStatus();
  const { mutateAsync: updateStatus, isPending: loadingUpdateStatus } =
    updateWorkspaceStatus();

  // type workspace-detail
  const { createWorkspaceCardStatus, updateWorkspaceCardStatus } =
    useWorkspaceCardStatus();

  const { mutateAsync: createCardStatus, isPending: loadingCreateCardStatus } =
    createWorkspaceCardStatus(workspaceId);
  const { mutateAsync: updateCardStatus, isPending: loadingUpdateCardStatus } =
    updateWorkspaceCardStatus(workspaceId);

  useEffect(() => {
    reset(defaultValues);
  }, [open]);

  const isPending =
    loadingCreateStatus ||
    loadingUpdateStatus ||
    loadingCreateCardStatus ||
    loadingUpdateCardStatus;

  const onSubmit = async (data: FormInputWorkspaceType) => {
    try {
      if (isEdit) {
        if (type === "workspace") {
          await updateStatus({ ...data, id: status?.id });
        } else {
          await updateCardStatus({ data: data, status_id: status?.id });
        }
      } else {
        if (type === "workspace") {
          await createStatus(data);
        } else {
          await createCardStatus(data);
        }
      }
    } catch (error) {
      console.error(error);
    }

    onOpenChange();
    socket?.send(
      JSON.stringify({ type: "mark-progress", workspace_id: workspaceId })
    );
  };

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>
            {isEdit ? "Update" : "New"} Workspace Status
          </DialogTitle>
          <DialogDescription>
            {isEdit ? "Update" : "Create"} Workspace Status Form
          </DialogDescription>
        </DialogHeader>

        <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col gap-2">
          <div className="flex flex-col gap-2">
            <label htmlFor="username" className="text-sm">
              Name
            </label>
            <Input
              type="text"
              placeholder="Input workspace name..."
              className="bg-transparent!"
              disabled={isPending}
              {...register("name")}
            />
            <small className="text-red-500">{errors.name?.message}</small>
          </div>

          <div className="flex flex-col gap-2">
            <label htmlFor="username" className="text-sm">
              Color
            </label>
            <Select
              value={watch("color")}
              onValueChange={(value) => {
                setValue("color", value);
                clearErrors("color");
              }}
              {...register("color")}
              disabled={isPending}
            >
              <SelectTrigger
                id="type"
                className="w-full !bg-transparent capitalize"
              >
                <SelectValue
                  placeholder="Workspace Type"
                  className="capitalize"
                />
              </SelectTrigger>
              <SelectContent>
                {colors?.map((color, key) => (
                  <SelectItem
                    key={key}
                    value={color}
                    className="capitalize hover:bg-zinc-50"
                  >
                    {color}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
            <small className="text-red-500">{errors.color?.message}</small>
          </div>

          <button
            className={cn(
              "border border-zinc-200 h-10 bg-zinc-700 rounded-lg shadow w-full cursor-pointer duration-200",
              isPending && "cursor-not-allowed opacity-60"
            )}
            disabled={isPending}
          >
            <p className="text-white text-sm">{isEdit ? "Save" : "Create"}</p>
          </button>
        </form>
      </DialogContent>
    </Dialog>
  );
}
