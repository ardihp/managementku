"use client";

import React from "react";
import LayoutPage from "@/components/layout";
import { useWorkspace } from "@/hooks/use-workspace";
import { useParams } from "next/navigation";
import { Loader2, Plus } from "lucide-react";
import Link from "next/link";
import { paths } from "@/const/routes";
import { Skeleton } from "@/components/ui/skeleton";
import ColorfulBadge from "@/components/layout/colorful-badge";
import { WorkspaceCardByStatus } from "@/types/workspace";

export default function DetailWorkspaceView() {
  const params = useParams();
  const { getDetailWorkspace, getListWorkspaceCardStatus } = useWorkspace();
  const { data: detailWorkspace, isLoading: loadingWorkspace } =
    getDetailWorkspace(Number(params?.id || 0));
  const { data: workspaceStatus } = getListWorkspaceCardStatus(
    Number(params?.id || 0),
    "false"
  );

  return (
    <div className="flex flex-col h-full">
      <section className="flex items-center justify-between border-b border-zinc-200 p-8">
        <div className="h-9 flex items-center gap-1.5">
          <Link href={paths.dashboard.workspace}>
            <p className="text-2xl text-foreground/40 font-light my-auto">
              Workspace{" "}
            </p>
          </Link>

          <div className="text-2xl text-foreground my-auto flex items-center gap-1">
            /{" "}
            {loadingWorkspace ? (
              <Skeleton className="w-60 h-7" />
            ) : (
              <span>{detailWorkspace?.name}</span>
            )}
          </div>
        </div>
      </section>

      <section className="flex flex-col h-full gap-6">
        {loadingWorkspace ? (
          <div className="flex flex-col h-full gap-2 items-center justify-center">
            <Loader2 className="size-10 stroke-[1.5px] text-zinc-700 animate-spin" />

            <p className="text-zinc-700">Loading get workspace data...</p>
          </div>
        ) : (
          <div className="flex gap-5 overflow-auto p-8 h-full">
            {workspaceStatus?.map(
              (status: WorkspaceCardByStatus, key: number) => (
                <div key={key} className="flex flex-col min-w-[300px] gap-4">
                  <div className="flex items-center justify-between gap-2 rounded-lg bg-white border opacity-100 py-2 px-4">
                    <p className="font-medium">{status.name}</p>

                    <ColorfulBadge color={status.color} text="0" />
                  </div>

                  <div className="flex flex-col gap-4">
                    {status.cards.map((card, cardKey: number) => (
                      <div key={cardKey}>{card.title}</div>
                    ))}

                    <div className="flex items-center gap-2 rounded-lg bg-white border opacity-100 py-2 px-4 cursor-pointer hover:bg-zinc-50 transition-colors">
                      <Plus className="size-4" />
                      <p className="font-medium">Add a card</p>
                    </div>
                  </div>
                </div>
              )
            )}
          </div>
        )}
      </section>
    </div>
  );
}
