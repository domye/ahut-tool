<!-- frontend/src/App.vue -->
// 应用根组件，包含路由视图和自动登录轮询

<script setup lang="ts">
import { onMounted, onUnmounted } from 'vue'
import { useUserStore } from './store/user'
import { useElectricityStore } from './store/electricity'
import { message } from 'ant-design-vue'

const userStore = useUserStore()
const electricityStore = useElectricityStore()

const POLLING_INTERVAL = 30 * 60 * 1000
let pollingTimer: number | null = null

const performLogin = async () => {
  try {
    await userStore.login()
    await electricityStore.login()
  } catch (error) {
    if (userStore.message === '请先配置教务系统') {
      message.warning('请先配置教务系统')
    } else if (electricityStore.message === '请先配置缴费系统') {
      message.warning('请先配置缴费系统')
    } else if (userStore.message === '登录失败' || electricityStore.message === '登录失败') {
      message.error('登录失败')
    }
    console.error('登录失败:', error)
  }
}

const startPolling = () => {
  performLogin()
  pollingTimer = window.setInterval(() => {
    performLogin()
  }, POLLING_INTERVAL)
}

const stopPolling = () => {
  if (pollingTimer !== null) {
    window.clearInterval(pollingTimer)
    pollingTimer = null
  }
}

onMounted(() => {
  startPolling()
})

onUnmounted(() => {
  stopPolling()
})
</script>

<template>
  <router-view v-slot="{ Component }">
    <transition name="page" mode="out-in">
      <component :is="Component" />
    </transition>
  </router-view>
</template>

