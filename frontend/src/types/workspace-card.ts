import { Member, WorkspaceStatus } from "./workspace";

export interface WorkspaceCardMember {
  id: number;
  name: string;
  username: string;
  email: string;
}

export interface WorkspaceCard {
  id: number;
  title: string;
  description: string;
  status: WorkspaceStatus;
  members: WorkspaceCardMember[];
  comment_count: number;
  is_archived: boolean;
  is_completed: boolean;
  user_created: string;
  created_at: Date;
  updated_at: Date;
}

export interface WorkspaceCardByStatus extends WorkspaceStatus {
  cards: WorkspaceCard[];
}

export interface WorkspaceCardActivity {
  id: number;
  user_created: string;
  message: string;
  type: string;
  created_at: Date;
  updated_at: Date;
}
