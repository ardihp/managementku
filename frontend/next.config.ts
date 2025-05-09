import { paths } from "@/const/routes";
import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  /* config options here */
  reactStrictMode: true,
  async redirects() {
    return [
      // Basic redirect
      {
        source: "/",
        destination: paths.dashboard.root,
        permanent: true,
      },
    ];
  },
};

export default nextConfig;
