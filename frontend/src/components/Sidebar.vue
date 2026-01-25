<script lang="ts" setup>
  import { ref, onMounted, onUnmounted } from 'vue'
  import { useRouter } from 'vue-router'

  const router = useRouter()
  const activeItem = ref('home')
  const isMobile = ref(false)

  const navItems = [
    { id: 'home', label: '主页', icon: '🏠' },
    { id: 'dev1', label: '待开发', icon: '🚧' },
    { id: 'dev2', label: '待开发', icon: '🚧' },
    { id: 'about', label: '关于', icon: 'ℹ️' }
  ]

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
  <div class="sidebar" :class="{ 'top-bar': isMobile }">
    <div class="logo" v-if="!isMobile">
      <h3>安徽工业大学</h3>
      <p class="subtitle">生活工具</p>
    </div>
    <div class="logo-mobile" v-else>
      <h3>安徽工业大学</h3>
    </div>
    <nav class="nav-menu" :class="{ 'horizontal': isMobile }">
      <div
              v-for="item in navItems"
              :key="item.id"
              class="nav-item"
              :class="{ active: activeItem === item.id }"
              @click="handleNavigate(item.id)"
      >
        <span class="icon">{{ item.icon }}</span>
        <span class="text" v-if="!isMobile">{{ item.label }}</span>
      </div>
    </nav>
    <div class="footer" v-if="!isMobile">
      <p>© 2026 Domye</p>
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
    transition: all 0.3s ease;
    z-index: 1000;
  }

  /* 移动端顶栏样式 */
  .sidebar.top-bar {
    width: 100%;
    height: auto;
    flex-direction: row;
    padding: 12px 16px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.05);
    align-items: center;
    justify-content: space-between;
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

  .logo-mobile {
    margin-right: 16px;
  }

  .logo-mobile h3 {
    margin: 0;
    font-size: 1.1rem;
    font-weight: 600;
    color: #4a5568;
  }

  .nav-menu {
    flex: 1;
    display: flex;
    flex-direction: column;
    gap: 16px;
  }

  /* 移动端水平导航菜单 */
  .nav-menu.horizontal {
    flex-direction: row;
    gap: 8px;
    justify-content: center;
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

  /* 移动端导航项样式调整 */
  .top-bar .nav-item {
    padding: 8px 12px;
    border-radius: 12px;
  }

  .top-bar .nav-item .icon {
    margin-right: 0;
    font-size: 1.2rem;
  }
</style>
