import { endpoint } from "@/const/endpoint";
import Api from "@/services/api";
import { useQuery } from "@tanstack/react-query";

export const useWorkspace = () => {
  const GetListWorkspace = () => {
    return useQuery({
      queryFn: async () => {
        const response = await Api.get(endpoint.workspace.list);

        return response.data;
      },
      queryKey: ["workspaces"],
    });
  };

  return { getListWorkspace: GetListWorkspace };
};
