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

    status: {
      list: "/workspace/status",
      create: "/workspace/status/create",
    },

    type: {
      list: "/workspace/type",
      create: "/workspace/type",
      delete: (id: number) => `/workspace/type/${id}`,
    },

    card: {
      list: (id: number) => `/workspace/${id}/card`,
      create: (id: number) => `/workspace/${id}/card/create`,
      status_order: (id: number) => `/workspace/${id}/card/status-order`,
      update_order: (id: number) => `/workspace/${id}/card/update-order`,
    },
  },
};
