export const endpoint = {
  auth: {
    sign_in: "/login",
    sign_up: "/register",
  },
  workspace: {
    list: "/workspace",
    create: "/workspace/create",
    update: (id: number) => `/workspace/update/${id}`,
    archive: (id: number) => `/workspace/archive/${id}`,
    unarchive: (id: number) => `/workspace/unarchive/${id}`,
    delete: (id: number) => `/workspace/delete/${id}`,
    list_archived: "/workspace/archived",
    update_order: "/workspace/update-order",
    assign_member: "/workspace/assign",

    status: {
      list: "/workspace/status",
      list_archived: "/workspace/status/archived",
      create: "/workspace/status/create",
      update: "/workspace/status/update",
      update_order: "/workspace/status/update-order",
      archive: (id: number) => `/workspace/status/archive/${id}`,
      unarchive: (id: number) => `/workspace/status/unarchive/${id}`,
      delete: (id: number) => `/workspace/status/delete/${id}`,
    },

    type: {
      list: "/workspace/type",
      create: "/workspace/type/create",
      update: "/workspace/type/update",
      delete: (id: number) => `/workspace/type/delete/${id}`,
    },

    card: {
      history_log: (id: number) => `/workspace/${id}/history-logs`,
      list: (id: number) => `/workspace/${id}/card`,
      list_archived: (id: number) => `/workspace/${id}/card/archived`,
      create: (id: number) => `/workspace/${id}/card/create`,
      update: (workspace_id: number, card_id: number) =>
        `/workspace/${workspace_id}/card/update/${card_id}`,
      archive: (workspace_id: number, card_id: number) =>
        `/workspace/${workspace_id}/card/archive/${card_id}`,
      unarchive: (workspace_id: number, card_id: number) =>
        `/workspace/${workspace_id}/card/unarchive/${card_id}`,
      delete: (workspace_id: number, card_id: number) =>
        `/workspace/${workspace_id}/card/delete/${card_id}`,
      status_order: (id: number) => `/workspace/${id}/card/status-order`,
      update_order: (id: number) => `/workspace/${id}/card/update-order`,

      status: {
        list: (workspace_id: number) => `/workspace/${workspace_id}/status`,
        list_archived: (workspace_id: number) =>
          `/workspace/${workspace_id}/status/archived`,
        create: (workspace_id: number) =>
          `/workspace/${workspace_id}/status/create`,
        update: (workspace_id: number, status_id: number) =>
          `/workspace/${workspace_id}/status/update/${status_id}`,
        archive: (workspace_id: number, status_id: number) =>
          `/workspace/${workspace_id}/status/archive/${status_id}`,
        unarchive: (workspace_id: number, status_id: number) =>
          `/workspace/${workspace_id}/status/unarchive/${status_id}`,
        delete: (workspace_id: number, status_id: number) =>
          `/workspace/${workspace_id}/status/delete/${status_id}`,
      },
    },
  },
  user: {
    list_employee: "/employee",
    history_logs: "/history-logs",
  },
};
