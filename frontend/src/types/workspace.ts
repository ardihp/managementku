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

export interface Workspace {
  id: number;
  name: string;
  description: string;
  tags: string[];
  type: WorkspaceType;
  members: Member[];
  created_at: string;
  updated_at: string;
}
