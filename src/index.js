/**
 * Cloudflare Worker for mergit PDF merger
 * Serves static files (HTML and bundled JS)
 */

import htmlContent from '../public/index.html';
import pdfLibJs from '../public/pdf-lib.min.js';

export default {
  async fetch(request, env, ctx) {
    const url = new URL(request.url);
    
    // Serve pdf-lib.min.js
    if (url.pathname === '/pdf-lib.min.js') {
      return new Response(pdfLibJs, {
        headers: {
          'Content-Type': 'application/javascript',
          'Cache-Control': 'public, max-age=31536000, immutable',
          'X-Content-Type-Options': 'nosniff',
        },
      });
    }
    
    // Serve the HTML for all other requests
    return new Response(htmlContent, {
      headers: {
        'Content-Type': 'text/html;charset=UTF-8',
        'Cache-Control': 'public, max-age=3600',
        // Security headers
        'X-Content-Type-Options': 'nosniff',
        'X-Frame-Options': 'DENY',
        'X-XSS-Protection': '1; mode=block',
        'Referrer-Policy': 'strict-origin-when-cross-origin',
        // CSP - pdf-lib is now bundled locally
        'Content-Security-Policy': "default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline'; connect-src 'self'; img-src 'self' data:;",
      },
    });
  },
};
