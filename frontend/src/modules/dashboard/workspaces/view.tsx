import React, { useEffect, useState } from "react";
import LayoutPage from "@/components/layout";
import { useAuth, UserDetail } from "@/hooks/auth/use-auth";
import { useWorkspace } from "@/hooks/use-workspace";
import { Workspace } from "@/types/workspace";
import { Skeleton } from "@/components/ui/skeleton";
import FormWorkspace from "./components/form";
import { Archive } from "lucide-react";
import WorkspaceCard from "./components/card";

export default function WorkspacesView() {
  const [userDetail, setUserDetail] = useState<UserDetail | null>();
  const [selectedItem, setSelectedItem] = useState<Workspace | null>(null);

  const { user } = useAuth();
  const { getListWorkspace } = useWorkspace();

  const { data: listWorkspace, isLoading: loadingWorkspace } = getListWorkspace(
    { is_archived: false }
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
            />
          )}
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
            <WorkspaceCard
              key={key}
              workspace={workspace}
              setSelectedItem={setSelectedItem}
              isArchived={false}
            />
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
