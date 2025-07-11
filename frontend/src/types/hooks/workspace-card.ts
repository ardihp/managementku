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
  card_id: number;
  order_status: CardOrderItem[];
}

export interface AssignWorkspaceCardMemberRequest {
  card_id: number;
  user_id: number[];
}
