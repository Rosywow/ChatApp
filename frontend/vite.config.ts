import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'
import { fileURLToPath, URL } from 'node:url'
import sveltePreprocess from 'svelte-preprocess';
import { less } from 'svelte-preprocess-less';


// https://vitejs.dev/config/
export default defineConfig({
  plugins: [svelte({
    preprocess: {
      style: less(),
    }
  })
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  server: {
    host: 'localhost',
    port: 8082,
    open: false,
    https: false,
    proxy: {
      '/api': 'http://localhost:9090/',
    },
  },
})
