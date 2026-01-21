<template>
  <div class="main-layout">
    <Sidebar />
    <div class="main-content" :class="{ 'mobile-content': isMobile }">
      <Header v-if="!isMobile" />
      <div class="content" :class="{ 'mobile-content-inner': isMobile }">
        <router-view />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '../store/user'
import Sidebar from '../components/Sidebar.vue'
import Header from '../components/Header.vue'

const router = useRouter()
const userStore = useUserStore()
const isMobile = ref(false)

// 检查屏幕宽度是否为移动设备
const checkScreenSize = () => {
  isMobile.value = window.innerWidth <= 768
}

// 组件挂载时检查屏幕尺寸
onMounted(() => {
  checkScreenSize()
  window.addEventListener('resize', checkScreenSize)
})

// 组件卸载时移除事件监听
onUnmounted(() => {
  window.removeEventListener('resize', checkScreenSize)
})
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
  transition: all 0.3s ease;
}

/* 移动端内容区域样式 */
.main-content.mobile-content {
  margin-left: 0;
  margin-top: 60px; /* 为顶栏留出空间 */
}



.content {
  flex: 1;
  padding: 24px;
  background-color: #e0e5ec;
  overflow-y: auto;
  overflow-x: hidden;
  min-height: 0;
}

/* 移动端内容区域样式 */
.content.mobile-content-inner {
  padding: 16px;
}
</style>
