import { defineStore } from "pinia";

export const useMainStore = defineStore("main", {
  state: () => ({
    title: "Hello Element Plus"
  })
});

