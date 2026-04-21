import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  transpilePackages: ["@workspace/mocks", "@workspace/ui"],
};

export default nextConfig;
