import { endpoint } from "@/const/endpoint";
import Api from "@/services/api";
import {
  CreateWorkspaceCardRequest,
  UpdateWorkspaceCardOrderRequest,
  UpdateWorkspaceCardRequest,
} from "@/types/hooks/workspace-card";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { toast } from "sonner";

export const useWorkspaceCard = () => {
  const queryClient = useQueryClient();

  const GetListWorkspaceHistoryLog = (workspace_id: number) => {
    return useQuery({
      queryKey: ["workspace-card/history-log", workspace_id],
      queryFn: async () => {
        const response = await Api.get(
          endpoint.workspace.card.history_log(workspace_id)
        );
        return response.data.data;
      },
      staleTime: Infinity,
      enabled: workspace_id !== 0,
    });
  };

  const GetListWorkspaceCard = (workspace_id: number) => {
    return useQuery({
      queryKey: ["workspace-card/list", workspace_id],
      queryFn: async () => {
        const response = await Api.get(
          endpoint.workspace.card.list(workspace_id)
        );
        return response.data.data;
      },
      staleTime: Infinity,
      enabled: workspace_id !== 0,
    });
  };

  const GetListArchivedWorkspaceCard = (workspace_id: number) => {
    return useQuery({
      queryKey: ["workspace-card/list-archived", workspace_id],
      queryFn: async () => {
        const response = await Api.get(
          endpoint.workspace.card.list_archived(workspace_id)
        );
        return response.data.data;
      },
      staleTime: Infinity,
      enabled: workspace_id !== 0,
    });
  };

  const CreateWorkspaceCard = (workspace_id: number) => {
    return useMutation({
      mutationKey: ["workspace-card/create"],
      mutationFn: async ({ data }: { data: CreateWorkspaceCardRequest }) => {
        const response = await Api.post(
          endpoint.workspace.card.create(workspace_id),
          data
        );
        return response.data;
      },
      onSuccess: () => {
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/list", workspace_id],
        });
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/history-log", workspace_id],
        });
        toast.success("Card created successfully");
      },
    });
  };

  const UpdateWorkspaceCard = (workspace_id: number) => {
    return useMutation({
      mutationKey: ["workspace-card/update"],
      mutationFn: async ({
        card_id,
        data,
      }: {
        card_id: number;
        data: UpdateWorkspaceCardRequest;
      }) => {
        const response = await Api.put(
          endpoint.workspace.card.update(workspace_id, card_id),
          data
        );
        return response.data;
      },
      onSuccess: () => {
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/list", workspace_id],
        });
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/list-archived", workspace_id],
        });
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/history-log", workspace_id],
        });
        toast.success("Card updated successfully");
      },
    });
  };

  const ArchivedWorkspaceCard = (workspace_id: number) => {
    return useMutation({
      mutationKey: ["workspace-card/archive"],
      mutationFn: async (card_id: number) => {
        const response = await Api.put(
          endpoint.workspace.card.archive(workspace_id, card_id)
        );
        return response.data.data;
      },
      onSuccess: () => {
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/list", workspace_id],
        });
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/list-archived", workspace_id],
        });
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/history-log", workspace_id],
        });
        toast.success("Card archived successfully");
      },
    });
  };

  const UnarchivedWorkspaceCard = (workspace_id: number) => {
    return useMutation({
      mutationKey: ["workspace-card/unarchive"],
      mutationFn: async (card_id: number) => {
        const response = await Api.put(
          endpoint.workspace.card.unarchive(workspace_id, card_id)
        );
        return response.data.data;
      },
      onSuccess: () => {
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/list", workspace_id],
        });
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/list-archived", workspace_id],
        });
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/history-log", workspace_id],
        });
        toast.success("Card activated successfully");
      },
    });
  };

  const DeleteWorkspaceCard = (workspace_id: number) => {
    return useMutation({
      mutationKey: ["workspace-card/delete"],
      mutationFn: async (card_id: number) => {
        const response = await Api.delete(
          endpoint.workspace.card.delete(workspace_id, card_id)
        );
        return response.data.data;
      },
      onSuccess: () => {
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/list", workspace_id],
        });
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/list-archived", workspace_id],
        });
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/history-log", workspace_id],
        });
        toast.success("Card deleted successfully");
      },
    });
  };

  const UpdateWorkspaceCardStatusOrder = (workspace_id: number) => {
    return useMutation({
      mutationKey: ["workspace-card/status-order"],
      mutationFn: async ({ data }: { data: { order_status: number[] } }) => {
        const response = await Api.post(
          endpoint.workspace.card.status_order(workspace_id),
          data
        );
        return response.data;
      },
      onSuccess: () => {
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/list", workspace_id],
        });
      },
    });
  };

  const UpdateWorkspaceCardOrder = (workspace_id: number) => {
    return useMutation({
      mutationKey: ["workspace-card/card-order"],
      mutationFn: async ({
        data,
      }: {
        data: UpdateWorkspaceCardOrderRequest;
      }) => {
        const response = await Api.post(
          endpoint.workspace.card.update_order(workspace_id),
          data
        );
        return response.data;
      },
      onSuccess: () => {
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/list", workspace_id],
        });
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/history-log", workspace_id],
        });
      },
    });
  };

  return {
    getListWorkspaceHistoryLog: GetListWorkspaceHistoryLog,
    getListWorkspaceCard: GetListWorkspaceCard,
    getListArchivedWorkspaceCard: GetListArchivedWorkspaceCard,
    createWorkspaceCard: CreateWorkspaceCard,
    updateWorkspaceCard: UpdateWorkspaceCard,
    archivedWorkspaceCard: ArchivedWorkspaceCard,
    unarchivedWorkspaceCard: UnarchivedWorkspaceCard,
    deleteWorkspaceCard: DeleteWorkspaceCard,
    updateWorkspaceCardStatusOrder: UpdateWorkspaceCardStatusOrder,
    updateWorkspaceCardOrder: UpdateWorkspaceCardOrder,
  };
};
