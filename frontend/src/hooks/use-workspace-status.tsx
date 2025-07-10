import { endpoint } from "@/const/endpoint";
import Api from "@/services/api";
import {
  CreateWorkspaceStatusRequest,
  UpdateWorkspaceStatusOrderRequest,
  UpdateWorkspaceStatusRequest,
} from "@/types/hooks/workspace-status";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { toast } from "sonner";
import { useAuth } from "./auth/use-auth";

export const useWorkspaceStatus = () => {
  const queryClient = useQueryClient();
  const { user } = useAuth();

  const GetListWorkspaceStatus = () => {
    return useQuery({
      queryKey: ["workspace-status/list"],
      queryFn: async () => {
        const response = await Api.get(endpoint.workspace.status.list);
        return response.data.data;
      },
      staleTime: Infinity,
    });
  };

  const GetListArchivedWorkspaceStatus = () => {
    return useQuery({
      queryKey: ["workspace-status/list-archived"],
      queryFn: async () => {
        const response = await Api.get(endpoint.workspace.status.list_archived);
        return response.data.data;
      },
      staleTime: Infinity,
      enabled: user?.role === "manager",
    });
  };

  const CreateWorkspaceStatus = () => {
    return useMutation({
      mutationKey: ["workspace-status/create"],
      mutationFn: async (data: CreateWorkspaceStatusRequest) => {
        const response = await Api.post(endpoint.workspace.status.create, data);
        return response.data;
      },
      onSuccess: () => {
        queryClient.invalidateQueries({ queryKey: ["workspace-status/list"] });
        queryClient.invalidateQueries({ queryKey: ["workspace/list"] });
        toast.success("Workspace status created successfully");
      },
    });
  };

  const UpdateWorkspaceStatus = () => {
    return useMutation({
      mutationKey: ["workspace-status/update"],
      mutationFn: async (data: UpdateWorkspaceStatusRequest) => {
        const response = await Api.put(endpoint.workspace.status.update, data);
        return response.data;
      },
      onSuccess: () => {
        queryClient.invalidateQueries({ queryKey: ["workspace-status/list"] });
        queryClient.invalidateQueries({ queryKey: ["workspace/list"] });
        toast.success("Workspace status updated successfully");
      },
    });
  };

  const UpdateWorkspaceStatusOrder = () => {
    return useMutation({
      mutationKey: ["workspace-status/update-order"],
      mutationFn: async (data: UpdateWorkspaceStatusOrderRequest) => {
        const response = await Api.put(
          endpoint.workspace.status.update_order,
          data
        );
        return response.data;
      },
      onSuccess: () => {
        queryClient.invalidateQueries({ queryKey: ["workspace/list"] });
        queryClient.invalidateQueries({ queryKey: ["user/history-logs"] });
      },
    });
  };

  const ArchiveWorkspaceStatus = () => {
    return useMutation({
      mutationKey: ["workspace-status/archive"],
      mutationFn: async (id: number) => {
        const response = await Api.put(endpoint.workspace.status.archive(id));
        return response.data;
      },
      onSuccess: () => {
        queryClient.invalidateQueries({ queryKey: ["workspace-status/list"] });
        queryClient.invalidateQueries({
          queryKey: ["workspace-status/list-archived"],
        });
        queryClient.invalidateQueries({ queryKey: ["workspace/list"] });
        toast.success("Workspace status archived successfully");
      },
    });
  };

  const UnarchiveWorkspaceStatus = () => {
    return useMutation({
      mutationKey: ["workspace-status/unarchive"],
      mutationFn: async (id: number) => {
        const response = await Api.put(endpoint.workspace.status.unarchive(id));
        return response.data;
      },
      onSuccess: () => {
        queryClient.invalidateQueries({ queryKey: ["workspace-status/list"] });
        queryClient.invalidateQueries({
          queryKey: ["workspace-status/list-archived"],
        });
        queryClient.invalidateQueries({ queryKey: ["workspace/list"] });
        toast.success("Workspace status activated successfully");
      },
    });
  };

  const DeleteWorkspaceStatus = () => {
    return useMutation({
      mutationKey: ["workspace-status/delete"],
      mutationFn: async (id: number) => {
        const response = await Api.delete(endpoint.workspace.status.delete(id));
        return response.data;
      },
      onSuccess: () => {
        queryClient.invalidateQueries({ queryKey: ["workspace-status/list"] });
        queryClient.invalidateQueries({
          queryKey: ["workspace-status/list-archived"],
        });
        queryClient.invalidateQueries({ queryKey: ["workspace/list"] });
        toast.success("Workspace status deleted successfully");
      },
    });
  };

  return {
    getListWorkspaceStatus: GetListWorkspaceStatus,
    getListArchivedWorkspaceStatus: GetListArchivedWorkspaceStatus,
    createWorkspaceStatus: CreateWorkspaceStatus,
    updateWorkspaceStatus: UpdateWorkspaceStatus,
    updateWorkspaceStatusOrder: UpdateWorkspaceStatusOrder,
    archiveWorkspaceStatus: ArchiveWorkspaceStatus,
    unarchiveWorkspaceStatus: UnarchiveWorkspaceStatus,
    deleteWorkspaceStatus: DeleteWorkspaceStatus,
  };
};
