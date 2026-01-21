<script setup lang="ts">
import { onMounted, onUnmounted } from 'vue'
import { useUserStore } from './store/user'
import { useElectricityStore } from './store/electricity'

const userStore = useUserStore()
const electricityStore = useElectricityStore()

// 轮询间隔时间（30分钟）
const POLLING_INTERVAL = 30 * 60 * 1000
let pollingTimer: number | null = null

// 执行登录操作
const performLogin = async () => {
  try {
    await userStore.login()
    await electricityStore.login()
  } catch (error) {
    console.error('登录失败:', error)
  }
}

// 启动轮询
const startPolling = () => {
  // 先执行一次登录
  performLogin()
  // 设置定时器，每30分钟执行一次登录
  pollingTimer = window.setInterval(() => {
    performLogin()
  }, POLLING_INTERVAL)
}

// 停止轮询
const stopPolling = () => {
  if (pollingTimer !== null) {
    window.clearInterval(pollingTimer)
    pollingTimer = null
  }
}

// 在应用挂载时自动登录并启动轮询
onMounted(() => {
  startPolling()
})

// 在应用卸载时停止轮询
onUnmounted(() => {
  stopPolling()
})
</script>

<template>
  <router-view />
</template>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

#app {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Fira Sans', 'Droid Sans', 'Helvetica Neue', sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  background-color: #e0e5ec;
  min-height: 100vh;
}

/* 新拟态风格基础样式 */
.neumorphic {
  background-color: #e0e5ec;
  box-shadow: 9px 9px 16px rgb(163,177,198,0.6), -9px -9px 16px rgba(255,255,255, 0.5);
  border-radius: 16px;
}

.neumorphic-inset {
  background-color: #e0e5ec;
  box-shadow: inset 6px 6px 10px 0 rgba(163,177,198, 0.7), inset -6px -6px 10px 0 rgba(255,255,255, 0.8);
  border-radius: 16px;
}

.neumorphic-btn {
  background-color: #e0e5ec;
  box-shadow: 6px 6px 10px 0 rgba(163,177,198, 0.7), -6px -6px 10px 0 rgba(255,255,255, 0.8);
  border-radius: 12px;
  border: none;
  color: #4a5568;
  transition: all 0.3s ease;
  cursor: pointer;
}

.neumorphic-btn:hover {
  box-shadow: 4px 4px 8px 0 rgba(163,177,198, 0.7), -4px -4px 8px 0 rgba(255,255,255, 0.8);
}

.neumorphic-btn:active, .neumorphic-btn.active {
  box-shadow: inset 4px 4px 8px 0 rgba(163,177,198, 0.7), inset -4px -4px 8px 0 rgba(255,255,255, 0.8);
}

.neumorphic-card {
  background-color: #e0e5ec;
  box-shadow: 8px 8px 16px rgb(163,177,198,0.6), -8px -8px 16px rgba(255,255,255, 0.5);
  border-radius: 20px;
  padding: 24px;
}
</style>
