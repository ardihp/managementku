"use server";

import { endpoint } from "@/const/endpoint";
import { paths } from "@/const/routes";
import { AuthContext } from "@/context/auth";
import Api from "@/services/api";
import { useMutation } from "@tanstack/react-query";
import Cookies from "js-cookie";
import { cookies } from "next/headers";
import { useRouter } from "next/router";
import { useContext } from "react";
import { toast } from "sonner";

interface LoginPayload {
  username: string;
  password: string;
}

interface RegisterPayload {
  name: string;
  username: string;
  email: string;
  role: string;
  password: string;
}

export interface UserDetail {
  id: number;
  name: string;
  username: string;
  email: string;
  role: string;
}

export const useAuth = () => {
  const user = Cookies.get("user");
  const { setIsAuth } = useContext(AuthContext);
  const router = useRouter();

  const HandleSignIn = () => {
    return useMutation({
      mutationFn: async (payload: LoginPayload) => {
        const response = await Api.post(endpoint.auth.sign_in, payload);

        return response.data;
      },
    });
  };

  const HandleSignUp = () => {
    return useMutation({
      mutationFn: async (payload: RegisterPayload) => {
        const response = await Api.post(endpoint.auth.sign_up, payload);

        return response.data;
      },
    });
  };

  const HandleLogout = () => {
    const toastID = "sign-out";
    toast("Signing out...", { id: toastID });
    Cookies.remove("token");
    Cookies.remove("user");
    setIsAuth(false);
    router.push(paths.auth.sign_in);
    toast("Sign out success.", { id: toastID });
  };

  return {
    user: user ? (JSON.parse(user!) as UserDetail) : null,
    userSignIn: HandleSignIn,
    userSignUp: HandleSignUp,
    userSignOut: HandleLogout,
  };
};
