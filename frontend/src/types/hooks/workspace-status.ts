export interface CreateWorkspaceStatusRequest {
  name: string;
  color: string;
}

export interface UpdateWorkspaceStatusRequest {
  id: number;
  name: string;
  color: string;
}

export interface UpdateWorkspaceStatusOrderRequest {
  order_status: number[];
}
