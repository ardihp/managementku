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

  return {
    getListEmployee: GetListEmployee,
  };
};
