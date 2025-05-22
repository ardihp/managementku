import React from "react";
import LayoutPage from "@/components/layout";
import { Plus } from "lucide-react";
import { useAuth } from "@/hooks/auth/use-auth";
import { useWorkspace } from "@/hooks/use-workspace";
import { Workspace } from "@/types/workspace";

export default function WorkspacesView() {
  const { user } = useAuth();
  const { getListWorkspace } = useWorkspace();
  const { data: listWorkspace, isLoading: loadingWorkspace } =
    getListWorkspace();

  return (
    <LayoutPage>
      <div className="flex flex-col">
        <section className="flex items-center justify-between p-8 border-b border-zinc-200">
          <p className="text-2xl text-foreground/40 font-light">
            Workspace <span className="text-foreground font-normal">/ All Workspace</span>
          </p>

          <div className="flex">
            {user?.role === "manager" && (
              <button className="flex items-center gap-2 bg-primary py-2 px-3 rounded-md cursor-pointer">
                <Plus className="size-4 text-white" />
                <p className="text-white text-sm">Add New</p>
              </button>
            )}
          </div>
        </section>

        <section className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 xl:grid-cols-4 p-8">
          {listWorkspace?.data?.map((workspace: Workspace, key: number) => (
            <div key={key} className="w-full h-10 border rounded-md bg-white"></div>
          ))}
        </section>
      </div>
    </LayoutPage>
  );
}
