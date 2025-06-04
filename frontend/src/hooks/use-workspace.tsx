import { endpoint } from "@/const/endpoint";
import Api from "@/services/api";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { toast } from "sonner";

export const useWorkspace = () => {
  const queryClient = useQueryClient();

  const GetListWorkspace = () => {
    return useQuery({
      queryKey: ["workspace/list"],
      queryFn: async () => {
        const response = await Api.get(endpoint.workspace.list);
        return response.data;
      },
      staleTime: Infinity,
    });
  };

  const CreateWorkspace = () => {
    return useMutation({
      mutationKey: ["workspace/create"],
      mutationFn: async (data: any) => {
        const response = await Api.post(endpoint.workspace.create, data);
        return response.data;
      },
      onSuccess: () => {
        queryClient.invalidateQueries({ queryKey: ["workspace/list"] });
        toast.success("Workspace created successfully");
      },
    });
  };

  const UpdateWorkspace = (id: number) => {
    return useMutation({
      mutationKey: ["workspace/update"],
      mutationFn: async (data: any) => {
        const response = await Api.put(endpoint.workspace.update(id), data);
        return response.data;
      },
      onSuccess: () => {
        queryClient.invalidateQueries({ queryKey: ["workspace/list"] });
        toast.success("Workspace updated successfully");
      },
    });
  };

  const ArchiveWorkspace = () => {
    return useMutation({
      mutationKey: ["workspace/archive"],
      mutationFn: async (id: number) => {
        const response = await Api.put(endpoint.workspace.archive(id));
        return response.data;
      },
      onSuccess: () => {
        queryClient.invalidateQueries({ queryKey: ["workspace/list"] });
        toast.success("Workspace archived successfully");
      },
    });
  };

  const GetListWorkspaceType = () => {
    return useQuery({
      queryKey: ["workspace/type/list"],
      queryFn: async () => {
        const response = await Api.get(endpoint.workspace.type.list);
        return response.data;
      },
      staleTime: Infinity,
    });
  };

  return {
    getListWorkspace: GetListWorkspace,
    createWorkspace: CreateWorkspace,
    updateWorkspace: UpdateWorkspace,
    archiveWorkspace: ArchiveWorkspace,

    getListWorkspaceType: GetListWorkspaceType,
  };
};
