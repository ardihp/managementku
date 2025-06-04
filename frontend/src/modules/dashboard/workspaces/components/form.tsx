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
import { Workspace, WorkspaceType } from "@/types/workspace";
import { Plus, X } from "lucide-react";
import { cn } from "@/lib/utils";

interface FormInputWorkspace {
  name: string;
  description: string;
  type_id: string;
  tags: string[];
}

const schema = yup.object().shape({
  name: yup.string().required("Workspace name is required"),
  description: yup.string().required("Workspace description is required"),
  type_id: yup.string().required("Workspace type is required"),
  tags: yup.array().required("Workspace tag is required"),
});

interface FormWorkspaceProps {
  selectedItem?: Workspace | null;
  resetSelectedItem?: () => void;
}

export default function FormWorkspace({
  selectedItem,
  resetSelectedItem = () => {},
}: FormWorkspaceProps) {
  const [openDialog, setOpenDialog] = useState(false);

  const { getListWorkspaceType, createWorkspace, updateWorkspace } =
    useWorkspace();
  const { data: listWorkspaceType, isLoading: loadingWorkspaceType } =
    getListWorkspaceType();
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
  } = useForm<FormInputWorkspace>({
    resolver: yupResolver(schema),
  });
  useEffect(() => {
    if (!openDialog) {
      reset({});
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

  console.log(watch());

  const handleRemoveWorkspaceTag = (tag: string) => {
    const newTags = watch("tags")?.filter((t) => t !== tag) || [];
    setValue("tags", [...newTags]);
  };

  const onSubmit = (data: FormInputWorkspace) => {
    const formData = { ...data, type_id: parseInt(data.type_id) };

    console.log(data);

    if (isEdit) {
      update({ id: selectedItem.id, data: formData });
    } else {
      create(formData);
    }

    setOpenDialog(false);
    resetSelectedItem();
    reset({});
  };

  return (
    <Dialog onOpenChange={setOpenDialog} open={openDialog}>
      <DialogTrigger className="flex items-center gap-2 bg-primary py-2 px-3 rounded-md cursor-pointer">
        <Plus className="size-4 text-white" />
        <p className="text-white text-sm">Add New</p>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>{isEdit ? "Update" : "New"} Workspace</DialogTitle>
          <DialogDescription>
            {isEdit ? "Update" : "Create"} Workspace Form
          </DialogDescription>
        </DialogHeader>

        <form id="workspace_tags" onSubmit={handleAddWorkspaceTag} />

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
            <label htmlFor="username" className="text-sm">
              Type
            </label>
            <Select
              value={watch("type_id")}
              onValueChange={(value) => setValue("type_id", value)}
              {...register("type_id")}
              disabled={loadingWorkspaceType || loadingCreate || loadingUpdate}
            >
              <SelectTrigger className="w-full !bg-transparent">
                <SelectValue placeholder="Workspace Type" />
              </SelectTrigger>
              <SelectContent>
                {listWorkspaceType?.data?.length >= 1 ? (
                  listWorkspaceType?.data?.map(
                    (type: WorkspaceType, key: number) => (
                      <SelectItem
                        key={key}
                        value={type.id.toString()}
                        className="capitalize hover:bg-zinc-50"
                      >
                        {type.name}
                      </SelectItem>
                    )
                  )
                ) : (
                  <SelectItem
                    value="-"
                    className="capitalize hover:bg-zinc-50"
                    disabled
                  >
                    There's no data
                  </SelectItem>
                )}
              </SelectContent>
            </Select>
            <small className="text-red-500">{errors.type_id?.message}</small>
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
                className="flex items-center justify-center bg-black h-full px-3 rounded-lg cursor-pointer"
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
