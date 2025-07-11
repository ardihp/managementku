import { endpoint } from "@/const/endpoint";
import Api from "@/services/api";
import {
  AssignWorkspaceCardMemberRequest,
  CreateWorkspaceCardRequest,
  UpdateWorkspaceCardOrderRequest,
  UpdateWorkspaceCardRequest,
} from "@/types/hooks/workspace-card";
import { WorkspaceCard } from "@/types/workspace-card";
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

  const GetListWorkspaceCardMember = (card_id: number) => {
    return useQuery({
      queryKey: ["workspace-card/list-member", card_id],
      queryFn: async () => {
        const response = await Api.get(
          endpoint.workspace.card.list_member(card_id)
        );
        return response.data.data;
      },
      staleTime: Infinity,
      enabled: card_id !== 0,
    });
  };

  const AssignWorkspaceCardMember = (workspace_id: number, card_id: number) => {
    return useMutation({
      mutationKey: ["workspace-card/assign"],
      mutationFn: async (data: AssignWorkspaceCardMemberRequest) => {
        const response = await Api.post(
          endpoint.workspace.card.assign_member(workspace_id),
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
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/list-activity", card_id],
        });
        queryClient.invalidateQueries({ queryKey: ["user/history-logs"] });
      },
    });
  };

  const WatchWorkspaceCard = (workspace_id: number, card_id: number) => {
    return useMutation({
      mutationKey: ["workspace-card/watch"],
      mutationFn: async (data: { user_id: number }) => {
        const response = await Api.put(
          endpoint.workspace.card.watch(workspace_id, card_id),
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
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/list-activity", card_id],
        });
        queryClient.invalidateQueries({ queryKey: ["user/history-logs"] });
      },
    });
  };

  const UnwatchWorkspaceCard = (workspace_id: number, card_id: number) => {
    return useMutation({
      mutationKey: ["workspace-card/unwatch"],
      mutationFn: async (data: { user_id: number }) => {
        const response = await Api.put(
          endpoint.workspace.card.unwatch(workspace_id, card_id),
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
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/list-activity", card_id],
        });
        queryClient.invalidateQueries({ queryKey: ["user/history-logs"] });
      },
    });
  };

  const GetListWorkspaceCardActivities = (
    workspace_id: number,
    card_id: number
  ) => {
    return useQuery({
      queryKey: ["workspace-card/list-activity", card_id],
      queryFn: async () => {
        const response = await Api.get(
          endpoint.workspace.card.activity.list(workspace_id, card_id)
        );
        return response.data.data;
      },
      staleTime: Infinity,
      enabled: card_id !== 0,
    });
  };

  const CreateWorkspaceCardActivity = (
    workspace_id: number,
    card_id: number
  ) => {
    return useMutation({
      mutationKey: ["workspace-card/create-activity"],
      mutationFn: async (data: { message: string }) => {
        const response = await Api.post(
          endpoint.workspace.card.activity.create(workspace_id, card_id),
          data
        );
        return response.data;
      },
      onSuccess: () => {
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/list", workspace_id],
        });
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/list-activity", card_id],
        });
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/history-log", workspace_id],
        });
      },
    });
  };

  const UpdateWorkspaceCardActivity = (
    workspace_id: number,
    card_id: number
  ) => {
    return useMutation({
      mutationKey: ["workspace-card/update-activity"],
      mutationFn: async ({
        activity_id,
        data,
      }: {
        activity_id: number;
        data: { message: string };
      }) => {
        const response = await Api.put(
          endpoint.workspace.card.activity.update(
            workspace_id,
            card_id,
            activity_id
          ),
          data
        );
        return response.data.data;
      },
      onSuccess: () => {
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/list", workspace_id],
        });
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/list-activity", card_id],
        });
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/list-archived", workspace_id],
        });
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/history-log", workspace_id],
        });
      },
    });
  };

  const DeleteWorkspaceCardActivity = (
    workspace_id: number,
    card_id: number
  ) => {
    return useMutation({
      mutationKey: ["workspace-card/delete-activity"],
      mutationFn: async (activity_id: number) => {
        const response = await Api.delete(
          endpoint.workspace.card.activity.delete(
            workspace_id,
            card_id,
            activity_id
          )
        );
        return response.data.data;
      },
      onSuccess: () => {
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/list", workspace_id],
        });
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/list-activity", card_id],
        });
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/list-archived", workspace_id],
        });
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/history-log", workspace_id],
        });
      },
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
      onSuccess: (response: { data: WorkspaceCard }) => {
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/list", workspace_id],
        });
        queryClient.invalidateQueries({
          queryKey: ["workspace-card/list-activity", response.data.id],
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
    getListWorkspaceCardMember: GetListWorkspaceCardMember,
    assignWorkspaceCardMember: AssignWorkspaceCardMember,
    watchWorkspaceCard: WatchWorkspaceCard,
    unwatchWorkspaceCard: UnwatchWorkspaceCard,
    getListWorkspaceCardActivities: GetListWorkspaceCardActivities,
    createWorkspaceCardActivity: CreateWorkspaceCardActivity,
    updateWorkspaceCardActivity: UpdateWorkspaceCardActivity,
    deleteWorkspaceCardActivity: DeleteWorkspaceCardActivity,
    createWorkspaceCard: CreateWorkspaceCard,
    updateWorkspaceCard: UpdateWorkspaceCard,
    archivedWorkspaceCard: ArchivedWorkspaceCard,
    unarchivedWorkspaceCard: UnarchivedWorkspaceCard,
    deleteWorkspaceCard: DeleteWorkspaceCard,
    updateWorkspaceCardStatusOrder: UpdateWorkspaceCardStatusOrder,
    updateWorkspaceCardOrder: UpdateWorkspaceCardOrder,
  };
};
