<template>
  <div>
    <ElButton type="primary" @click="loginAdmin">切换到管理员</ElButton>
    <ElButton type="primary" @click="loginUser">切换到用户</ElButton>
  </div>
</template>

<script setup>
import useUserStore from '@/stores/useUserStore'
import { ElMessage } from 'element-plus'
const userStore = useUserStore()

function loginAdmin() {
  userStore
    .loginApp({
      name: 'admin',
      password: '123'
    })
    .then(() => {
      ElMessage.success('切换管理员角色成功')
    })
}

function loginUser() {
  userStore
    .loginApp({
      name: 'user',
      password: '123'
    })
    .then(() => {
      ElMessage.success('切换用户角色成功')
    })
}

</script>

<style scoped></style>
