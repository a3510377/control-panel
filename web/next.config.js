/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,

  async rewrites() {
    // Rewrite all API routes to the API server
    return [
      {
        source: '/api',
        destination: 'http://localhost:8000/api',
      },
    ];
  },
  future: { webpack5: true },
  trailingSlash: true,
  swcMinify: true,
  experimental: {
    appDir: true,
  },
  images: {
    unoptimized: true, // FIXME
  },
};

module.exports = nextConfig;
