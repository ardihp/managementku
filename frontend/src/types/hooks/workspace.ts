export interface UpdateWorkspace {
  name: string;
  description: string;
  type_id: number;
  status_id: number;
  tags?: string[];
}

export interface OrderStatusItem {
  status_id: number;
  workspace_id: number[];
}

export interface UpdateWorkspaceOrder {
  order_status: OrderStatusItem[];
}
