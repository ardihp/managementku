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
import { useWorkspace } from "@/hooks/use-workspace";
import { useWorkspaceCard } from "@/hooks/use-workspace-card";
import { cn } from "@/lib/utils";
import { Member } from "@/types/workspace";
import { WorkspaceCard } from "@/types/workspace-card";
import { Check } from "lucide-react";
import React, { useContext, useEffect, useState } from "react";

interface AssignMemberProps {
  open: boolean;
  onOpenChange: () => void;
  workspaceId?: number;
  card: WorkspaceCard;
}

export default function AssignMember({
  open,
  onOpenChange,
  workspaceId = 0,
  card,
}: AssignMemberProps) {
  const [search, setSearch] = useState("");
  const [assignedMember, setAssignedMember] = useState<number[]>([]);
  const { socket } = useContext(SocketContext);

  const { getListWorkspaceMember } = useWorkspace();
  const { assignWorkspaceCardMember } = useWorkspaceCard();

  const { data: workspaceMembers, isLoading: loadingMember } =
    getListWorkspaceMember(workspaceId);
  const { mutateAsync: assign, isPending: loadingAssign } =
    assignWorkspaceCardMember(workspaceId, card.id);

  useEffect(() => {
    if (open && !loadingMember) {
      setAssignedMember(card.members.map((member) => member.id));
    }
  }, [open, loadingMember]);

  const handleAssignedMember = (id: number) => {
    let members = Array.from(assignedMember);

    if (members.includes(id)) {
      members = members.filter((member_id) => member_id !== id);
    } else {
      members.push(id);
    }

    setAssignedMember([...members]);
  };

  const handleAssign = async () => {
    try {
      await assign(
        {
          card_id: card.id,
          user_id: assignedMember,
        },
        {
          onSuccess: () => {
            onOpenChange();
          },
        }
      );
    } catch (error) {
      console.error(error);
    }

    onOpenChange();
    socket?.send(
      JSON.stringify({
        type: "detail-card",
        workspace_id: workspaceId,
        card_id: card.id,
      })
    );
  };

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Assign Member Workspace Card</DialogTitle>
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
              {loadingMember
                ? [...new Array(4)].map((item, key) => (
                    <Skeleton key={key} className="w-full h-[51px]" />
                  ))
                : workspaceMembers
                    .filter((member: Member) =>
                      member.name.toLowerCase().includes(search.toLowerCase())
                    )
                    .map((member: Member, key: number) => (
                      <div
                        key={key}
                        className="flex items-center gap-4 border rounded-lg p-2 cursor-pointer hover:bg-zinc-50/50"
                        onClick={() => handleAssignedMember(member.id)}
                      >
                        <Avatar className="size-10">
                          <AvatarFallback className="uppercase">
                            {member.username[0]}
                          </AvatarFallback>
                        </Avatar>

                        <p className="font-quicksand capitalize">
                          {member.username}
                        </p>

                        {assignedMember.includes(member.id) && (
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
