export interface CreateWorkspaceCardRequest {
  status_id: number;
  title: string;
  description?: string;
}

export interface UpdateWorkspaceCardRequest {
  status_id: number;
  title: string;
  description?: string;
}

export interface CardOrderItem {
  status_id: number;
  card_id: number[];
}

export interface UpdateWorkspaceCardOrderRequest {
  order_status: CardOrderItem[];
}
