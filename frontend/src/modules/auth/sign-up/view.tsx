import AuthLayout from "@/components/layout/auth";
import { Input } from "@/components/ui/input";
import { paths } from "@/const/routes";
import { useAuth } from "@/hooks/auth/use-auth";
import { cn } from "@/lib/utils";
import { yupResolver } from "@hookform/resolvers/yup";
import { Eye, EyeClosed } from "lucide-react";
import { useRouter } from "next/router";
import React, { useState } from "react";
import { useForm } from "react-hook-form";
import { toast } from "sonner";
import * as yup from "yup";
import Cookies from "js-cookie";

interface FormInputValue {
  name: string;
  username: string;
  email: string;
  password: string;
}

const schema = yup.object().shape({
  name: yup.string().required(""),
  username: yup.string().required(""),
  email: yup.string().email().required(""),
  password: yup.string().required(""),
});

export default function SignInView() {
  const [showPassword, setShowPassword] = useState(false);
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormInputValue>({
    resolver: yupResolver(schema),
  });
  const router = useRouter();
  const { userSignUp } = useAuth();
  const { mutate, isPending } = userSignUp();

  const onSubmit = (data: FormInputValue) => {
    const formData = {
      ...data,
      role: "employee",
    };

    mutate(formData, {
      onSuccess: (response) => {
        toast(response?.message);
        Cookies.set("token", response?.data?.token);
        Cookies.set("user", JSON.stringify(response?.data));
        router.push(paths.dashboard.root);
      },
    });
  };

  return (
    <AuthLayout>
      <div className="relative -top-10 xl:-top-0 z-[2] flex flex-col w-full max-w-[420px] gap-10 xl:justify-center px-6">
        <p className="text-2xl font-semibold">Sign Up</p>
        <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col gap-4">
          <div className="flex flex-col gap-1">
            <label htmlFor="name" className="text-sm">
              Name
            </label>
            <Input
              type="text"
              placeholder="Input your name..."
              className="bg-transparent!"
              disabled={isPending}
              {...register("name")}
            />
            <small className="text-red-500">{errors.name?.message}</small>
          </div>

          <div className="flex flex-col gap-1">
            <label htmlFor="email" className="text-sm">
              Email
            </label>
            <Input
              type="email"
              placeholder="Input your email..."
              className="bg-transparent!"
              disabled={isPending}
              {...register("email")}
            />
            <small className="text-red-500">{errors.email?.message}</small>
          </div>

          <div className="flex flex-col gap-1">
            <label htmlFor="username" className="text-sm">
              Username
            </label>
            <Input
              type="text"
              placeholder="Input your username..."
              className="bg-transparent!"
              disabled={isPending}
              {...register("username")}
            />
            <small className="text-red-500">{errors.username?.message}</small>
          </div>

          <div className="flex flex-col gap-1">
            <label htmlFor="password" className="text-sm">
              Password
            </label>

            <Input
              type={showPassword ? "text" : "password"}
              placeholder="Input you password..."
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

          <div className="flex flex-col gap-5 items-center">
            <button
              className={cn(
                "border border-zinc-200 h-10 bg-zinc-700 rounded-lg shadow w-full cursor-pointer duration-200",
                isPending && "cursor-not-allowed opacity-60"
              )}
              disabled={isPending}
            >
              <p className="text-white text-sm leading-none">
                {isPending ? "Loading..." : "Sign Up"}
              </p>
            </button>

            <p className="text-sm text-zinc-500">
              {`Already have an account? `}
              <span
                className="text-sm text-zinc-700 cursor-pointer font-semibold"
                onClick={() => router.push(paths.auth.sign_in)}
              >
                Sign In
              </span>
            </p>
          </div>
        </form>
      </div>
    </AuthLayout>
  );
}
