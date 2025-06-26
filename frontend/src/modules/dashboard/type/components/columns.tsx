"use client";

import { ColumnDef } from "@tanstack/react-table";

export interface WorkspaceType {
  id: number;
  name: string;
  color: string;
}

export const columns: ColumnDef<WorkspaceType>[] = [
  {
    accessorKey: "name",
    header: "Name",
  },
  {
    accessorKey: "color",
    header: "Color",
  },
];
