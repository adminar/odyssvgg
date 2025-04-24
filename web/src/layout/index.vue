<template>
  <div class="common-layout">
    <el-container>
      <el-aside width="200px">
        <collapse/>
      </el-aside>
      <el-container>
        <el-header>
          <navigation />
        </el-header>
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
import navigation from '@/layout/navigation/navigation.vue';
</script>