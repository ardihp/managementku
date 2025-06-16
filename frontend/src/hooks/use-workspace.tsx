import { endpoint } from "@/const/endpoint";
import Api from "@/services/api";
import { UpdateWorkspace, UpdateWorkspaceOrder } from "@/types/hooks/workspace";
import { Workspace, WorkspaceStatus } from "@/types/workspace";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { toast } from "sonner";

export const useWorkspace = () => {
  const queryClient = useQueryClient();

  const GetListWorkspace = () => {
    return useQuery({
      queryKey: ["workspace/list"],
      queryFn: async () => {
        const response = await Api.get(endpoint.workspace.list);
        return response.data.data;
      },
      staleTime: Infinity,
    });
  };

  const GetDetailWorkspace = (id: number) => {
    return useQuery({
      queryKey: ["workspace/detail", id],
      queryFn: async () => {
        const response: { data: { data: Workspace } } = await Api.get(
          `${endpoint.workspace.list}/${id}`
        );
        return response.data.data;
      },
      staleTime: Infinity,
      enabled: id !== 0,
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

  const UpdateWorkspace = () => {
    return useMutation({
      mutationKey: ["workspace/update"],
      mutationFn: async ({
        id,
        data,
      }: {
        id: number;
        data: UpdateWorkspace;
      }) => {
        const response = await Api.put(endpoint.workspace.update(id), data);
        return response.data;
      },
      onSuccess: () => {
        queryClient.invalidateQueries({ queryKey: ["workspace/list"] });
        toast.success("Workspace updated successfully");
      },
    });
  };

  const UpdateWorkspaceOrder = () => {
    return useMutation({
      mutationKey: ["worspace/update-order"],
      mutationFn: async (data: UpdateWorkspaceOrder) => {
        const response = await Api.post(endpoint.workspace.update_order, data);
        return response.data;
      },
      onSuccess: () =>
        queryClient.invalidateQueries({ queryKey: ["workspace/list"] }),
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
        queryClient.invalidateQueries({
          queryKey: ["workspace/archived/list"],
        });
        toast.success("Workspace archived successfully");
      },
    });
  };

  const UnarchiveWorkspace = () => {
    return useMutation({
      mutationKey: ["workspace/unarchive"],
      mutationFn: async (id: number) => {
        const response = await Api.put(endpoint.workspace.unarchive(id));
        return response.data;
      },
      onSuccess: () => {
        queryClient.invalidateQueries({ queryKey: ["workspace/list"] });
        queryClient.invalidateQueries({
          queryKey: ["workspace/archived/list"],
        });
        toast.success("Workspace activated successfully");
      },
    });
  };

  const DeleteWorkspace = () => {
    return useMutation({
      mutationKey: ["workspace/delete"],
      mutationFn: async (id: number) => {
        const response = await Api.delete(endpoint.workspace.delete(id));
        return response.data;
      },
      onSuccess: () => {
        queryClient.invalidateQueries({ queryKey: ["workspace/list"] });
        queryClient.invalidateQueries({
          queryKey: ["workspace/archived/list"],
        });
        toast.success("Workspace deleted successfully");
      },
    });
  };

  const GetListWorkspaceType = () => {
    return useQuery({
      queryKey: ["workspace/type/list"],
      queryFn: async () => {
        const response = await Api.get(endpoint.workspace.type.list);
        return response.data.data;
      },
      staleTime: Infinity,
    });
  };

  const GetListWorkspaceStatus = (is_archived: "true" | "false") => {
    return useQuery({
      queryKey: ["workspace/status/list", is_archived],
      queryFn: async () => {
        const response = await Api.get(
          `${endpoint.workspace.status.list}?is_archived=${is_archived}`
        );
        return response.data.data;
      },
      staleTime: Infinity,
    });
  };

  const GetListWorkspaceCardStatus = (
    id: number,
    is_archived: "true" | "false"
  ) => {
    return useQuery({
      queryKey: ["workspace/card-status/list", id, is_archived],
      queryFn: async () => {
        const response = await Api.get(
          `${endpoint.workspace.card_status.list(
            id
          )}?is_archived=${is_archived}`
        );
        return response.data.data;
      },
      staleTime: Infinity,
      enabled: id !== 0,
    });
  };

  const GetListArchivedWorkspace = () => {
    return useQuery({
      queryKey: ["workspace/archived/list"],
      queryFn: async () => {
        const response = await Api.get(endpoint.workspace.list_archived);
        return response.data.data;
      },
      staleTime: Infinity,
    });
  };

  return {
    getListWorkspace: GetListWorkspace,
    getDetailWorkspace: GetDetailWorkspace,
    createWorkspace: CreateWorkspace,
    updateWorkspace: UpdateWorkspace,
    updateWorkspaceOrder: UpdateWorkspaceOrder,
    archiveWorkspace: ArchiveWorkspace,
    unarchiveWorkspace: UnarchiveWorkspace,
    deleteWorkspace: DeleteWorkspace,
    getListWorkspaceStatus: GetListWorkspaceStatus,
    getListArchivedWorkspace: GetListArchivedWorkspace,

    getListWorkspaceType: GetListWorkspaceType,

    getListWorkspaceCardStatus: GetListWorkspaceCardStatus,
  };
};
