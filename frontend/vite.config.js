import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    react({
      // 禁用 Strict Mode 以避免一些不必要的警告
      strictMode: false,
    }),
  ],
  server: {
    port: 5173,
    host: true,
  },
});
