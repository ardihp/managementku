export interface HistoryLog {
  id: number;
  user: string;
  type: "general" | "workspace";
  action: string;
  description: string;
  created_at: Date;
  update_at: Date;
}
