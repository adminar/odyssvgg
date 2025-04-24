import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd())

  return {
    plugins: [vue()],
    resolve: {
      alias: {
        '@': path.resolve(__dirname, 'src')
      }
    },
    server: {
      host: '0.0.0.0',
      port: 8001,
      open: true,
      proxy: {
        '/api': {
          target: env.VITE_APP_BASE_URL,
          changeOrigin: true,
          rewrite: path => path.replace(/^\/api/, '')
        }
      }
    },
    base: env.VITE_PUBLIC_PATH || '/',
  }
})
