import { endpoint } from "@/const/endpoint";
import Api from "@/services/api";
import {
  CreateWorkspaceCardStatusRequest,
  UpdateWorkspaceCardStatusRequest,
} from "@/types/hooks/workspace-card-status";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { toast } from "sonner";

export const useWorkspaceCardStatus = () => {
  const queryClient = useQueryClient();

  const GetListWorkspaceCardStatus = (workspace_id: number) => {
    return useQuery({
      queryKey: ["workspace-card-status/list"],
      queryFn: async () => {
        const response = await Api.get(
          endpoint.workspace.card.status.list(workspace_id)
        );
        return response.data.data;
      },
      enabled: workspace_id !== 0,
      staleTime: Infinity,
    });
  };

  const GetListArchivedWorkspaceCardStatus = (workspace_id: number) => {
    return useQuery({
      queryKey: ["workspace-card-status/list-archived"],
      queryFn: async () => {
        const response = await Api.get(
          endpoint.workspace.card.status.list_archived(workspace_id)
        );
        return response.data.data;
      },
      enabled: workspace_id !== 0,
      staleTime: Infinity,
    });
  };

  const CreateWorkspaceCardStatus = (workspace_id: number) => {
    return useMutation({
      mutationKey: ["workspace-card-status/create"],
      mutationFn: async (data: CreateWorkspaceCardStatusRequest) => {
        const response = await Api.post(
          endpoint.workspace.card.status.create(workspace_id),
          data
        );
        return response.data;
      },
      onSuccess: () => {
        queryClient.invalidateQueries({
          queryKey: ["workspace-card-status/list"],
        });
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/list", workspace_id],
        });
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/history-log", workspace_id],
        });
        toast.success("Status created successfully");
      },
    });
  };

  const UpdateWorkspaceCardStatus = (workspace_id: number) => {
    return useMutation({
      mutationKey: ["workspace-card-status/update"],
      mutationFn: async ({
        data,
        status_id,
      }: {
        data: UpdateWorkspaceCardStatusRequest;
        status_id: number;
      }) => {
        const response = await Api.put(
          endpoint.workspace.card.status.update(workspace_id, status_id),
          data
        );
        return response.data;
      },
      onSuccess: () => {
        queryClient.invalidateQueries({
          queryKey: ["workspace-card-status/list"],
        });
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/list", workspace_id],
        });
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/history-log", workspace_id],
        });
        toast.success("Status updated successfully");
      },
    });
  };

  const ArchiveWorkspaceCardStatus = (workspace_id: number) => {
    return useMutation({
      mutationKey: ["workspace-card-status/archive"],
      mutationFn: async (status_id: number) => {
        const response = await Api.put(
          endpoint.workspace.card.status.archive(workspace_id, status_id)
        );
        return response.data;
      },
      onSuccess: () => {
        queryClient.invalidateQueries({
          queryKey: ["workspace-card-status/list-archived"],
        });
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/list", workspace_id],
        });
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/history-log", workspace_id],
        });
        toast.success("Status archived successfully");
      },
    });
  };

  const UnarchiveWorkspaceCardStatus = (workspace_id: number) => {
    return useMutation({
      mutationKey: ["workspace-card-status/unarchive"],
      mutationFn: async (status_id: number) => {
        const response = await Api.put(
          endpoint.workspace.card.status.unarchive(workspace_id, status_id)
        );
        return response.data;
      },
      onSuccess: () => {
        queryClient.invalidateQueries({
          queryKey: ["workspace-card-status/list-archived"],
        });
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/list", workspace_id],
        });
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/history-log", workspace_id],
        });
        toast.success("Status activated successfully");
      },
    });
  };

  const DeleteWorkspaceCardStatus = (workspace_id: number) => {
    return useMutation({
      mutationKey: ["workspace-card-status/delete"],
      mutationFn: async (status_id: number) => {
        const response = await Api.delete(
          endpoint.workspace.card.status.delete(workspace_id, status_id)
        );
        return response.data;
      },
      onSuccess: () => {
        queryClient.invalidateQueries({
          queryKey: ["workspace-card-status/list-archived"],
        });
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/list", workspace_id],
        });
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/history-log", workspace_id],
        });
        toast.success("Status deleted successfully");
      },
    });
  };

  return {
    getListWorkspaceCardStatus: GetListWorkspaceCardStatus,
    getListArchivedWorkspaceCardStatus: GetListArchivedWorkspaceCardStatus,
    createWorkspaceCardStatus: CreateWorkspaceCardStatus,
    updateWorkspaceCardStatus: UpdateWorkspaceCardStatus,
    archiveWorkspaceCardStatus: ArchiveWorkspaceCardStatus,
    unarchiveWorkspaceCardStatus: UnarchiveWorkspaceCardStatus,
    deleteWorkspaceCardStatus: DeleteWorkspaceCardStatus,
  };
};
