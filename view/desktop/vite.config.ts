import { resolve } from "path";
import { defineConfig } from "vite";

import tailwindcss from "@tailwindcss/vite";
import react from "@vitejs/plugin-react";

export default defineConfig({
  server: {
    watch: {
      ignored: ["**/*.spec.ts", "**/*.test.ts", "storage/**"],
    },
  },
  plugins: [tailwindcss(), react()],
  assetsInclude: "src/renderer/assets/**", //TODO what's this assetsInclude for???
  resolve: {
    alias: {
      "@": resolve("src"),
      "@/hooks": resolve("src/hooks"),
      "@/assets": resolve("src/assets"),
      "@/components": resolve("src/components"),
    },
  },
});
