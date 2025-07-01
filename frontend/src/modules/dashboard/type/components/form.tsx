import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
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
import { useWorkspaceType } from "@/hooks/use-workspace-type";
import { cn } from "@/lib/utils";
import { WorkspaceType } from "@/types/workspace";
import { yupResolver } from "@hookform/resolvers/yup";
import React, { ReactNode, useEffect } from "react";
import { useForm } from "react-hook-form";
import { toast } from "sonner";
import * as yup from "yup";

interface FormWorkspaceTypeProps {
  open: boolean;
  onOpenChange: () => void;
  type?: WorkspaceType;
}

interface FormInputWorkspaceType {
  name: string;
  color: string;
}

const schema = yup.object({
  name: yup.string().required("Name is required"),
  color: yup.string().required("Color is required"),
});

export default function FormWorkspaceType({
  open,
  onOpenChange,
  type,
}: FormWorkspaceTypeProps) {
  const isEdit = !!type;
  const defaultValues = {
    name: type?.name || "",
    color: type?.color || "",
  };

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
  const { createWorkspaceType, updateWorkspaceType } = useWorkspaceType();
  const { mutateAsync: createType, isPending: loadingCreate } =
    createWorkspaceType();
  const { mutateAsync: updateType, isPending: loadingUpdate } =
    updateWorkspaceType();

  const isPending = loadingCreate || loadingUpdate;

  useEffect(() => {
    reset(defaultValues);
  }, [open]);

  const onSubmit = async (data: FormInputWorkspaceType) => {
    try {
      if (isEdit) {
        await updateType({ ...data, id: type.id });
      } else {
        await createType(data);
      }
    } catch (error: any) {
      console.log(error);
    }

    onOpenChange();
  };

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>{isEdit ? "Update" : "New"} Workspace Type</DialogTitle>
          <DialogDescription>
            {isEdit ? "Update" : "Create"} Workspace Type Form
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
            <small className="text-red-500">{errors.name?.message}</small>
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
