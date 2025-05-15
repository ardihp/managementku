import Image from "next/image";
import React, { ReactNode } from "react";

export default function AuthLayout({ children }: { children: ReactNode }) {
  return (
    <section className="flex flex-col-reverse xl:flex-row h-screen">
      <div className="w-full h-full flex xl:items-center justify-center">
        {children}
      </div>

      <div className="relative w-full h-full max-h-[250px] xl:max-h-full border-l border-zinc-200">
        <div className="absolute w-full h-full bg-gradient-to-t from-white xl:from-zinc-300/10 from-30% to-transparent xl:to-zinc-300/10 z-[1]" />
        <div className="relative w-full h-full">
          <Image
            src="/assets/kanban-example.png"
            alt="Auth Illustration"
            fill
            sizes="800px"
            className="object-cover object-center xl:object-left"
          />
        </div>
      </div>
    </section>
  );
}
