<template>
  <div>
    <div class="login_form_box">
      <a-form :rules="rules" :model="form" layout="vertical" @submit="onSubmit">
        <a-form-item field="username" :hide-asterisk="true">
          <a-input v-model="form.username" allow-clear placeholder="请输入账号">
            <template #prefix>
              <icon-user />
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
        <a-form-item field="verifyCode" :hide-asterisk="true">
          <div class="verifyCode">
            <a-input style="width: 160px" v-model="form.verifyCode" allow-clear placeholder="请输入验证码" />
            <VerifyCode :content-height="40" :content-width="100" @verify-code-change="verifyCodeChange" />
          </div>
        </a-form-item>
        <a-form-item field="remember">
          <div class="remember">
            <a-checkbox v-model="form.remember">记住密码</a-checkbox>
            <div class="forgot-password">忘记密码</div>
          </div>
        </a-form-item>
        <a-form-item>
          <a-button long type="primary" html-type="submit">登录</a-button>
          <a-button long type="secondary" @click="loginWithOIDC">使用 OIDC 登录</a-button>
        </a-form-item>
      </a-form>
    </div>
    <div class="register" @click="goToRegister">注册账号</div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import { Message } from "@arco-design/web-vue";
import { useRouter } from "vue-router";
import { useUserInfoStore } from "@/store/modules/user-info";
import axios from "axios";

const API_BASE_URL = import.meta.env.VITE_APP_BASE_URL;
const router = useRouter();
const form = ref({
  username: null,
  password: null,
  verifyCode: null,
  remember: false
});
const rules = ref({
  username: [
    {
      required: true,
      message: "请输入账号"
    },
  ],
  password: [
    {
      required: true,
      message: "请输入密码"
    },
  ],
  verifyCode: [
    {
      required: true,
      message: "请输入验证码"
    },
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
  username: null,
  password: null,
  verifyCode: ""
});
const verifyCodeChange = (code: string) => (verify.value.verifyCode = code);

// 定义类型
interface RolesPro {
  project: string[];
  env: string[];
  roles: string[];
}

// 登录提交处理逻辑
const onSubmit = async ({ errors }: any) => {
  if (errors) return;
  try {
    const params = {
      username: form.value.username ?? '',
      password: form.value.password ?? ''
    };
    const response = await axios.get(`${API_BASE_URL}/login`, { params });
    if (response.data.status === 200)  {
      Message.success(response.data.message);
      await fetchUserInfo(response.data.result);
      router.replace("/home");
    } else {
      Message.error(response.data.message);
    }
  } catch (error) {
    console.log(error)
    Message.error("登录失败，请重试");
  }
};

// 登录提交处理逻辑
const fetchUserInfo = async (token: string) => {
  try {
    console.log(document.cookie);
    const response = await axios.get(`${API_BASE_URL}/login/fetch/user`, { headers: { Authorization: `Bearer ${token}` } });
    if (response.data.status === 200)  {
      //const dataRoles = response.data.result.roles;
      //const rolesArray = dataRoles.split(",").map((role: string) => role.trim());
      let parsedRolesData: RolesPro | null = null;
      try {
        // 先解析 JSON 字符串
        parsedRolesData = JSON.parse(response.data.result.roles) as RolesPro;
      } catch (error) {
        console.error("JSON 解析失败:", error);
      }
      let stores = useUserInfoStore();
      let account = {
        id: response.data.result.id,
        username: response.data.result.username,
        project: parsedRolesData?.project || [],
        env: parsedRolesData?.env || [],
        roles: parsedRolesData?.roles || [],
      };
      stores.setAccount(account);
      stores.setToken(response.data.result.auth_token);
    } else {
      Message.error(response.data.message);
    }
  } catch (error) {
    console.log(error)
    Message.error("用户信息获取失败，请重试");
  }
};

// 注册跳转处理
const goToRegister = () => {
  router.push("/register"); // 跳转到注册页面
};

// OIDC 登录处理
const handleOIDCLogin = async () => {
  const urlParams = new URLSearchParams(window.location.search);
  const token = urlParams.get("token");
  if (token) {
    await fetchUserInfo(token);
    // 使用 replaceState 清除 URL 参数
    window.history.replaceState({}, document.title, window.location.pathname);
  }
  router.replace("/home");
};

// 触发 OIDC 登录跳转
const loginWithOIDC = () => {
  window.location.href = `${API_BASE_URL}/login/oidc`;
};

// 在组件挂载时处理 OIDC 登录
onMounted(handleOIDCLogin);

</script>

<style lang="scss" scoped>
.login_form_box {
  margin-top: 28px;
  .verifyCode {
    width: 100%;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
  .remember {
    width: 100%;
    display: flex;
    justify-content: space-between;
    align-items: center;
    .forgot-password {
      color: $color-primary;
      cursor: pointer;
    }
  }
}
.register {
  text-align: center;
  color: $color-text-3;
  font-size: $font-size-body-1;
  cursor: pointer;
  &:hover {
    color: $color-primary; // 鼠标悬停时变色
  }
}
</style>
