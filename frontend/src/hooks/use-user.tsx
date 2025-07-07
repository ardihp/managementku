import { endpoint } from "@/const/endpoint";
import Api from "@/services/api";
import { useQuery } from "@tanstack/react-query";
import { useAuth } from "./auth/use-auth";

export const useUser = () => {
  const { user } = useAuth();

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
    getListEmployee: GetListEmployee,
    getHistoryLog: GetHistoryLog,
  };
};
