import { endpoint } from "@/const/endpoint";
import Api from "@/services/api";
import {
  CreateWorkspaceTypeParams,
  UpdateWorkspaceTypeParams,
} from "@/types/hooks/workspace-type";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { toast } from "sonner";

export const useWorkspaceType = () => {
  const queryClient = useQueryClient();

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

  const CreateWorkspaceType = () => {
    return useMutation({
      mutationKey: ["workspace-type/create"],
      mutationFn: async (data: CreateWorkspaceTypeParams) => {
        const response = await Api.post(endpoint.workspace.type.create, data);
        return response.data.data;
      },
      onSuccess: () => {
        queryClient.invalidateQueries({ queryKey: ["workspace-type/list"] });
        toast.success("Workspace type created successfully");
      },
    });
  };

  const UpdateWorkspaceType = () => {
    return useMutation({
      mutationKey: ["workspace-type/update"],
      mutationFn: async (data: UpdateWorkspaceTypeParams) => {
        const response = await Api.post(endpoint.workspace.type.update, data);
        return response.data.data;
      },
      onSuccess: () => {
        queryClient.invalidateQueries({ queryKey: ["workspace-type/list"] });
        toast.success("Workspace type updated successfully");
      },
    });
  };

  const DeleteWorkspaceType = () => {
    return useMutation({
      mutationKey: ["workspace-type/delete"],
      mutationFn: async (id) => {
        const response = await Api.post(endpoint.workspace.type.update, { id });
        return response.data.data;
      },
      onSuccess: () => {
        queryClient.invalidateQueries({ queryKey: ["workspace-type/list"] });
        toast.success("Workspace type deleted successfully");
      },
    });
  };

  return {
    getWorkspaceType: GetWorkspaceType,
    createWorkspaceType: CreateWorkspaceType,
    updateWorkspaceType: UpdateWorkspaceType,
    deleteWorkspaceType: DeleteWorkspaceType,
  };
};
