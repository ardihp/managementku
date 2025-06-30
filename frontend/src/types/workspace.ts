import { colorsType } from "@/const";

export interface Member {
  id: number;
  name: string;
  username: string;
  email: string;
  role: string;
}

export interface WorkspaceType {
  id: number;
  name: string;
  color: colorsType;
}

export interface WorkspaceRole {
  id: number;
  name: string;
}

export interface WorkspaceByStatus extends WorkspaceStatus {
  workspaces: Workspace[];
}

export interface WorkspaceStatus {
  id: number;
  name: string;
  color: colorsType;
  is_archived: boolean;
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
