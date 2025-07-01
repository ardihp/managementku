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
import { useWorkspaceStatus } from "@/hooks/use-workspace-status";

interface FormInputWorkspace {
  name: string;
  description: string;
  type_id: string;
  status_id: string;
  tags?: string[];
}

const schema = yup.object().shape({
  name: yup.string().required("Workspace name is required"),
  description: yup.string().required("Workspace description is required"),
  type_id: yup.string().required("Workspace type is required"),
  status_id: yup.string().required("Workspace status is required"),
});

interface FormWorkspaceProps {
  selectedItem?: Workspace | null;
  resetSelectedItem?: () => void;
  buttonClassName?: string;
  defaultStatus?: string;
  children: React.ReactNode;
}

export default function FormWorkspace({
  selectedItem,
  resetSelectedItem = () => {},
  buttonClassName,
  defaultStatus,
  children,
}: FormWorkspaceProps) {
  const [openDialog, setOpenDialog] = useState(false);

  const { getListWorkspaceType, createWorkspace, updateWorkspace } =
    useWorkspace();
  const { getListWorkspaceStatus } = useWorkspaceStatus();

  const { data: listWorkspaceType, isLoading: loadingWorkspaceType } =
    getListWorkspaceType();
  const { data: listWorkspaceStatus, isLoading: loadingWorkspaceStatus } =
    getListWorkspaceStatus();

  const { mutate: create, isPending: loadingCreate } = createWorkspace();
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
  } = useForm<FormInputWorkspace>({
    resolver: yupResolver(schema),
  });

  useEffect(() => {
    if (!openDialog) {
      reset({
        name: "",
        description: "",
        type_id: "",
        status_id: defaultStatus || "",
        tags: [],
      });
      resetSelectedItem();
    }
  }, [openDialog]);

  useEffect(() => {
    setOpenDialog(isEdit);

    if (isEdit) {
      reset({
        name: selectedItem.name || "",
        description: selectedItem.description || "",
        type_id: selectedItem.type.id.toString() || "",
        status_id: selectedItem.status.id.toString() || "",
        tags: selectedItem.tags || [],
      });
    }
  }, [selectedItem]);

  const handleAddWorkspaceTag = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    // @ts-ignore: Tag is available in form component
    if (e.target.tag.value !== "") {
      setValue(
        "tags",
        // @ts-ignore: Tag is available in form component
        new Set([...(watch("tags") || []), e.target.tag.value])
          .values()
          .toArray()
      );
      // @ts-ignore: Tag is available in form component
      e.target.tag.value = "";
    }
  };

  const handleRemoveWorkspaceTag = (tag: string) => {
    const newTags = watch("tags")?.filter((t) => t !== tag) || [];
    setValue("tags", [...newTags]);
  };

  const onSubmit = (data: FormInputWorkspace) => {
    const formData = {
      ...data,
      type_id: parseInt(data.type_id),
      status_id: parseInt(data.status_id),
    };

    if (isEdit) {
      update(
        { id: selectedItem.id, data: formData },
        {
          onSuccess(response) {
            if (response.data.status.name === "Completed") handleConffeti();
          },
        }
      );
    } else {
      create(formData);
    }

    setOpenDialog(false);
    resetSelectedItem();
    reset({ name: "", description: "", type_id: "", status_id: "", tags: [] });
  };

  return (
    <Dialog onOpenChange={setOpenDialog} open={openDialog}>
      <DialogTrigger className={buttonClassName}>{children}</DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>{isEdit ? "Update" : "New"} Workspace</DialogTitle>
          <DialogDescription>
            {isEdit ? "Update" : "Create"} Workspace Form
          </DialogDescription>
        </DialogHeader>

        <form id="workspace_tags" onSubmit={handleAddWorkspaceTag} hidden />

        <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col gap-2">
          <div className="flex flex-col gap-2">
            <label htmlFor="username" className="text-sm">
              Name
            </label>
            <Input
              type="text"
              placeholder="Input workspace name..."
              className="bg-transparent!"
              disabled={loadingCreate || loadingUpdate}
              {...register("name")}
            />
            <small className="text-red-500">{errors.name?.message}</small>
          </div>

          <div className="flex flex-col gap-2">
            <label className="text-sm">Type</label>
            <Select
              value={watch("type_id")}
              onValueChange={(value) => {
                setValue("type_id", value);
                clearErrors("type_id");
              }}
              {...register("type_id")}
              disabled={loadingWorkspaceType || loadingCreate || loadingUpdate}
            >
              <SelectTrigger id="type" className="w-full !bg-transparent">
                <SelectValue placeholder="Workspace Type" />
              </SelectTrigger>
              <SelectContent>
                {loadingWorkspaceType || listWorkspaceType?.length < 1 ? (
                  <SelectItem
                    value="-"
                    className="capitalize hover:bg-zinc-50"
                    disabled
                  >
                    There's no data
                  </SelectItem>
                ) : (
                  listWorkspaceType?.map((type: WorkspaceType, key: number) => (
                    <SelectItem
                      key={key}
                      value={type.id.toString()}
                      className="capitalize hover:bg-zinc-50"
                    >
                      {type.name}
                    </SelectItem>
                  ))
                )}
              </SelectContent>
            </Select>
            <small className="text-red-500">{errors.type_id?.message}</small>
          </div>

          <div className="flex flex-col gap-2">
            <label className="text-sm">Status</label>
            <Select
              value={watch("status_id")}
              onValueChange={(value) => {
                setValue("status_id", value);
                clearErrors("status_id");
              }}
              disabled={
                loadingWorkspaceStatus ||
                loadingCreate ||
                loadingUpdate ||
                !!defaultStatus
              }
            >
              <SelectTrigger id="status" className="w-full !bg-transparent">
                <SelectValue placeholder="Workspace Status" />
              </SelectTrigger>
              <SelectContent>
                {loadingWorkspaceStatus || listWorkspaceStatus?.length < 1 ? (
                  <SelectItem
                    value="-"
                    className="capitalize hover:bg-zinc-50"
                    disabled
                  >
                    There's no data
                  </SelectItem>
                ) : (
                  listWorkspaceStatus?.map(
                    (status: WorkspaceStatus, key: number) => (
                      <SelectItem
                        key={key}
                        value={status.id.toString()}
                        className="capitalize hover:bg-zinc-50"
                      >
                        {status.name}
                      </SelectItem>
                    )
                  )
                )}
              </SelectContent>
            </Select>
            <small className="text-red-500">{errors.status_id?.message}</small>
          </div>

          <div className="flex flex-col gap-2">
            <label htmlFor="username" className="text-sm">
              Tags
            </label>
            <div className="flex items-center gap-2">
              <Input
                name="tag"
                type="text"
                placeholder="Input workspace tags..."
                className="bg-transparent!"
                wrapperClassname="w-full"
                form="workspace_tags"
              />

              <button
                className="flex items-center justify-center bg-zinc-700 h-full px-3 rounded-lg cursor-pointer"
                form="workspace_tags"
              >
                <p className="text-white text-sm">Add</p>
              </button>
            </div>
            {(watch("tags") || []).length >= 1 && (
              <div className="flex flex-wrap gap-2 item-center">
                {watch("tags")?.map((tag, key) => (
                  <div
                    key={key}
                    className="flex gap-2 items-center justify-center py-1 px-2 bg-zinc-100 rounded-lg"
                  >
                    <p className="text-sm">{tag}</p>

                    <X
                      className="size-4 cursor-pointer"
                      onClick={() => handleRemoveWorkspaceTag(tag)}
                    />
                  </div>
                ))}
              </div>
            )}
            <small className="text-red-500">{errors.tags?.message}</small>
          </div>

          <div className="flex flex-col gap-2">
            <label htmlFor="username" className="text-sm">
              Description
            </label>
            <Textarea
              rows={4}
              placeholder="Input workspace description..."
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
