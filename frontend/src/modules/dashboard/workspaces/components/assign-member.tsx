import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Skeleton } from "@/components/ui/skeleton";
import { SocketContext } from "@/context/socket";
import { useUser } from "@/hooks/use-user";
import { useWorkspace } from "@/hooks/use-workspace";
import { cn } from "@/lib/utils";
import { Member, Workspace, WorkspaceRole } from "@/types/workspace";
import { Check } from "lucide-react";
import React, { useContext, useEffect, useState } from "react";

interface AssignMemberProps {
  open: boolean;
  onOpenChange: () => void;
  workspace: Workspace;
}

export default function AssignMember({
  open,
  onOpenChange,
  workspace,
}: AssignMemberProps) {
  const [search, setSearch] = useState("");
  const [assignedMember, setAssignedMember] = useState<number[]>([]);
  const { socket } = useContext(SocketContext);

  const { getListEmployee } = useUser();
  const { assignWorkspaceMember } = useWorkspace();

  const { data: employees, isLoading: loadingEmployee } = getListEmployee();
  const { mutate: assign, isPending: loadingAssign } = assignWorkspaceMember();

  useEffect(() => {
    if (!loadingEmployee) {
      setAssignedMember(workspace.members.map((member) => member.id));
    }
  }, [loadingEmployee]);

  const handleAssignedMember = (id: number) => {
    let members = Array.from(assignedMember);

    if (members.includes(id)) {
      members = members.filter((member_id) => member_id !== id);
    } else {
      members.push(id);
    }

    setAssignedMember([...members]);
  };

  const handleAssign = () => {
    assign(
      {
        workspace_id: workspace.id,
        user_id: assignedMember,
      },
      {
        onSuccess: () => {
          onOpenChange();
        },
      }
    );
    socket?.send(
      JSON.stringify({ type: "assign-workspace", workspace_id: workspace.id })
    );
  };

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Assign Member Workspace</DialogTitle>
          <DialogDescription></DialogDescription>
        </DialogHeader>

        <div className="flex flex-col gap-6">
          <Input
            value={search}
            onChange={(e) => setSearch(e.target.value)}
            placeholder="Search by username"
          />

          <div className="flex flex-col gap-3">
            <p className="text-sm">Members</p>

            <div className="flex flex-col gap-2 max-h-[300px] overflow-auto">
              {loadingEmployee
                ? [...new Array(4)].map((item, key) => (
                    <Skeleton key={key} className="w-full h-[51px]" />
                  ))
                : employees
                    .filter((employee: Member) =>
                      employee.name.toLowerCase().includes(search.toLowerCase())
                    )
                    .map((employee: Member, key: number) => (
                      <div
                        key={key}
                        className="flex items-center gap-4 border rounded-lg p-2 cursor-pointer hover:bg-zinc-50/50"
                        onClick={() => handleAssignedMember(employee.id)}
                      >
                        <Avatar className="size-10">
                          <AvatarFallback className="uppercase">
                            {employee.username[0]}
                          </AvatarFallback>
                        </Avatar>

                        <p className="font-quicksand capitalize">
                          {employee.username}
                        </p>

                        {assignedMember.includes(employee.id) && (
                          <div className="ml-auto mr-2">
                            <Check className="size-5" />
                          </div>
                        )}
                      </div>
                    ))}
            </div>
          </div>
        </div>

        <button
          className={cn(
            "border border-zinc-200 h-10 bg-zinc-700 rounded-lg shadow w-full cursor-pointer duration-200",
            loadingAssign && "cursor-not-allowed opacity-60"
          )}
          onClick={handleAssign}
          disabled={loadingAssign}
        >
          <p className="text-white text-sm">Save</p>
        </button>
      </DialogContent>
    </Dialog>
  );
}
