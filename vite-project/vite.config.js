import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'


// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  server: {
    host: true,
    port: 8080, // This is the port which we will use in docker
    proxy: {
      // 請參考 https://vitejs.dev/config/server-options.html
      '/room-list': {
        target: 'http://127.0.0.1:8080',
        changeOrigin: true,
      },
      '/restaurant-list': {
        target: 'http://127.0.0.1:8080',
        changeOrigin: true,
      },
      '/login': {
        target: 'http://127.0.0.1:8080',
        changeOrigin: true,
      }
    },
  },
})


