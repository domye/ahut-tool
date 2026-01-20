<script lang="ts" setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const activeItem = ref('home')

const navItems = [
  { id: 'home', label: '主页', icon: '🏠' },
  { id: 'dev1', label: '待开发', icon: '🚧' },
  { id: 'dev2', label: '待开发', icon: '🚧' },
  { id: 'about', label: '关于', icon: 'ℹ️' }
]

function handleNavigate(id: string) {
  activeItem.value = id
  if (id === 'home') {
    router.push('/main/home')
  } else if (id === 'about') {
    router.push('/main/about')
  } else {
    // 待开发页面
    router.push('/main/coming-soon')
  }
}
</script>

<template>
  <div class="sidebar">
    <div class="logo">
      <h3>安徽工业大学</h3>
      <p class="subtitle">生活工具</p>
    </div>
    <nav class="nav-menu">
      <div
        v-for="item in navItems"
        :key="item.id" 
        class="nav-item" 
        :class="{ active: activeItem === item.id }"
        @click="handleNavigate(item.id)"
      >
        <span class="icon">{{ item.icon }}</span>
        <span class="text">{{ item.label }}</span>
      </div>
    </nav>
    <div class="footer">
      <p>© 2024 安徽工业大学</p>
    </div>
  </div>
</template>

<style scoped>
.sidebar {
  width: 240px;
  height: 100vh;
  background-color: #e0e5ec;
  color: #4a5568;
  display: flex;
  flex-direction: column;
  position: fixed;
  left: 0;
  top: 0;
  padding: 24px 16px;
  box-shadow: 4px 0 8px rgba(0, 0, 0, 0.05);
}

.logo {
  margin-bottom: 32px;
  text-align: center;
}

.logo h3 {
  margin: 0 0 8px 0;
  font-size: 1.3rem;
  font-weight: 600;
  color: #4a5568;
}

.subtitle {
  margin: 0;
  font-size: 0.9rem;
  color: #718096;
}

.nav-menu {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.nav-item {
  display: flex;
  align-items: center;
  padding: 16px;
  cursor: pointer;
  transition: all 0.3s;
  position: relative;
  border-radius: 16px;
  background-color: #e0e5ec;
  box-shadow: 6px 6px 10px 0 rgba(163,177,198, 0.7), -6px -6px 10px 0 rgba(255,255,255, 0.8);
}

.nav-item:hover {
  box-shadow: 4px 4px 8px 0 rgba(163,177,198, 0.7), -4px -4px 8px 0 rgba(255,255,255, 0.8);
}

.nav-item.active {
  box-shadow: inset 4px 4px 8px 0 rgba(163,177,198, 0.7), inset -4px -4px 8px 0 rgba(255,255,255, 0.8);
}

.nav-item .icon {
  font-size: 1.3rem;
  margin-right: 12px;
}

.nav-item .text {
  flex: 1;
  font-weight: 500;
}

.footer {
  margin-top: 24px;
  font-size: 0.85rem;
  color: #718096;
  text-align: center;
}

.footer p {
  margin: 0;
}
</style>
