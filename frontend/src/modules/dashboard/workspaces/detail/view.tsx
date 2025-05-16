import React from "react";
import LayoutPage from "@/components/layout";
import { useRouter } from "next/router";

export default function DetailWorkspaceView() {
  const router = useRouter();

  return (
    <LayoutPage>
      <p>Workspace {router.query.id} Page</p>
    </LayoutPage>
  );
}
