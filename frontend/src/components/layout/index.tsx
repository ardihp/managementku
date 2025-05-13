import React, { ReactNode } from "react";
import Sidebar from "./sidebar";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import AuthProvider from "@/context/auth";

export default function LayoutPage({ children }: { children: ReactNode }) {
  const queryClient = new QueryClient();

  return (
    <AuthProvider>
      <QueryClientProvider client={queryClient}>
        <section className="grid lg:grid-cols-[minmax(0,300px)_minmax(400px,1fr)] h-screen">
          <Sidebar />
          <div className="p-4">{children}</div>
        </section>
      </QueryClientProvider>
    </AuthProvider>
  );
}
