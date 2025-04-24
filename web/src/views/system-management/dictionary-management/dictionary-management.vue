<template>
  <div class="odyssey-page">
    |-- 0:状态未知 --|-- 
    1:待构建 --|-- 
    2:构建中 --|-- 
    3:构建成功,待审核(prod环境) --|-- 
    3:构建成功 --|-- 
    4:构建失败 --|-- 
    5:审核中 --|-- 
    6:审核通过,待发布 --|-- 
    7:审核已拒绝 --|-- 
    8:发布中 --|-- 
    9:发布成功 --|-- 
    10:发布失败 --|
  </div>
</template>

<script setup lang="ts"></script>

<style lang="scss" scoped></style>
