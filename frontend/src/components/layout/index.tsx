import React, { ReactNode } from "react";
import Sidebar from "./sidebar";
import AuthProvider from "@/context/auth";
import Header from "./header";

export default function LayoutPage({ children }: { children: ReactNode }) {
  return (
    <AuthProvider>
      <section className="grid lg:grid-cols-[minmax(0,300px)_minmax(400px,1fr)] h-screen bg-zinc-100/30">
        <Sidebar />
        <div className="flex flex-col">
          <Header />

          {children}
        </div>
      </section>
    </AuthProvider>
  );
}
