import { cn } from "@/lib/utils";
import React, { useEffect, useState } from "react";

interface ColorfulBadgeProps {
  text: string;
  color?:
    | "default"
    | "gray"
    | "brown"
    | "orange"
    | "yellow"
    | "green"
    | "blue"
    | "purple"
    | "pink"
    | "red"
    | string;
}

export default function ColorfulBadge({
  text,
  color = "default",
}: ColorfulBadgeProps) {
  const [colorset, setColorset] = useState({
    background: "",
    text: "",
  });

  useEffect(() => {
    switch (color) {
      case "gray":
        setColorset({ background: "bg-[#F1F1EF]", text: "text-[#787774]" });
        break;
      case "brown":
        setColorset({ background: "bg-[#F3EEEE]", text: "text-[#976D57]" });
        break;
      case "orange":
        setColorset({ background: "bg-[#F8ECDF]", text: "text-[#CC782F]" });
        break;
      case "yellow":
        setColorset({ background: "bg-[#FAF3DD]", text: "text-[#C29343]" });
        break;
      case "green":
        setColorset({ background: "bg-[#EEF3ED]", text: "text-[#548164]" });
        break;
      case "blue":
        setColorset({ background: "bg-[#E9F3F7]", text: "text-[#487CA5]" });
        break;
      case "purple":
        setColorset({ background: "bg-[#F6F3F8]", text: "text-[#8A67AB]" });
        break;
      case "pink":
        setColorset({ background: "bg-[#F9F2F5]", text: "text-[#B35488]" });
        break;
      case "red":
        setColorset({ background: "bg-[#FAECEC]", text: "text-[#C4554D]" });
        break;
      default:
        setColorset({ background: "bg-[#FFFFFF]", text: "text-[#373530]" });
        break;
    }
  }, [color]);

  return (
    <div
      className={cn(
        "flex items-center justify-center py-2 px-[10px] w-fit rounded-md",
        colorset.background
      )}
    >
      <p
        className={cn(
          "text-[10px] font-quicksand uppercase font-bold",
          colorset.text
        )}
      >
        {text}
      </p>
    </div>
  );
}
