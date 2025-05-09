import React from "react";
import LayoutPage from "@/components/layout";
import { useRouter } from "next/router";

export default function ProjectView() {
  const router = useRouter();

  return (
    <LayoutPage>
      <p>Project Page</p>

      <p>{router.query.id}</p>
    </LayoutPage>
  );
}
