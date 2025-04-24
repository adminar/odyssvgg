import { createRouter, createWebHistory } from "vue-router";

const routes = [
  {
    path: '/',
    component: () => import('@/layout/index.vue'),
    children: [
      {
        path: '',
        name: 'Home',
        component: () => import('@/views/Home.vue')
      }
      // 其他需要显示导航栏的页面
    ]
  }
];

export default createRouter({
  history: createWebHistory(),
  routes
});

