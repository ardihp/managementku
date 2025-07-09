import { endpoint } from "@/const/endpoint";
import Api from "@/services/api";
import {
  AssignWorkspaceMemberRequest,
  UpdateWorkspace,
  UpdateWorkspaceOrder,
} from "@/types/hooks/workspace";
import { Workspace } from "@/types/workspace";
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
        queryClient.invalidateQueries({ queryKey: ["user/history-logs"] });
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
        queryClient.invalidateQueries({ queryKey: ["user/history-logs"] });
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
      onSuccess: () => {
        queryClient.invalidateQueries({ queryKey: ["workspace/list"] });
        queryClient.invalidateQueries({ queryKey: ["user/history-logs"] });
      },
    });
  };

  const AssignWorkspaceMember = () => {
    return useMutation({
      mutationKey: ["workspace/assign"],
      mutationFn: async (data: AssignWorkspaceMemberRequest) => {
        const response = await Api.post(endpoint.workspace.assign_member, data);
        return response.data;
      },
      onSuccess: () => {
        queryClient.invalidateQueries({ queryKey: ["workspace/list"] });
        queryClient.invalidateQueries({ queryKey: ["user/history-logs"] });
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
        queryClient.invalidateQueries({
          queryKey: ["workspace/archived/list"],
        });
        queryClient.invalidateQueries({ queryKey: ["user/history-logs"] });
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
        queryClient.invalidateQueries({ queryKey: ["user/history-logs"] });
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
        queryClient.invalidateQueries({ queryKey: ["user/history-logs"] });
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
    assignWorkspaceMember: AssignWorkspaceMember,
    archiveWorkspace: ArchiveWorkspace,
    unarchiveWorkspace: UnarchiveWorkspace,
    deleteWorkspace: DeleteWorkspace,

    getListArchivedWorkspace: GetListArchivedWorkspace,
    getListWorkspaceType: GetListWorkspaceType,
  };
};
