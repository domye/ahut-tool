// frontend/src/router/index.ts
// 路由配置文件，包含懒加载和路由守卫

import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    redirect: '/main/home'
  },
  {
    path: '/main',
    name: 'Main',
    component: () => import('../views/Main.vue'),
    children: [
      {
        path: 'home',
        name: 'Home',
        component: () => import('../views/Home.vue')
      },
      {
        path: 'grades',
        name: 'Grades',
        component: () => import('../views/Grades.vue')
      },
      {
        path: 'electricity',
        name: 'Electricity',
        component: () => import('../views/Electricity.vue')
      },
      {
        path: 'schedule',
        name: 'Schedule',
        component: () => import('../views/Schedule.vue')
      },
      {
        path: 'news',
        name: 'News',
        component: () => import('../views/News.vue')
      },
      {
        path: 'about',
        name: 'About',
        component: () => import('../views/About.vue')
      },
      {
        path: 'coming-soon',
        name: 'ComingSoon',
        component: () => import('../views/ComingSoon.vue')
      },
      {
        path: 'config',
        name: 'Config',
        component: () => import('../views/Config.vue')
      }
    ]
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  next()
})

export default router
