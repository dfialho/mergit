/**
 * Cloudflare Worker for mergit PDF merger
 * Serves only the index.html file
 */

import htmlContent from '../public/index.html';

export default {
  async fetch(request, env, ctx) {
    // Serve the HTML for all requests (only index.html is deployed)
    return new Response(htmlContent, {
      headers: {
        'Content-Type': 'text/html;charset=UTF-8',
        'Cache-Control': 'public, max-age=3600',
        // Security headers
        'X-Content-Type-Options': 'nosniff',
        'X-Frame-Options': 'DENY',
        'X-XSS-Protection': '1; mode=block',
        'Referrer-Policy': 'strict-origin-when-cross-origin',
        // CSP to allow pdf-lib from unpkg CDN
        'Content-Security-Policy': "default-src 'self'; script-src 'self' 'unsafe-inline' https://unpkg.com; style-src 'self' 'unsafe-inline'; connect-src 'self'; img-src 'self' data:;",
      },
    });
  },
};
