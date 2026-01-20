<template>
  <div class="main-layout">
    <Sidebar />
    <div class="main-content">
      <div class="header">
        <div class="user-info">
          <span>欢迎，{{ userStore.userId }}</span>
        </div>
        <button class="logout-btn" @click="handleLogout">
          <span class="logout-icon">🚪</span>
          <span class="logout-text">退出登录</span>
        </button>
      </div>
      <div class="content">
        <router-view />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useUserStore } from '../store/user'
import Sidebar from '../components/Sidebar.vue'

const router = useRouter()
const userStore = useUserStore()

function handleLogout() {
  userStore.logout()
  router.push('/login')
}
</script>

<style scoped>
.main-layout {
  display: flex;
  height: 100vh;
  background-color: #e0e5ec;
  overflow: hidden;
}

.main-content {
  flex: 1;
  margin-left: 240px;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.header {
  background-color: #e0e5ec;
  padding: 16px 24px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  flex-shrink: 0;
}

.user-info {
  font-size: 1rem;
  color: #4a5568;
  font-weight: 500;
}

.logout-btn {
  background-color: #e0e5ec;
  box-shadow: 6px 6px 10px 0 rgba(163,177,198, 0.7), -6px -6px 10px 0 rgba(255,255,255, 0.8);
  border-radius: 12px;
  border: none;
  color: #4a5568;
  padding: 10px 20px;
  font-size: 0.95rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
}

.logout-btn:hover {
  box-shadow: 4px 4px 8px 0 rgba(163,177,198, 0.7), -4px -4px 8px 0 rgba(255,255,255, 0.8);
}

.logout-btn:active {
  box-shadow: inset 4px 4px 8px 0 rgba(163,177,198, 0.7), inset -4px -4px 8px 0 rgba(255,255,255, 0.8);
}

.logout-icon {
  font-size: 1.1rem;
  margin-right: 8px;
}

.logout-text {
  font-size: 0.95rem;
}

.content {
  flex: 1;
  padding: 24px;
  background-color: #e0e5ec;
  overflow-y: auto;
  overflow-x: hidden;
  min-height: 0;
}
</style>
