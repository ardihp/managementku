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

    type: {
      list: "/workspace/type",
      create: "/workspace/type",
      delete: (id: number) => `/workspace/type/${id}`,
    },
  },
};
