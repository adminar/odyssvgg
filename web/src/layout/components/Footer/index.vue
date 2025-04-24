<template>
  <a-layout-footer class="footer">
    <div class="footer_title" @click="onFooter">2024 © odyssey by riyoukou</div>
  </a-layout-footer>
</template>

<script setup lang="ts">
const onFooter = () => {
  window.open("https://github.com/Riyoukou/odyssey", "_blank");
};
</script>

<style lang="scss" scoped>
.footer {
  height: 30px;
  box-sizing: border-box;
  border-top: $border-1 solid $color-border-2;
  display: flex;
  align-items: center;
}
.footer_title {
  text-align: center;
  margin: 0 auto;
  color: $color-text-4;
  cursor: pointer;
}
</style>
