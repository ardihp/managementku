import LayoutPage from "@/components/layout";
import { Skeleton } from "@/components/ui/skeleton";
import { useWorkspace } from "@/hooks/use-workspace";
import React from "react";
import WorkspaceCard from "../workspaces/components/card";
import { Workspace } from "@/types/workspace";
import { Archive } from "lucide-react";

export default function ArchiveViewF() {
  const { getListWorkspace } = useWorkspace();
  const { data: listWorkspace, isLoading: loadingWorkspace } = getListWorkspace(
    { is_archived: true }
  );

  return (
    <div className="flex flex-col h-full">
      <section className="flex items-center justify-between border-b border-zinc-200 p-8">
        <div className="h-9 flex flex-col">
          <p className="text-2xl text-foreground/40 font-light my-auto">
            Workspace{" "}
            <span className="text-foreground font-normal">/ Archived</span>
          </p>
        </div>
      </section>

      {loadingWorkspace ? (
        <section className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 p-8 gap-6">
          {[...Array(5)].map((item, key) => (
            <Skeleton key={key} className="w-full h-[177px]" />
          ))}
        </section>
      ) : listWorkspace?.data?.length >= 1 ? (
        <section className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 p-8 gap-6">
          {listWorkspace?.data?.map((workspace: Workspace, key: number) => (
            <WorkspaceCard key={key} workspace={workspace} isArchived={true} />
          ))}
        </section>
      ) : (
        <div className="flex flex-col h-full gap-2 items-center justify-center">
          <Archive className="size-10 stroke-[1.5px] text-zinc-700" />

          <p className="text-zinc-700">There's no workspace archived</p>
        </div>
      )}
    </div>
  );
}
