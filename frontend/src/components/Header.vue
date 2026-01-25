<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useUserStore } from '../store/user'

const router = useRouter()
const userStore = useUserStore()

function handleConfig() {
  router.push('/main/config')
}
</script>

<template>
  <div class="header">
    <div class="user-info">
      <span>欢迎使用</span>
    </div>
    <button class="logout-btn" @click="handleConfig">
      <span class="logout-icon">⚙️</span>
      <span class="logout-text">配置</span>
    </button>
  </div>
</template>

<style scoped>
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
</style>
