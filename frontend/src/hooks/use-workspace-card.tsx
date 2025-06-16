import { endpoint } from "@/const/endpoint";
import Api from "@/services/api";
import {
  CreateWorkspaceCardRequest,
  UpdateWorkspaceCardOrderRequest,
} from "@/types/hooks/workspace-card";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { toast } from "sonner";

export const useWorkspaceCard = () => {
  const queryClient = useQueryClient();

  const GetListWorkspaceCard = (id: number) => {
    return useQuery({
      queryKey: ["workspace/card/list", id],
      queryFn: async () => {
        const response = await Api.get(endpoint.workspace.card.list(id));
        return response.data.data;
      },
      staleTime: Infinity,
      enabled: id !== 0,
    });
  };

  const CreateWorkspaceCard = () => {
    return useMutation({
      mutationKey: ["workspace/card/create"],
      mutationFn: async ({
        id,
        data,
      }: {
        id: number;
        data: CreateWorkspaceCardRequest;
      }) => {
        const response = await Api.post(
          endpoint.workspace.card.create(id),
          data
        );
        return response.data;
      },
      onSuccess: () => {
        queryClient.invalidateQueries({ queryKey: ["workspace/card/list"] });
        toast.success("Card created successfully");
      },
    });
  };

  const UpdateWorkspaceCardStatusOrder = () => {
    return useMutation({
      mutationKey: ["workspace/card/status-order"],
      mutationFn: async ({
        id,
        data,
      }: {
        id: number;
        data: { order_status: number[] };
      }) => {
        const response = await Api.post(
          endpoint.workspace.card.status_order(id),
          data
        );
        return response.data;
      },
    });
  };

  const UpdateWorkspaceCardOrder = () => {
    return useMutation({
      mutationKey: ["workspace/card/card-order"],
      mutationFn: async ({
        id,
        data,
      }: {
        id: number;
        data: UpdateWorkspaceCardOrderRequest;
      }) => {
        const response = await Api.post(
          endpoint.workspace.card.update_order(id),
          data
        );
        return response.data;
      },
      onSuccess: () => {
        queryClient.invalidateQueries({ queryKey: ["workspace/card/list"] });
      },
    });
  };

  return {
    getListWorkspaceCard: GetListWorkspaceCard,
    createWorkspaceCard: CreateWorkspaceCard,
    updateWorkspaceCardStatusOrder: UpdateWorkspaceCardStatusOrder,
    updateWorkspaceCardOrder: UpdateWorkspaceCardOrder,
  };
};
