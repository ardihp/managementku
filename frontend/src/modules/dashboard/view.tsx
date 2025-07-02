import { colorsCode } from "@/const";
import { useWorkspace } from "@/hooks/use-workspace";
import { cn } from "@/lib/utils";
import { WorkspaceByStatus } from "@/types/workspace";
import { Folder, FolderOpen } from "lucide-react";
import React from "react";

export default function DashboardView() {
  const { getListWorkspace } = useWorkspace();

  const { data: workspaces, isLoading: ladingWorkspace } = getListWorkspace();

  return (
    <div className="flex flex-col h-full">
      <section className="flex items-center justify-between border-b border-zinc-200 p-8">
        <div className="h-9 flex flex-col">
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
          <p className="text-2xl">Recent Activity</p>
          <hr className="mt-3" />

          <div></div>
        </div>
      </div>
    </div>
  );
}
