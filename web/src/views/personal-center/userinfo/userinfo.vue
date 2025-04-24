<template>
  <div class="odyssey-page">
    <div class="odyssey-inner-page">
      <div class="my-avatar">
        <a-avatar :size="100">
         <!-- <img alt="avatar" :src="myImage" /> -->
        </a-avatar>
        <div class="my-name">
          <div class="my-title title-size-1">车联网研发部</div>
          <div class="my-local"><icon-location />北京朝阳</div>
        </div>
      </div>
      <a-grid :cols="{ xs: 1, sm: 1, md: 1, lg: 2, xl: 3 }" :col-gap="16" :row-gap="16">
        <a-grid-item>
          <div class="title-size-1">👋关于我</div>
          <div class="type-row" v-for="item in list.aboutMe" :key="item.label">
            <div class="column-title-size-1">{{ item.label }}：</div>
            <div v-for="(type, index) in item.value" :key="index">
              <a-tag :color="item.theme">
                {{ type }}
              </a-tag>
            </div>
          </div>
        </a-grid-item>
        <a-grid-item>
          <div class="title-size-1">💬联系我</div>
          <div class="type-row" v-for="item in list.callMe" :key="item.label">
            <div class="column-title-size-2">{{ item.label }}：</div>
            <div :class="item.link ? 'link-hover' : ''" @click="onLink(item)">{{ item.value }}</div>
          </div>
        </a-grid-item>
        <a-grid-item>
          <div class="title-size-1">🐾其它</div>
          <div class="type-row" v-for="item in list.other" :key="item.label">
            <div class="column-title-size-2">{{ item.label }}：</div>
            <div>{{ item.value }}</div>
          </div>
          <!-- <a-image width="100" :src="officialAccount" /> -->
        </a-grid-item>
        <a-grid-item>
        <!-- <a-image default-scale height="220" :src="otherImage" /> -->
        </a-grid-item>
        <a-grid-item>
          <div class="title-size-2">✨Hello 朋友们✨</div>
          <div class="type-row" v-for="item in list.introduce" :key="item.label">
            <div class="text-ellipsis">{{ item.label }}</div>
          </div>
        </a-grid-item>
      </a-grid>
    </div>
  </div>
</template>

<script setup lang="ts">
//import myImage from "@/assets/img/my-image.jpg";
//import otherImage from "@/assets/img/other-image.jpg";
//import officialAccount from "@/assets/img/official-account.png";

const list = reactive({
  aboutMe: [
    { label: "💼职业工作", value: ["neolix"], theme: "orangered" },
  ],
  callMe: [
    { label: "📧邮箱", value: "@neolix.cn", link: false },
  ],
  other: [
    { label: "🧭公众号", value: "neolix" },
  ],
  introduce: [
    { label: "打造移动服务空间，让美好生活触手可得" }
  ]
});

const onLink = (item: any) => {
  if (!item.link) return;
  window.open(item.url);
};
</script>

<style lang="scss" scoped>
.border {
  border: 1px solid red;
}
.my-avatar {
  display: flex;
  align-items: center;
  margin-bottom: calc($margin * 2);
  .my-name {
    height: 60px;
    margin-left: $margin;
    .my-title {
      height: 30px;
    }
    .my-local {
      font-size: $font-size-body-3;
      height: 30px;
      display: flex;
      align-items: end;
    }
  }
}

.type-row {
  height: 25px;
  margin-top: $margin;
  display: flex;
  align-items: center;
  column-gap: $margin;
  .column-title-size-1 {
    min-width: 90px;
  }
  .column-title-size-2 {
    min-width: 68px;
  }
  .text-ellipsis {
    white-space: nowrap;
  }
}

.title-size-1 {
  font-size: $font-size-title-1;
  font-weight: bold;
}
.title-size-2 {
  font-size: $font-size-title-2;
  font-weight: bold;
}

.link-hover {
  cursor: pointer;
  color: $color-primary;
}
</style>
