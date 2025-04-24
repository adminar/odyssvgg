<template>
  <div class="container">
    <div class="login">
      <LoginBanner />
      <div class="login_box">
        <div class="login_title">Welcome To Odyssey</div>
        <div class="login_title_desc">基于GO X Vue3.x的运维中心，功能管理应有尽有</div>
        <div class="login_title_desc">CI/CD功能全覆盖</div>
        <LoginForm />
        <div class="author">by devops</div> 
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import LoginBanner from "@/views/login/components/login-banner.vue";
import LoginForm from "@/views/login/components/login-form.vue";
</script>

<style lang="scss" scoped>
.container {
  height: 100vh;
  position: relative;
  overflow: hidden;
  .login {
    width: 1000px;
    height: 500px;
    left: 50%;
    top: 50%;
    position: absolute;
    transform: translate(-50%, -50%);
    display: flex;
    align-items: center;
    box-shadow: 0 0 8px 1px $color-fill-2;

    .login_box {
      width: 350px;
      height: 100%;
      box-sizing: border-box;
      padding: 40px 30px 30px 30px;
      position: relative;
      .login_title {
        font-size: $font-size-title-2;
        color: $color-text-1;
        margin-bottom: $margin-text;
      }
      .login_title_desc {
        font-size: $font-size-body-1;
        color: $color-text-3;
      }
      .author {
        color: $color-text-4;
        font-size: $font-size-body-1;
        position: absolute;
        bottom: 30px;
      }
    }
  }
}
/* 手机端适配 */
@media (max-width: 768px) {
  .login {
    flex-direction: column;
    width: 100%;
    min-height: auto;
    box-shadow: none;
  }

  .banner {
    width: 100%;
    padding: 20px;
  }

  .login_box {
    width: 100%;
    padding: 20px;
  }

  .login_title {
    font-size: 20px;
  }

  .login_title_desc {
    font-size: 14px;
  }

  .author {
    font-size: 12px;
  }
}
</style>
