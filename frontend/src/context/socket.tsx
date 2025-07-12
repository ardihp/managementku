import { useAuth } from "@/hooks/auth/use-auth";
import { useQueryClient } from "@tanstack/react-query";
import { createContext, ReactNode, useEffect, useState } from "react";

export const SocketContext = createContext<{
  socket: WebSocket | null;
}>({
  socket: null,
});

export default function SocketProvider({ children }: { children: ReactNode }) {
  const [socket, setSocket] = useState<WebSocket | null>(null);
  const { user } = useAuth();
  const queryClient = useQueryClient();

  useEffect(() => {
    const ws = new WebSocket("ws://localhost:8000/ws");

    if (user) {
      setSocket(ws);
    }
  }, []);

  useEffect(() => {
    if (socket) {
      socket.onmessage = (event) => {
        const content = JSON.parse(event.data)?.content;

        if (content?.type === "card-activity") {
          queryClient.invalidateQueries({
            queryKey: ["workspace-card/list-activity", content?.card_id],
          });
          queryClient.invalidateQueries({
            queryKey: ["workspace-card/list", content?.workspace_id],
          });
        }

        if (content?.type === "detail-card") {
          queryClient.invalidateQueries({
            queryKey: ["workspace-card/list", content?.workspace_id],
          });
        }

        if (content?.type === "mark-progress") {
          queryClient.invalidateQueries({
            queryKey: ["workspace-card/list", content?.workspace_id],
          });
          queryClient.invalidateQueries({
            queryKey: ["workspace-card/list-archived", content?.workspace_id],
          });
          queryClient.invalidateQueries({
            queryKey: [
              "workspace-card-status/list-archived",
              content?.workspace_id,
            ],
          });
          queryClient.invalidateQueries({ queryKey: ["workspace/list"] });
        }

        if (content?.type === "refresh-workspace") {
          queryClient.invalidateQueries({ queryKey: ["workspace/list"] });
          queryClient.invalidateQueries({
            queryKey: ["workspace/archived/list"],
          });
          queryClient.invalidateQueries({
            queryKey: ["workspace-status/list"],
          });
        }

        if (content?.type === "assign-workspace") {
          queryClient.invalidateQueries({ queryKey: ["workspace/list"] });
          queryClient.invalidateQueries({
            queryKey: ["workspace/detail", content?.workspace_id],
          });
        }

        if (content?.workspace_id) {
          queryClient.invalidateQueries({
            queryKey: ["workspace-card/history-log", content?.workspace_id],
          });
        }
        queryClient.invalidateQueries({
          queryKey: ["user/history-logs", user?.id],
        });
      };
    }
  });

  return (
    <SocketContext.Provider value={{ socket: socket }}>
      {children}
    </SocketContext.Provider>
  );
}
