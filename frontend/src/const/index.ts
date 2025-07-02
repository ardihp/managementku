export const colors = [
  "default",
  "gray",
  "brown",
  "orange",
  "yellow",
  "green",
  "blue",
  "purple",
  "pink",
  "red",
] as const;
export type colorsType = (typeof colors)[number];

export const colorsCode = [
  {
    color: "gray",
    background: "bg-[#F1F1EF]",
    text: "text-[#787774]",
  },
  {
    color: "brown",
    background: "bg-[#F3EEEE]",
    text: "text-[#976D57]",
  },
  {
    color: "orange",
    background: "bg-[#F8ECDF]",
    text: "text-[#CC782F]",
  },
  {
    color: "yellow",
    background: "bg-[#FAF3DD]",
    text: "text-[#C29343]",
  },
  {
    color: "green",
    background: "bg-[#EEF3ED]",
    text: "text-[#548164]",
  },
  {
    color: "blue",
    background: "bg-[#E9F3F7]",
    text: "text-[#487CA5]",
  },
  {
    color: "purple",
    background: "bg-[#F6F3F8]",
    text: "text-[#8A67AB]",
  },
  {
    color: "pink",
    background: "bg-[#F9F2F5]",
    text: "text-[#B35488]",
  },
  {
    color: "red",
    background: "bg-[#FAECEC]",
    text: "text-[#C4554D]",
  },
  {
    color: "default",
    background: "bg-[#FFFFFF]",
    text: "text-[#373530]",
  },
];
