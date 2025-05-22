export interface Member {
  id: string;
  name: string;
  username: string;
  email: string;
  role: string;
}

export interface Workspace {
  id: number;
  name: string;
  desription: string;
  tags: string[];
  members: Member[];
  created_at: string;
  updated_at: string;
}
