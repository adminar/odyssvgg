#!/bin/bash

# 创建项目文件夹
PROJECT_NAME="odyssey"
mkdir $PROJECT_NAME && cd $PROJECT_NAME

# 初始化 Vite + Vue3 + TS 项目
pnpm create vite . --template vue-ts

# 安装依赖
pnpm add vue-router@4 pinia element-plus

# 安装开发依赖
pnpm add -D sass vite-plugin-svg-icons @vitejs/plugin-vue typescript

# 创建目录结构
mkdir -p src/{views,components,router,store,assets/icons,style,api,utils,composables}

# 创建基础文件
echo 'import { createRouter, createWebHistory } from "vue-router";

const routes = [
  {
    path: "/",
    name: "Home",
    component: () => import("@/views/Home.vue")
  }
];

export default createRouter({
  history: createWebHistory(),
  routes
});
' > src/router/index.ts

echo 'import { defineStore } from "pinia";

export const useMainStore = defineStore("main", {
  state: () => ({
    title: "Hello Element Plus"
  })
});
' > src/store/index.ts

echo '<template>
  <el-container style="height: 100vh">
    <el-header>Header</el-header>
    <el-main>Main Content</el-main>
  </el-container>
</template>

<script setup lang="ts"></script>
' > src/views/Home.vue

echo 'import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import { createPinia } from "pinia";

import ElementPlus from "element-plus";
import "element-plus/dist/index.css";

import "./style/index.scss";

const app = createApp(App);

app.use(createPinia());
app.use(router);
app.use(ElementPlus);

app.mount("#app");
' > src/main.ts

echo '<script setup lang="ts">
</script>

<template>
  <router-view />
</template>
' > src/App.vue

echo '@import "./reset.scss";' > src/style/index.scss
touch src/style/reset.scss src/style/variables.scss

# 修改 vite.config.ts 添加 @ 路径别名
sed -i '' 's|resolve: {|resolve: {\n    alias: { "@": path.resolve(__dirname, "src") },|' vite.config.ts

# 初始化完成提示
echo "✅ Vue 3 + Element Plus 项目初始化完成！"
echo "👉 运行以下命令启动项目："
echo "cd $PROJECT_NAME && pnpm install && pnpm dev"
