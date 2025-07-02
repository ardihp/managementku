import { colorsCode } from "@/const";
import { SocketContext } from "@/context/socket";
import { useWorkspace } from "@/hooks/use-workspace";
import { cn } from "@/lib/utils";
import { WorkspaceByStatus } from "@/types/workspace";
import { FolderOpen } from "lucide-react";
import React, { useContext, useEffect } from "react";

export default function DashboardView() {
  const { getListWorkspace } = useWorkspace();
  const { socket } = useContext(SocketContext);

  useEffect(() => {
    if (socket) {
      socket.onopen = () => {
        console.log("connected");
      };

      socket.onmessage = (event) => {
        console.log("message received: ", JSON.parse(event.data || "{}"));
      };

      socket.onclose = () => {
        console.log("connection close");
      };
    }
  });

  const { data: workspaces, isLoading: ladingWorkspace } = getListWorkspace();

  return (
    <div className="flex flex-col h-full">
      <section className="flex items-center justify-between border-b border-zinc-200 p-8">
        <div
          className="h-9 flex flex-col"
          onClick={() =>
            socket?.send(JSON.stringify({ text: "haha", value: [1, 2] }))
          }
        >
          <p className="text-2xl text-foreground font-light my-auto">
            Dashboard Overview
          </p>
        </div>
      </section>

      <div className="grid grid-cols-[2fr_minmax(200px,350px)] gap-4 h-full p-8">
        <div className="flex flex-col gap-4">
          <div className="grid grid-cols-3 gap-4">
            {workspaces?.map((status: WorkspaceByStatus, key: number) => (
              <div
                key={key}
                className={cn(
                  "flex flex-col border rounded-xl p-4 gap-3 relative overflow-hidden h-32",
                  colorsCode.find((code) => code.color === status.color)
                    ?.background
                )}
              >
                <div className="absolute -bottom-6 -right-4">
                  <FolderOpen
                    className={cn(
                      "opacity-20 size-24",
                      colorsCode.find((code) => code.color === status.color)
                        ?.text
                    )}
                  />
                </div>

                <p className={cn("capitalize text-lg")}>{status.name}</p>
                <p
                  className={cn(
                    "font-bold text-3xl",
                    colorsCode.find((code) => code.color === status.color)?.text
                  )}
                >
                  {status.workspaces.length}
                </p>
              </div>
            ))}
          </div>
        </div>

        <div className="flex flex-col border rounded-xl p-4">
          <p className="text-xl">Recent Activity</p>
          <hr className="mt-3" />

          <div></div>
        </div>
      </div>
    </div>
  );
}
