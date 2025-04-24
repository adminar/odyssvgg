<template>
  <div class="common-layout">
    <el-container>
      <el-aside width="200px">
        <collapse/>
      </el-aside>
      <el-container>
        <el-header>Header</el-header>
        <el-main>
          <router-view />
        </el-main>
        <el-footer>Footer</el-footer>
      </el-container>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import collapse from '@/layout/collapse/collapse.vue';
import header from '@/layout/header/header.vue';
</script>