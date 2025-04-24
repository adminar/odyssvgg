import { defineStore } from "pinia";
import persistedstateConfig from "@/store/config/index";
/**
 * 用户信息
 * @methods setAccount 设置账号信息
 * @methods setToken 设置token
 * @methods logOut 退出登录
 */

// 定义账户信息的类型
interface Account {
  id: number;
  username: string;
  roles: string[];
  project: string[];
  env: string[];
}

// 定义状态的类型
interface UserInfoState {
  account: Account;
  token: string;
}

export const useUserInfoStore = defineStore("user-info", {
  state: (): UserInfoState => ({
    account: {
      id: 0,
      username: "",
      roles: [],
      project: [],
      env: []
    }, // 账号信息
    token: "" // token
  }),
  actions: {
    async setAccount(data: Account) {
      this.account = data;
    },
    async setToken(data: string) {
      this.token = data;
    },
    async logOut() {
      this.account = { id: 0, username: "", roles: [], project: [], env: [] };
      this.token = "";
    }
  },
  persist: persistedstateConfig("user-info")
});
