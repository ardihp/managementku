import { endpoint } from "@/const/endpoint";
import Api from "@/services/api";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { useAuth } from "./auth/use-auth";
import { UserCreateRequest, UserUpdateRequest } from "@/types/user";
import { toast } from "sonner";

export const useUser = () => {
  const { user } = useAuth();
  const queryClient = useQueryClient();

  const GetListUser = () => {
    return useQuery({
      queryKey: ["user/list"],
      queryFn: async () => {
        const response = await Api.get(endpoint.user.list_user);
        return response.data.data;
      },
      staleTime: Infinity,
    });
  };

  const CreateUser = () => {
    return useMutation({
      mutationKey: ["user/create"],
      mutationFn: async (payload: UserCreateRequest) => {
        const response = await Api.post(endpoint.auth.sign_up, payload);
        return response.data;
      },
      onSuccess: () => {
        queryClient.invalidateQueries({ queryKey: ["user/list"] });
        toast.success("Success create user");
      },
    });
  };

  const UpdateUser = () => {
    return useMutation({
      mutationKey: ["user/update"],
      mutationFn: async ({
        user_id,
        data,
      }: {
        user_id: number;
        data: UserUpdateRequest;
      }) => {
        const response = await Api.put(
          endpoint.user.update_user(user_id),
          data
        );
        return response.data;
      },
      onSuccess: () => {
        queryClient.invalidateQueries({ queryKey: ["user/list"] });
        toast.success("Success update user");
      },
    });
  };

  const DeleteUser = () => {
    return useMutation({
      mutationKey: ["user/delete"],
      mutationFn: async (user_id: number) => {
        const response = await Api.put(endpoint.user.delete_user(user_id));
        return response.data;
      },
      onSuccess: () => {
        queryClient.invalidateQueries({ queryKey: ["user/list"] });
        toast.success("Success delete user");
      },
    });
  };

  const GetListEmployee = () => {
    return useQuery({
      queryKey: ["employee/list"],
      queryFn: async () => {
        const response = await Api.get(endpoint.user.list_employee);
        return response.data.data;
      },
      staleTime: Infinity,
    });
  };

  const GetHistoryLog = (type: "general" | "workspace") => {
    return useQuery({
      queryKey: ["user/history-logs", user?.id],
      queryFn: async () => {
        const response = await Api.get(
          `${endpoint.user.history_logs}?type=${type}`
        );
        return response.data.data;
      },
      staleTime: Infinity,
    });
  };

  return {
    getListUser: GetListUser,
    createUser: CreateUser,
    updateUser: UpdateUser,
    deleteUser: DeleteUser,
    getListEmployee: GetListEmployee,
    getHistoryLog: GetHistoryLog,
  };
};
