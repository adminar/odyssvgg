<template>
  <div class="center-wrapper">
    <div class="register_form_box">
      <a-form :rules="rules" :model="form" layout="vertical" @submit="onSubmit">
        <a-form-item field="username" :hide-asterisk="true">
          <a-input v-model="form.username" allow-clear placeholder="请输入账号">
            <template #prefix>
              <icon-user />
            </template>
          </a-input>
        </a-form-item>
        <a-form-item field="email" :hide-asterisk="true">
          <a-input v-model="form.email" allow-clear placeholder="请输入邮箱">
            <template #prefix>
              <icon-mail />
            </template>
          </a-input>
        </a-form-item>
        <a-form-item field="phone" :hide-asterisk="true">
          <a-input v-model="form.phone" allow-clear placeholder="请输入手机号" type="number" @input="validatePhoneInput">
            <template #prefix>
              <icon-mail />
            </template>
          </a-input>
        </a-form-item>
        <a-form-item field="password" :hide-asterisk="true">
          <a-input-password v-model="form.password" allow-clear placeholder="请输入密码">
            <template #prefix>
              <icon-lock />
            </template>
          </a-input-password>
        </a-form-item>
        <a-form-item field="confirmPassword" :hide-asterisk="true">
          <a-input-password v-model="form.confirmPassword" allow-clear placeholder="请确认密码">
            <template #prefix>
              <icon-lock />
            </template>
          </a-input-password>
        </a-form-item>
        <a-form-item field="verifyCode" :hide-asterisk="true">
          <div class="verifyCode">
            <a-input style="width: 160px" v-model="form.verifyCode" allow-clear placeholder="请输入验证码" />
            <VerifyCode :content-height="40" :content-width="100" @verify-code-change="verifyCodeChange" />
          </div>
        </a-form-item>
        <a-form-item>
          <a-button long type="primary" html-type="submit">注册</a-button>
        </a-form-item>
      </a-form>
    </div>
    <div class="login" @click="redirectToLogin">已有账号？登录</div>
  </div>
</template>

<script setup lang="ts">
import { Message } from "@arco-design/web-vue";
import { useRouter } from "vue-router";
import axios from "axios";
import { ref } from "vue";

// @ts-ignore
const API_BASE_URL = import.meta.env.VITE_APP_BASE_URL;
const router = useRouter();

const form = ref({
  username: null,
  email: null,
  phone: null,
  password: null,
  confirmPassword: null,
  verifyCode: null
});

const rules = ref({
  username: [
    { required: true, message: "请输入账号" }
  ],
  email: [
    { required: true, message: "请输入邮箱" },
    { type: "email", message: "请输入有效的邮箱" }
  ],
  phone: [
    { required: true, message: "请输入手机号" },
    {
      validator: (value: string, cb: any) => {
        const phonePattern = /^\d{11}$/; // 只允许11位数字
        if (!phonePattern.test(value)) {
          cb("手机号必须为11位数字");
        } else {
          cb();
        }
      }
    }
  ],
  password: [
    { required: true, message: "请输入密码" }
  ],
  confirmPassword: [
    { required: true, message: "请确认密码" },
    {
      validator: (value: string, cb: any) => {
        if (value !== form.value.password) {
          cb("两次输入的密码不一致");
        } else {
          cb();
        }
      }
    }
  ],
  verifyCode: [
    { required: true, message: "请输入验证码" },
    {
      validator: (value: string, cb: any) => {
        if (value !== verify.value.verifyCode) {
          cb("请输入正确的验证码");
        } else {
          cb();
        }
      }
    }
  ]
});

const verify = ref({
  verifyCode: ""
});

const verifyCodeChange = (code: string) => (verify.value.verifyCode = code);

const validatePhoneInput = (event: Event) => {
  const input = event.target as HTMLInputElement;
  // 只保留数字，删除其他字符
  input.value = input.value.replace(/[^\d]/g, '');
  
  // 限制为11位数字
  if (input.value.length > 11) {
    input.value = input.value.slice(0, 11);
  }
  
  // 更新表单数据
  form.value.phone = input.value; 
};

const onSubmit = async ({ errors }: any) => {
  if (errors) return;
  const params = {
    username: form.value.username ?? '',
    email: form.value.email ?? '',
    phone: form.value.phone ?? '',
    password: form.value.password ?? ''
  };
  try {
    // 发送注册请求到后端
    const response = await axios.get(`${API_BASE_URL}/login/register`, { params });
    // 注册成功，跳转到登录页面
    if (response.data.status === 200) {
      Message.success(response.data.message);
      router.replace("/login");
    } else {
      Message.error(response.data.message);
    }
  } catch (error) {
    Message.error("注册失败，请重试");
  }
};

// 跳转到登录页面
const redirectToLogin = () => {
  router.push("/login");
};
</script>

<style lang="scss" scoped>
.center-wrapper {
  display: flex;
  justify-content: center;  /* 水平居中 */
  align-items: center;      /* 垂直居中 */
  height: 100vh;            /* 占满视口高度 */
  background-color: #f5f5f5; /* 背景色 */
}

.register_form_box {
  width: 400px;             /* 表单宽度 */
  padding: 20px;            /* 内边距 */
  background-color: #fff;   /* 白色背景 */
  box-shadow: 0px 4px 12px rgba(0, 0, 0, 0.1); /* 阴影效果 */
  border-radius: 8px;       /* 圆角 */
}

.verifyCode {
  width: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.login {
  text-align: center;
  color: #1890ff;
  font-size: 14px;
  margin-top: 16px;
  cursor: pointer;
}
</style>
