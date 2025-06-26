import { endpoint } from "@/const/endpoint";
import Api from "@/services/api";
import { useQuery } from "@tanstack/react-query";

export const useWorkspaceType = () => {
  const GetWorkspaceType = () => {
    return useQuery({
      queryKey: ["workspace-type/list"],
      queryFn: async () => {
        const response = await Api.get(endpoint.workspace.type.list);
        return response.data.data;
      },
      staleTime: Infinity,
    });
  };

  return {
    getWorkspaceType: GetWorkspaceType,
  };
};
