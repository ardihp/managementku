import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import { Skeleton } from "@/components/ui/skeleton";
import { useUser } from "@/hooks/use-user";
import { HistoryLog } from "@/types/history-logs";
import React from "react";

export default function RecentActivity() {
  const { getHistoryLog } = useUser();
  const { data: logs, isLoading: loadingLog } = getHistoryLog("general");

  return (
    <div className="flex flex-col border rounded-xl overflow-hidden pt-4">
      <p className="text-xl px-4">Recent Activity</p>

      <hr className="mt-3" />

      <div className="flex flex-col h-full overflow-auto max-h-[calc(100vh_-_253px)]">
        {loadingLog ? (
          [...new Array(5)].map((item, key) => (
            <div
              key={key}
              className="flex gap-4 items-center first:border-0 border-t p-4"
            >
              <Skeleton className="rounded-full size-10" />

              <div className="flex flex-col gap-2 w-full">
                <Skeleton className="w-32 h-4" />
                <Skeleton className="w-full h-4" />
              </div>
            </div>
          ))
        ) : logs.length >= 1 ? (
          logs.map((log: HistoryLog, key: number) => (
            <div
              key={key}
              className="flex gap-4 items-center first:border-0 border-t p-4 bg-transparent hover:bg-zinc-300/50 duration-200"
            >
              <Avatar className="size-10">
                <AvatarFallback className="capitalize">
                  {log.user?.[0]}
                </AvatarFallback>
              </Avatar>

              <div className="flex flex-col">
                <p className="font-medium capitalize truncate">{log.user}</p>
                <p className="text-xs line-clamp-2 opacity-80 font-light">
                  {log.description}
                </p>
              </div>
            </div>
          ))
        ) : (
          <div className="flex flex-col h-full items-center justify-center">
            <p className="text-sm">No recent activity.</p>
          </div>
        )}
      </div>
    </div>
  );
}
