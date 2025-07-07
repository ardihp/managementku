import { endpoint } from "@/const/endpoint";
import Api from "@/services/api";
import { useQuery } from "@tanstack/react-query";

export const useUser = () => {
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
      queryKey: ["user/history-logs"],
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
