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
import { useAuth } from "@/hooks/auth/use-auth";
import { useUser } from "@/hooks/use-user";
import { useWorkspaceType } from "@/hooks/use-workspace-type";
import { cn } from "@/lib/utils";
import { Member, WorkspaceType } from "@/types/workspace";
import { yupResolver } from "@hookform/resolvers/yup";
import { Eye, EyeClosed } from "lucide-react";
import React, { ReactNode, useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import { toast } from "sonner";
import * as yup from "yup";

interface FormUserProps {
  open: boolean;
  onOpenChange: () => void;
  afterUpdate?: (val: any) => void;
  user?: Member;
}

interface FormUser {
  name: string;
  username: string;
  email: string;
  role: string;
  password?: string;
}

const schema = yup.object().shape({
  name: yup.string().required("Name is required"),
  username: yup.string().required("Username is required"),
  email: yup.string().email().required("Email is required"),
  role: yup.string().required("Role is required"),
  password: yup.string().optional(),
});

export default function FormUser({
  open,
  onOpenChange,
  user,
  afterUpdate = () => {},
}: FormUserProps) {
  const isEdit = !!user;
  const defaultValues = {
    name: user?.name || "",
    username: user?.username || "",
    email: user?.email || "",
    role: user?.role || "",
    password: "",
  };
  const [showPassword, setShowPassword] = useState(false);

  const {
    register,
    handleSubmit,
    watch,
    setValue,
    clearErrors,
    reset,
    formState: { errors },
  } = useForm({
    resolver: yupResolver(schema),
    defaultValues,
  });
  const { updateUser, createUser } = useUser();
  const { mutateAsync: create, isPending: loadingCreate } = createUser();
  const { mutateAsync: update, isPending: loadingUpdate } = updateUser();

  const isPending = loadingCreate || loadingUpdate;

  useEffect(() => {
    reset(defaultValues);
    setShowPassword(false);
  }, [open]);

  const onSubmit = async (data: FormUser) => {
    try {
      if (isEdit) {
        await update(
          { user_id: user?.id, data },
          {
            onSuccess: (response) => {
              afterUpdate(response?.data);
            },
          }
        );
      } else {
        await create({ ...data, password: data.password! });
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
          <DialogTitle>{isEdit ? "Update" : "New"} User</DialogTitle>
          <DialogDescription>
            {isEdit ? "Update" : "Create"} User Form
          </DialogDescription>
        </DialogHeader>

        <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col gap-2">
          <div className="flex flex-col gap-2">
            <label htmlFor="name" className="text-sm">
              Fullname
            </label>
            <Input
              type="text"
              placeholder="Input fullname..."
              className="bg-transparent!"
              disabled={isPending}
              {...register("name")}
            />
            <small className="text-red-500">{errors.name?.message}</small>
          </div>

          <div className="flex flex-col gap-2">
            <label htmlFor="username" className="text-sm">
              Username
            </label>
            <Input
              type="text"
              placeholder="Input username..."
              className="bg-transparent!"
              disabled={isPending}
              {...register("username")}
            />
            <small className="text-red-500">{errors.username?.message}</small>
          </div>

          <div className="flex flex-col gap-2">
            <label htmlFor="email" className="text-sm">
              Email
            </label>
            <Input
              type="text"
              placeholder="Input email..."
              className="bg-transparent!"
              disabled={isPending}
              {...register("email")}
            />
            <small className="text-red-500">{errors.email?.message}</small>
          </div>

          <div className="flex flex-col gap-2">
            <label htmlFor="role" className="text-sm">
              Role
            </label>
            <Select
              value={watch("role")}
              onValueChange={(value) => {
                setValue("role", value);
                clearErrors("role");
              }}
              {...register("role")}
              disabled={isPending}
            >
              <SelectTrigger className="w-full !bg-transparent capitalize">
                <SelectValue placeholder="User role" className="capitalize" />
              </SelectTrigger>
              <SelectContent>
                {["manager", "employee"]?.map((role, key) => (
                  <SelectItem
                    key={key}
                    value={role}
                    className="capitalize hover:bg-zinc-50"
                  >
                    {role}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
            <small className="text-red-500">{errors.name?.message}</small>
          </div>

          <div className="flex flex-col gap-2">
            <label htmlFor="password" className="text-sm">
              {isEdit ? "New Password" : "Password"}
            </label>
            <Input
              type={showPassword ? "text" : "password"}
              placeholder={
                isEdit
                  ? "Just leave it blank if you don't want to change the password"
                  : "Input password..."
              }
              className="bg-transparent!"
              rightIcon={
                <div
                  className="flex items-center justify-center bg-white w-10 h-full right-0 cursor-pointer"
                  onClick={() => setShowPassword(!showPassword)}
                >
                  {showPassword ? (
                    <Eye className="size-4" />
                  ) : (
                    <EyeClosed className="size-4" />
                  )}
                </div>
              }
              disabled={isPending}
              {...register("password")}
            />
            <small className="text-red-500">{errors.password?.message}</small>
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
