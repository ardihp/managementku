import React, { useEffect, useState } from "react";
import LayoutPage from "@/components/layout";
import { useAuth, UserDetail } from "@/hooks/auth/use-auth";
import { useWorkspace } from "@/hooks/use-workspace";
import { Workspace } from "@/types/workspace";
import ColorfulBadge from "@/components/layout/colorful-badge";
import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import Link from "next/link";
import { paths } from "@/const/routes";
import { Skeleton } from "@/components/ui/skeleton";
import FormWorkspace from "./components/form";
import WorkspaceAction from "./components/action";

export default function WorkspacesView() {
  const [userDetail, setUserDetail] = useState<UserDetail | null>();
  const [selectedItem, setSelectedItem] = useState<Workspace | null>(null);

  const { user } = useAuth();
  const { getListWorkspace } = useWorkspace();

  const { data: listWorkspace, isLoading: loadingWorkspace } =
    getListWorkspace();

  useEffect(() => {
    setUserDetail(user);
  }, []);

  return (
    <LayoutPage>
      <div className="flex flex-col">
        <section className="flex items-center justify-between border-b border-zinc-200 p-8">
          <div className="h-9 flex flex-col">
            <p className="text-2xl text-foreground/40 font-light my-auto">
              Workspace{" "}
              <span className="text-foreground font-normal">
                / All Workspace
              </span>
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

        <section className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 p-8 gap-6">
          {loadingWorkspace
            ? [...Array(5)].map((item, key) => (
                <Skeleton key={key} className="w-full h-[177px]" />
              ))
            : listWorkspace?.data?.map((workspace: Workspace, key: number) => (
                <div key={key} className="relative">
                  <WorkspaceAction
                    setSelectedItem={setSelectedItem}
                    workspace={workspace}
                  />

                  <Link href={`${paths.dashboard.workspace}/${workspace.id}`}>
                    <div className="w-full border rounded-xl bg-white flex flex-col p-4 gap-2 min-h-[177px]">
                      <div className="flex items-center gap-2">
                        <ColorfulBadge
                          text={workspace.type.name}
                          color={workspace.type.color}
                        />

                        {workspace.tags.map((tag, key) => (
                          <ColorfulBadge key={key} text={tag} color="gray" />
                        ))}
                      </div>

                      <p
                        className="font-medium text-lg capitalize leading-[25px] line-clamp-2"
                        title={workspace.name}
                      >
                        {workspace.name}
                      </p>

                      <p
                        className="font-quicksand text-xs font-medium text-zinc-400 capitalize leading-[18px] line-clamp-2"
                        title={workspace.description}
                      >
                        {workspace.description}
                      </p>

                      <div className="*:data-[slot=avatar]:ring-zinc-200 *:data-[slot=avatar]:bg-zinc-100 flex -space-x-2 *:data-[slot=avatar]:ring-2 mt-3">
                        {workspace.members.map((member, key) => (
                          <Avatar
                            key={key}
                            title={member.username}
                            className="size-8"
                          >
                            <AvatarFallback className="uppercase text-sm">
                              {member.username
                                ?.split(" ")
                                ?.map((text) => text?.[0])}
                            </AvatarFallback>
                          </Avatar>
                        ))}
                      </div>
                    </div>
                  </Link>
                </div>
              ))}
        </section>
      </div>
    </LayoutPage>
  );
}
