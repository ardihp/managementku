import React, { useEffect, useState } from "react";
import { useAuth, UserDetail } from "@/hooks/auth/use-auth";
import { useWorkspace } from "@/hooks/use-workspace";
import { Workspace, WorkspaceByStatus } from "@/types/workspace";
import { Skeleton } from "@/components/ui/skeleton";
import FormWorkspace from "./components/form";
import { Archive, Plus } from "lucide-react";
import WorkspaceCard from "./components/card";
import ColorfulBadge from "@/components/layout/colorful-badge";

export default function WorkspacesView() {
  const [userDetail, setUserDetail] = useState<UserDetail | null>();
  const [selectedItem, setSelectedItem] = useState<Workspace | null>(null);

  const { user } = useAuth();
  const { getListWorkspace } = useWorkspace();

  const { data: listWorkspace, isLoading: loadingWorkspace } = getListWorkspace(
    {
      workspace_archived: "false",
      status_archived: "false",
    }
  );

  useEffect(() => {
    setUserDetail(user);
  }, []);

  return (
    <div className="flex flex-col h-full">
      <section className="flex items-center justify-between border-b border-zinc-200 p-8">
        <div className="h-9 flex flex-col">
          <p className="text-2xl text-foreground/40 font-light my-auto">
            Workspace{" "}
            <span className="text-foreground font-normal">/ All Workspace</span>
          </p>
        </div>

        <div className="flex">
          {userDetail?.role === "manager" && (
            <FormWorkspace
              selectedItem={selectedItem}
              resetSelectedItem={() => setSelectedItem(null)}
              buttonClassName="flex items-center gap-2 bg-primary py-2 px-3 rounded-md cursor-pointer h-9"
            >
              <Plus className="size-4 text-white" />
              <p className="text-white text-sm leading-none">Add New</p>
            </FormWorkspace>
          )}
        </div>
      </section>

      {loadingWorkspace ? (
        <section className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 p-8 gap-6">
          {[...Array(5)].map((item, key) => (
            <Skeleton key={key} className="w-full h-[177px]" />
          ))}
        </section>
      ) : listWorkspace?.length >= 1 ? (
        <section className="flex p-8 gap-6 overflow-auto h-full">
          {listWorkspace?.map((status: WorkspaceByStatus, key: number) => (
            <div className="flex flex-col gap-4 min-w-[300px]" key={key}>
              <div className="flex items-center justify-between gap-2 rounded-lg bg-white border opacity-100 py-2 px-4">
                <p className="font-medium">{status.name}</p>

                <ColorfulBadge
                  color={status.color}
                  text={status.workspaces.length.toString()}
                />
              </div>

              <div className="flex flex-col gap-4">
                {status.workspaces.map((workspace: Workspace, cardKey) => (
                  <WorkspaceCard
                    key={cardKey}
                    workspace={workspace}
                    setSelectedItem={setSelectedItem}
                    isArchived={workspace.is_archived}
                  />
                ))}

                {userDetail?.role === "manager" && (
                  <FormWorkspace
                    selectedItem={selectedItem}
                    resetSelectedItem={() => setSelectedItem(null)}
                    defaultStatus={status.id.toString()}
                    buttonClassName="flex items-center gap-2 rounded-lg bg-white border opacity-100 py-2 px-4 cursor-pointer hover:bg-zinc-50 transition-colors"
                  >
                    <Plus className="size-4" />
                    <p className="text-sm">Add New</p>
                  </FormWorkspace>
                )}
              </div>
            </div>
          ))}
        </section>
      ) : (
        <div className="flex flex-col h-full gap-2 items-center justify-center">
          <Archive className="size-10 stroke-[1.5px] text-zinc-700" />

          <p className="text-zinc-700">There's no workspace available</p>
        </div>
      )}
    </div>
  );
}
