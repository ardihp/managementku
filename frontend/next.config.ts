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
        destination: paths.dashboard,
        permanent: true,
      },
      {
        source: "/auth",
        destination: paths.auth.sign_in,
        permanent: true,
      },
    ];
  },
};

export default nextConfig;
