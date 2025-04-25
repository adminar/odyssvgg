<template>
  <ElDrawer v-model="appStore.themeDrawer" title="控制面板设置" size="320px" append-to-body>
    <div class="pb-6">
      <ElDivider>布局切换</ElDivider>
      <div class="grid grid-cols-3 gap-2">
        <ElTooltip content="经典" placement="bottom">
          <div
            class="layout-box hover:shadow-md flex flex-col justify-between"
            @click="appStore.layout = 'classics'"
          >
            <div class="h-[25%] bg-[var(--el-color-primary)] rounded-sm"></div>
            <div class="h-[65%] flex justify-between">
              <div class="w-[25%] h-full bg-[var(--el-color-primary-light-3)] rounded-sm"></div>
              <div class="main w-[70%]">
                <ElIcon
                  style="--color: var(--el-color-primary)"
                  v-if="appStore.layout == 'classics'"
                >
                  <CircleCheckFilled />
                </ElIcon>
              </div>
            </div>
          </div>
        </ElTooltip>
        <ElTooltip content="纵向" placement="bottom" @click="appStore.layout = 'vertical'">
          <div
            class="layout-box hover:shadow-md flex justify-between"
            @click="appStore.layout = 'vertical'"
          >
            <div class="w-[25%] h-full bg-[var(--el-color-primary)] rounded-sm"></div>
            <div class="w-[70%] h-full flex flex-col justify-between">
              <div class="h-[25%] bg-[var(--el-color-primary-light-3)] rounded-sm"></div>
              <div class="main h-[65%]">
                <ElIcon
                  style="--color: var(--el-color-primary)"
                  v-if="appStore.layout == 'vertical'"
                >
                  <CircleCheckFilled />
                </ElIcon>
              </div>
            </div>
          </div>
        </ElTooltip>
        <ElTooltip content="横向" placement="bottom">
          <div
            class="layout-box hover:shadow-md flex flex-col justify-between"
            @click="appStore.layout = 'horizontal'"
          >
            <div class="h-[25%] bg-[var(--el-color-primary)] rounded-sm"></div>
            <div class="main h-[65%]">
              <ElIcon
                style="--color: var(--el-color-primary)"
                v-if="appStore.layout == 'horizontal'"
              >
                <CircleCheckFilled />
              </ElIcon>
            </div>
          </div>
        </ElTooltip>
      </div>
    </div>
    <div class="pb-6">
      <ElDivider>组件设置</ElDivider>
      <div class="flex justify-between my-4 items-center">
        <div class="text-sm">是否开启导航条</div>
        <ElSwitch v-model="appStore.navTabs"></ElSwitch>
      </div>
      <div class="flex justify-between my-4 items-center">
        <div class="text-sm">是否开启页脚</div>
        <ElSwitch v-model="appStore.footer"></ElSwitch>
      </div>
      <div class="flex justify-between my-4 items-center">
        <div class="text-sm">是否开启面包屑</div>
        <ElSwitch v-model="appStore.breadcrumb"></ElSwitch>
      </div>
      <div class="flex justify-between my-4 items-center">
        <div class="text-sm">导航条风格</div>
        <ElRadioGroup v-model="appStore.navTabsIsBackground">
          <ElRadioButton :label="true">背景</ElRadioButton>
          <ElRadioButton :label="false">简约</ElRadioButton>
        </ElRadioGroup>
      </div>
      <div class="flex justify-between my-4 items-center">
        <div class="text-sm">组件语言</div>
        <ElRadioGroup v-model="appStore.language">
          <ElRadioButton label="en">英文</ElRadioButton>
          <ElRadioButton label="zhCn">中文</ElRadioButton>
        </ElRadioGroup>
      </div>
      <div class="flex justify-between my-4 items-center">
        <div class="text-sm">组件大小</div>
        <ElRadioGroup v-model="appStore.componentSize">
          <ElRadioButton label="large">大</ElRadioButton>
          <ElRadioButton label="default">中</ElRadioButton>
          <ElRadioButton label="small">小</ElRadioButton>
        </ElRadioGroup>
      </div>
    </div>
    <div class="pb-6">
      <ElDivider>全局主题</ElDivider>
      <div>
        <div class="flex justify-between my-4 items-center">
          <div class="text-sm">主题颜色</div>
          <ElColorPicker
            v-model="appStore.primaryColor"
            color-format="hex"
            :predefine="predefineColors"
          >
          </ElColorPicker>
        </div>
        <div class="flex justify-between my-4 items-center">
          <div class="text-sm">黑暗模式</div>
          <ElSwitch
            v-model="appStore.isDark"
            inline-prompt
            active-icon="Sunny"
            inactive-icon="Moon"
          >
          </ElSwitch>
        </div>
        <div class="flex justify-between my-4 items-center">
          <div class="text-sm">灰色模式</div>
          <ElSwitch v-model="appStore.isGray"></ElSwitch>
        </div>
      </div>
    </div>
  </ElDrawer>
</template>

<script setup>
import useAppStore from '@/stores/useAppStore'
const appStore = useAppStore()

const predefineColors = [
  '#2463eb',
  '#0C819F',
  '#409EFF',
  '#27ae60',
  '#ff5c93',
  '#e74c3c',
  '#fd726d',
  '#f39c12',
  '#9b59b6'
]
</script>

<style scoped>
.layout-box {
  border: 1px solid;
  height: 5rem; /* 20px -> 5rem */
  cursor: pointer;
  border-radius: 0.375rem; /* 0.375rem is equivalent to `rounded` */
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1); /* Default shadow */
  overflow: hidden;
  padding: 0.375rem; /* 1.5px -> 0.375rem */
  background-color: white;
}

.layout-box .main {
  background-color: var(--el-color-primary-light-8);
  display: flex;
  justify-content: flex-end;
  align-items: flex-end;
  padding: 0.25rem; /* Adjusted padding */
  border-radius: 0.125rem; /* rounded-sm */
  border: 1px dashed var(--el-color-primary);
}
</style>
