import { HistoryLog } from "@/types/history-logs";
import { ColumnDef } from "@tanstack/react-table";
import dayjs from "dayjs";

export const historyColumns: ColumnDef<HistoryLog>[] = [
  {
    accessorKey: "created_at",
    header: "Created At",
    cell: ({ row }) => {
      return (
        <p className="">
          {dayjs(row.original.created_at).format("DD/MM/YYYY, HH:mm:ss")}
        </p>
      );
    },
  },
  {
    accessorKey: "user",
    header: "User",
  },
  {
    accessorKey: "description",
    header: "Description",
    cell: ({ row }) => {
      return (
        <p className="text-wrap">{row.original.description}</p>
      );
    },
  },
];
