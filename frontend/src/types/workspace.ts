export interface Member {
  id: string;
  name: string;
  username: string;
  email: string;
  role: string;
}

export interface WorkspaceType {
  id: number;
  name: string;
  color: string;
}

export interface WorkspaceCard {
  id: number;
  title: string;
  description: string;
  created_at: string;
  updated_at: string;
}

export interface WorkspaceByStatus extends WorkspaceStatus {
  workspaces: Workspace[];
}

export interface WorkspaceStatus {
  id: number;
  name: string;
  color: string;
  is_archived: boolean;
}

export interface WorkspaceCardByStatus extends WorkspaceStatus {
  cards: WorkspaceCard[];
}

export interface Workspace {
  id: number;
  name: string;
  description: string;
  tags: string[];
  type: WorkspaceType;
  status: WorkspaceStatus;
  members: Member[];
  is_archived: boolean;
  created_at: string;
  updated_at: string;
}
