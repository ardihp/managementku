import React, { ReactNode } from "react";
import AuthProvider from "@/context/auth";
import Header from "./header";
import { useRouter } from "next/router";
import SocketProvider from "@/context/socket";
import { SidebarProvider, SidebarTrigger } from "../ui/sidebar";
import AppSidebar from "./sidebar";

export default function LayoutPage({ children }: { children: ReactNode }) {
  const router = useRouter();

  return (
    <SocketProvider>
      <AuthProvider>
        {router.pathname.includes("auth") ? (
          children
        ) : (
          <>
            <SidebarProvider>
              <AppSidebar />
              <main className="flex flex-col bg-zinc-100/30 overflow-auto w-full">
                <Header />

                {children}
              </main>
            </SidebarProvider>
          </>
        )}
      </AuthProvider>
    </SocketProvider>
  );
}
