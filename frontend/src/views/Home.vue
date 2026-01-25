<template>
  <div class="home-container">
    <h1 class="page-title">安徽工业大学 · 校园生活助手</h1>
    <p class="page-intro">一站式解决您的校园生活需求，让学习与生活更加便捷高效</p>

    <div class="tools-grid">
      <div 
        v-for="tool in tools" 
        :key="tool.id"
        class="tool-card"
        @click="handleToolClick(tool)"
      >
        <div class="tool-icon">{{ tool.icon }}</div>
        <div class="tool-name">{{ tool.name }}</div>
        <div class="tool-status" :class="{ 'status-dev': tool.status === 'dev' }">
          {{ tool.status === 'ready' ? '可用' : '开发中' }}
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'

const router = useRouter()

const tools = [
  { id: 'grades', name: '成绩查询', icon: '📊', status: 'ready' },
  { id: 'electricity', name: '电费查询', icon: '⚡', status: 'ready' },
  { id: 'schedule', name: '课程表', icon: '📅', status: 'ready' },
  { id: 'news', name: '校园新闻', icon: '📰', status: 'ready' },
  { id: 'library', name: '图书馆', icon: '📚', status: 'dev' },
  { id: 'campus', name: '校园地图', icon: '🗺️', status: 'dev' },
  { id: 'dining', name: '食堂查询', icon: '🍽️', status: 'dev' },
  { id: 'bus', name: '校车时刻', icon: '🚌', status: 'dev' }
]

function handleToolClick(tool: any) {
  if (tool.status === 'ready' && tool.id === 'grades') {
    router.push('/main/grades')
  } else if (tool.status === 'ready' && tool.id === 'electricity') {
    router.push('/main/electricity')
  } else if (tool.status === 'ready' && tool.id === 'schedule') {
    router.push('/main/schedule')
  } else if (tool.status === 'ready' && tool.id === 'news') {
    router.push('/main/news')
  } else {
    router.push('/main/coming-soon')
  }
}
</script>

<style scoped>
.home-container {
  padding: 24px;
  max-width: 1200px;
  margin: 0 auto;
}

.page-title {
  font-size: 2rem;
  color: #4a5568;
  margin-bottom: 16px;
  text-align: center;
  font-weight: 600;
}

.page-intro {
  font-size: 1.1rem;
  color: #718096;
  text-align: center;
  margin-bottom: 32px;
  max-width: 600px;
  margin-left: auto;
  margin-right: auto;
}

.tools-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 24px;
}

.tool-card {
  background-color: #e0e5ec;
  border-radius: 20px;
  padding: 24px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.3s ease;
  box-shadow: 6px 6px 10px 0 rgba(163,177,198, 0.7), -6px -6px 10px 0 rgba(255,255,255, 0.8);
  position: relative;
  overflow: hidden;
}

.tool-card::before {
  content: "";
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255,255,255,0.2), transparent);
  transition: left 0.5s ease;
}

.tool-card:hover::before {
  left: 100%;
}

.tool-card:hover {
  box-shadow: 4px 4px 8px 0 rgba(163,177,198, 0.7), -4px -4px 8px 0 rgba(255,255,255, 0.8);
  transform: translateY(-4px) scale(1.02);
}

.tool-card:active {
  box-shadow: inset 4px 4px 8px 0 rgba(163,177,198, 0.7), inset -4px -4px 8px 0 rgba(255,255,255, 0.8);
  transform: translateY(0) scale(0.98);
}

.tool-icon {
  font-size: 3rem;
  margin-bottom: 16px;
}

.tool-name {
  font-size: 1.1rem;
  color: #4a5568;
  font-weight: 500;
  margin-bottom: 12px;
}

.tool-status {
  font-size: 0.85rem;
  padding: 4px 12px;
  border-radius: 12px;
  background-color: #e0e5ec;
  color: #3f8600;
  box-shadow: 3px 3px 6px 0 rgba(163,177,198, 0.7), -3px -3px 6px 0 rgba(255,255,255, 0.8);
}

.tool-status.status-dev {
  color: #cf1322;
}

/* 工具卡片进入动画 */
.slide-up-enter-active {
  transition: all 0.5s ease;
}

.slide-up-enter-from {
  opacity: 0;
  transform: translateY(30px);
}

/* 为每个工具卡片添加延迟动画效果 */
.tool-card:nth-child(1) { animation-delay: 0.05s; }
.tool-card:nth-child(2) { animation-delay: 0.1s; }
.tool-card:nth-child(3) { animation-delay: 0.15s; }
.tool-card:nth-child(4) { animation-delay: 0.2s; }
.tool-card:nth-child(5) { animation-delay: 0.25s; }
.tool-card:nth-child(6) { animation-delay: 0.3s; }
.tool-card:nth-child(7) { animation-delay: 0.35s; }
.tool-card:nth-child(8) { animation-delay: 0.4s; }
.tool-card:nth-child(9) { animation-delay: 0.45s; }
</style>
