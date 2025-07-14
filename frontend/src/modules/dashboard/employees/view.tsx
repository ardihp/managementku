import React, { useState } from "react";
import DataTable from "@/components/layout/data-table";
import { columns } from "./components/columns";
import FormWorkspaceType from "./components/form";
import { Plus } from "lucide-react";
import { useUser } from "@/hooks/use-user";

export default function EmployeesView() {
  const [openDialog, setOpenDialog] = useState(false);

  const { getListUser } = useUser();
  const { data, isLoading } = getListUser();

  return (
    <>
      <FormWorkspaceType
        open={openDialog}
        onOpenChange={() => setOpenDialog(!openDialog)}
      />

      <div className="flex flex-col h-full">
        <section className="flex items-center justify-between border-b border-zinc-200 p-[28px]">
          <div className="h-9 flex flex-col">
            <p className="text-2xl text-foreground font-light my-auto">
              Employee List
            </p>
          </div>

          <button
            className="flex items-center gap-2 bg-primary py-2 px-3 rounded-md cursor-pointer h-9"
            onClick={() => setOpenDialog(true)}
          >
            <Plus className="size-4 text-white" />
            <p className="text-white text-sm leading-none">Add New</p>
          </button>
        </section>

        <section className="flex flex-col p-[28px]">
          <DataTable columns={columns} data={data} isLoading={isLoading} />
        </section>
      </div>
    </>
  );
}
