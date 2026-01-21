import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    redirect: '/main/home'
  },
  // {
  //   path: '/login',
  //   name: 'Login',
  //   component: () => import('../views/Login.vue')
  // },
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
        path: 'about',
        name: 'About',
        component: () => import('../views/About.vue')
      },
      {
        path: 'coming-soon',
        name: 'ComingSoon',
        component: () => import('../views/ComingSoon.vue')
      }
    ]
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// 路由守卫
router.beforeEach((to, from, next) => {
  // 由于应用启动时会自动登录，所以不再需要登录检查
  next()
})

export default router
