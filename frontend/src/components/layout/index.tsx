import React, { ReactNode } from "react";
import Sidebar from "./sidebar";
import AuthProvider from "@/context/auth";
import Header from "./header";
import { useRouter } from "next/router";
import SocketProvider from "@/context/socket";

export default function LayoutPage({ children }: { children: ReactNode }) {
  const router = useRouter();

  return (
    <SocketProvider>
      <AuthProvider>
        {router.pathname.includes("auth") ? (
          children
        ) : (
          <section className="grid lg:grid-cols-[minmax(0,300px)_minmax(400px,1fr)] h-screen bg-zinc-100/30">
            <Sidebar />
            <div className="flex flex-col">
              <Header />

              {children}
            </div>
          </section>
        )}
      </AuthProvider>
    </SocketProvider>
  );
}
