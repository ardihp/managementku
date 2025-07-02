import { useAuth } from "@/hooks/auth/use-auth";
import { createContext, ReactNode, useEffect, useState } from "react";

export const SocketContext = createContext<{ socket: WebSocket | null }>({
  socket: null,
});

export default function SocketProvider({ children }: { children: ReactNode }) {
  const [socket, setSocket] = useState<WebSocket | null>(null);
  const { user } = useAuth();

  useEffect(() => {
    const ws = new WebSocket("ws://localhost:8000/ws");

    if (user) {
      setSocket(ws);
    }

    return () => {
      ws.close();
    };
  }, []);

  return (
    <SocketContext.Provider value={{ socket: socket }}>
      {children}
    </SocketContext.Provider>
  );
}
