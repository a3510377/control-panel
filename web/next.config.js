const { PHASE_DEVELOPMENT_SERVER } = require('next/constants');

module.exports = (phase) => {
  const isDev = phase == PHASE_DEVELOPMENT_SERVER;
  /** @type {import('next').NextConfig} */
  const nextConfig = {
    reactStrictMode: true,
    future: { webpack5: true },
    trailingSlash: true,
    swcMinify: true,
    experimental: {
      appDir: true,
    },
    images: {
      unoptimized: true, // FIXME
    },
    env: {
      API_BASE_URL: isDev ? 'http://localhost:8000' : '',
    },
  };

  return nextConfig;
};
