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
}

export interface WorkspaceCardByStatus extends WorkspaceStatus {
  cards: WorkspaceCard[];
}
