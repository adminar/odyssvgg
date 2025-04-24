import { RouteRecordRaw, createRouter, createWebHistory } from "vue-router";

const routes: RouteRecordRaw[] = [
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
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/login/login.vue')
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
