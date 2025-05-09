import React, { ReactNode } from "react";
import Sidebar from "./sidebar";

export default function LayoutPage({ children }: { children: ReactNode }) {
  return (
    <section className="grid lg:grid-cols-[minmax(0,300px)_minmax(400px,1fr)] h-screen">
      <Sidebar />
      <div className="p-4">{children}</div>
    </section>
  );
}
